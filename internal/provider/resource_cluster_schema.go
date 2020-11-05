package provider

import (
	"encoding/base64"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/client-go/tools/clientcmd"
)

func resourceClusterSchema() map[string]*schema.Schema {
	return setClusterAttributesSchema(map[string]*schema.Schema{
		// Arguments
		"name": {
			Type:        schema.TypeString,
			Description: `the name of the cluster. corresponds to the --name flag on the kind cli.`,
			Required:    true,
			ForceNew:    true,
		},
		"config": {
			Type:        schema.TypeString,
			Description: `the cluster config as documented on https://kind.sigs.k8s.io/docs/user/configuration/`,
			Optional:    true,
			ForceNew:    true,
		},
		"image": {
			Type:          schema.TypeString,
			Description:   `The image to use for the kind nodes. corresponds to the --image flag on the cli.`,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"image_version"},
		},
		"image_version": {
			Type:          schema.TypeString,
			Description:   `Kubernetes major.minor version, which chooses the correct node image from the published SHAs matching this version of KIND`,
			Optional:      true,
			ForceNew:      true,
			ConflictsWith: []string{"image"},
		},
	})
}

func setClusterAttributesSchema(s map[string]*schema.Schema) map[string]*schema.Schema {
	s["ca_certificate_data"] = stringAttribute(`The base64-encoded CA Certificate used by the API Server`)
	s["client_certificate_data"] = stringAttribute(`The base64-encoded client certificate data for connecting the cluster`)
	s["client_key_data"] = stringAttribute(`The base64-encoded client private key data for connecting the cluster`)
	s["context"] = stringAttribute("The name of the context in KubeConfig")
	s["kubeconfig"] = stringAttribute(`The full text of the kubeconfig that can be used to connect to this cluster`)
	s["nodes"] = &schema.Schema{
		Type:        schema.TypeList,
		Description: `The list of nodes that were provisioned for this cluster`,
		Computed:    true,
		Elem:        nodeSchema(),
	}
	s["server"] = stringAttribute(`Kubernetes API Server URL`)
	return s
}

func setClusterAttributeData(c *Meta, name string, data *schema.ResourceData) (diags diag.Diagnostics) {
	kubeconfig, err := c.getKubeConfig(name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("could not export kubeconfig for cluster %q", name),
			Detail:   err.Error(),
		})
	}
	if err := setKubeConfigData(kubeconfig, data); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("could not set kubeconfig attributes for cluster %q", name),
			Detail:   err.Error(),
		})
	}
	nodes, err := c.getKindNodeList(name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("could not list nodes of cluster %q", name),
			Detail:   err.Error(),
		})
	}
	_ = data.Set("nodes", nodes)
	return
}

func setKubeConfigData(kubeconfig string, data *schema.ResourceData) error {
	cfg, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	if err != nil {
		return fmt.Errorf("could not parse kubeconfig: %w", err)
	}
	_ = data.Set("kubeconfig", kubeconfig)
	raw, err := cfg.RawConfig()
	if err != nil {
		return fmt.Errorf("could not parse kubeconfig: %w", err)
	}
	kctx := raw.Contexts[raw.CurrentContext]
	if kctx == nil {
		return fmt.Errorf("no kubeconfig context %q", raw.CurrentContext)
	}
	_ = data.Set("context", raw.CurrentContext)
	cluster := raw.Clusters[kctx.Cluster]
	if cluster == nil {
		return fmt.Errorf("no kubeconfig context %q", kctx.Cluster)
	}
	_ = data.Set("ca_certificate_data", base64.StdEncoding.EncodeToString(cluster.CertificateAuthorityData))
	_ = data.Set("server", cluster.Server)
	auth := raw.AuthInfos[kctx.AuthInfo]
	if auth == nil {
		return fmt.Errorf("no kubeconfig user %q", kctx.AuthInfo)
	}
	_ = data.Set("client_certificate_data", base64.StdEncoding.EncodeToString(auth.ClientCertificateData))
	_ = data.Set("client_key_data", base64.StdEncoding.EncodeToString(auth.ClientKeyData))
	return nil
}
