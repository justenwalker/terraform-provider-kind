package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/constants"
	"sort"
	"strings"
)

type Meta struct {
	ProviderName   string
	KubeConfigPath string
	Provider       *cluster.Provider
}

func idToName(id string) string {
	ss := strings.SplitN(id, "/", 2)
	if len(ss) > 1 {
		return ss[1]
	}
	return id
}

func (m *Meta) clusterExists(name string) (bool, error) {
	clusters, err := m.Provider.List()
	if err != nil {
		return false, err
	}
	for _, c := range clusters {
		if c == name {
			return true, nil
		}
	}
	return false, nil
}

func (m *Meta) id(name string) string {
	return fmt.Sprintf("%s/%s", m.ProviderName, name)
}

func (m *Meta) deleteCluster(name string) error {
	exists, err := m.clusterExists(name)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return m.Provider.Delete(name, m.KubeConfigPath)
}

func (m *Meta) setKindNodes(name string, data *schema.ResourceData) error {
	ns, err := m.Provider.ListNodes(name)
	if err != nil {
		return err
	}
	// so we get a predictable order
	sort.Slice(ns, func(i, j int) bool {
		return strings.Compare(ns[i].String(), ns[j].String()) < 0
	})
	nodes, err := mapKindNodeList(ns)
	if err != nil {
		return err
	}
	var cps []string
	for _, node := range nodes {
		if role := node["role"].(string); role == constants.ControlPlaneNodeRoleValue {
			name := node["name"].(string)
			cps = append(cps, name)
		}
	}
	_ = data.Set("nodes", nodes)
	_ = data.Set("control_plane_containers", cps)
	return nil
}
