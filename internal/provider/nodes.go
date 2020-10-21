package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
)

func nodeSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: `Name of the cluster node. Usually corresponds with the container name (in docker/podman)`,
				Computed:    true,
			},
			"role": {
				Type:        schema.TypeString,
				Description: "The node's role, `worker` or `control-plane`",
				Computed:    true,
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Description: `The internal ipv4 address of the node`,
				Computed:    true,
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Description: `The internal ipv6 address of the node`,
				Computed:    true,
			},
		},
	}
}

func mapKindNodeList(nodes []nodes.Node) (result []map[string]interface{}, err error) {
	for _, node := range nodes {
		n, err := mapKindNode(node)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}
	return
}

func mapKindNode(node nodes.Node) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	node.String()
	m["name"] = node.String()
	role, err := node.Role()
	if err != nil {
		return nil, err
	}
	m["role"] = role
	ipv4, ipv6, err := node.IP()
	if err != nil {
		return nil, err
	}
	if ipv4 != "" {
		m["ipv4_address"] = ipv4
	}
	if ipv6 != "" {
		m["ipv6_address"] = ipv6
	}
	return m, nil
}
