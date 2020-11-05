package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func stringAttribute(desc string) *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeString,
		Description: desc,
		Computed:    true,
	}
}
