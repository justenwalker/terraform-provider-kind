package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kubeconfig": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nodes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     nodeSchema(),
			},
		},
	}
}

func dataSourceClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	meta := m.(*Meta)
	name := data.Get("name").(string)
	diags = append(diags, setComputedResources(meta, name, data)...)
	data.SetId(meta.id(name))
	return
}
