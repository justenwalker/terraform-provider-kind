package provider

import (
	"fmt"
	"sigs.k8s.io/kind/pkg/cluster"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceCluster(t *testing.T) {
	t.Cleanup(destroyTestCluster)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			{
				PreConfig: func() {
					createTestCluster(t)
				},
				Config: testAccCheckClusterDataSource(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "name"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "kubeconfig"),
					resource.TestCheckResourceAttrSet("data.kind_cluster.new", "nodes"),
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
