package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/kind/pkg/cluster"
	"strings"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceClusterRead,
		CreateContext: resourceClusterCreate,
		DeleteContext: resourceClusterDelete,
		CustomizeDiff: resourceClusterDiff,
		Schema: map[string]*schema.Schema{
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
				Type:        schema.TypeString,
				Description: `The image to use for the kind nodes. corresponds to the --image flag on the cli.`,
				Optional:    true,
				ForceNew:    true,
			},
			"kubeconfig": {
				Type:        schema.TypeString,
				Description: `The full text of the kubeconfig that can be used to connect to this cluster`,
				Computed:    true,
			},
			"nodes": {
				Type:        schema.TypeList,
				Description: `The list of nodes that were provisioned for this cluster`,
				Computed:    true,
				Elem:        nodeSchema(),
			},
		},
	}
}

func resourceClusterDelete(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	meta := m.(*Meta)
	name := idToName(data.Id())
	if err := meta.deleteCluster(name); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "failed to delete cluster",
			Detail:   err.Error(),
		})
		return
	}
	return
}

func resourceClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	meta := m.(*Meta)
	name := idToName(data.Id())
	exists, err := meta.clusterExists(name)
	if err != nil {
		return diag.FromErr(err)
	}
	if !exists {
		return
	}
	setComputedResources(meta, name, data)
	return
}

func resourceClusterCreate(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	meta := m.(*Meta)
	name := data.Get("name").(string)
	copts := []cluster.CreateOption{
		cluster.CreateWithDisplayUsage(false),
		cluster.CreateWithDisplaySalutation(false),
	}
	if meta.KubeConfigPath != "" {
		copts = append(copts, cluster.CreateWithKubeconfigPath(meta.KubeConfigPath))
	}
	if v, ok := data.GetOk("config"); ok {
		config := strings.TrimSpace(v.(string))
		if len(config) > 0 {
			copts = append(copts, cluster.CreateWithRawConfig([]byte(config)))
		}
	}
	if v, ok := data.GetOk("image"); ok {
		image := strings.TrimSpace(v.(string))
		copts = append(copts, cluster.CreateWithNodeImage(image))
	}
	if err := meta.Provider.Create(name, copts...); err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("failed to create cluster %q", name),
			Detail:   err.Error(),
		})
		return
	}
	diags = append(diags, setComputedResources(meta, name, data)...)
	data.SetId(meta.id(name))
	return
}

func setComputedResources(c *Meta, name string, data *schema.ResourceData) (diags diag.Diagnostics) {
	kubeconfig, err := c.getKubeConfig(name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("could not export kubeconfig for cluster %q", name),
			Detail:   err.Error(),
		})
	}
	_ = data.Set("kubeconfig", kubeconfig)
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

func resourceClusterDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
