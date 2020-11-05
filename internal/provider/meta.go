package provider

import (
	"fmt"
	"sigs.k8s.io/kind/pkg/cluster"
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

func (m *Meta) getKubeConfig(name string) (string, error) {
	return m.Provider.KubeConfig(name, false)
}

func (m *Meta) getKindNodeList(name string) ([]map[string]interface{}, error) {
	nodes, err := m.Provider.ListNodes(name)
	if err != nil {
		return nil, err
	}
	// so we get a predictable order
	sort.Slice(nodes, func(i, j int) bool {
		return strings.Compare(nodes[i].String(), nodes[j].String()) < 0
	})
	return mapKindNodeList(nodes)
}
