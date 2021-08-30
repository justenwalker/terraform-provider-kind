package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"sigs.k8s.io/kind/pkg/cluster"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCluster(t *testing.T) {
	t.Cleanup(destroyTestCluster)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy: func(state *terraform.State) error {
			destroyTestCluster()
			return nil
		},
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					createTestCluster(t)
				},
				Config: testAccCheckClusterDataSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "%"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "ca_certificate_data"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "client_certificate_data"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "client_key_data"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "context"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "control_plane_containers.#"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "control_plane_containers.0"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "id"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "kubeconfig"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "kubeconfig_internal"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "name"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.#"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.0.%"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.0.ipv4_address"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.0.ipv6_address"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.0.name"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes.0.role"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "server"),
				),
			},
		},
	})
}

func createTestCluster(t *testing.T) {
	c := cluster.NewProvider(cluster.ProviderWithDocker())
	if err := c.Create(testClusterName, cluster.CreateWithKubeconfigPath(testClusterKubeConfigPath)); err != nil {
		t.Fatalf("could not set up cluster: %v", err)
	}
}

func destroyTestCluster() {
	c := cluster.NewProvider()
	_ = c.Delete(testClusterName, testClusterKubeConfigPath)
}

func testAccCheckClusterDataSource() string {
	return fmt.Sprintf(`
provider "kind" {
	provider   = "docker"
	kubeconfig = "%[2]s"
}
data "kind_cluster" "new" {
	name = "%[1]s"
}
	`, testClusterName, testClusterKubeConfigPath)
}
