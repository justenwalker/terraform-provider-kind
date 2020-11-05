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
		Schema:        resourceClusterSchema(),
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
	setClusterAttributeData(meta, name, data)
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
	} else if v, ok := data.GetOk("image_version"); ok {
		version := strings.TrimSpace(v.(string))
		image, ok := clusterImage[version]
		if !ok {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("no pre-defined node image for kubernetes version %s", version),
			})
			return
		}
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
	diags = append(diags, setClusterAttributeData(meta, name, data)...)
	data.SetId(meta.id(name))
	return
}

func resourceClusterDiff(ctx context.Context, d *schema.ResourceDiff, meta interface{}) error {
	return nil
}
