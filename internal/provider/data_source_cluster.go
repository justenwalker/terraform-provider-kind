package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceClusterRead,
		Schema: setClusterAttributesSchema(map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		}),
	}
}

func dataSourceClusterRead(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	meta := m.(*Meta)
	name := data.Get("name").(string)
	exists, err := meta.clusterExists(name)
	if err != nil {
		return diag.FromErr(err)
	}
	if !exists {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("cluster %q not foud", name),
		})
		return
	}
	diags = append(diags, setClusterAttributeData(meta, name, data)...)
	data.SetId(meta.id(name))
	return
}
