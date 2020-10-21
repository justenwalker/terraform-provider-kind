package provider

import (
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccCluster(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccCheckClusterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccClusterConfig(t, ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.0.name","test-control-plane"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.0.role","control-plane"),
				),
			},
			{
				Config: testAccClusterConfig(t, `
apiVersion: kind.x-k8s.io/v1alpha4
kind: Cluster
nodes:
- role: control-plane
- role: worker
- role: worker
`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("kind_cluster.new","name"),
					resource.TestCheckResourceAttrSet("kind_cluster.new","kubeconfig"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.0.name","test-control-plane"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.0.role","control-plane"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.1.name","test-worker"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.1.role","worker"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.2.name","test-worker2"),
					resource.TestCheckResourceAttr("kind_cluster.new","nodes.2.role","worker"),
				),
			},
		},
	})
}

func testAccCheckClusterDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*Meta)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cluster" {
			continue
		}

		name := rs.Primary.ID

		err := c.Provider.Delete(name, c.KubeConfigPath)
		if err != nil {
			return err
		}
	}

	return nil
}

var testTemplate = template.Must(template.New("tpl").Parse(`
provider "kind" {
	provider = "docker"
	kubeconfig = "{{ .KubeConfig }}"
	verbosity = 0
}
resource "kind_cluster" "new" {
	name = "{{ .Name }}"
{{- if .Config }}
	config = <<EOT
{{ .Config }}
EOT
{{- end }}
}
`))

type testTemplateData struct {
	KubeConfig string
	Name       string
	Config     string
}

func testAccClusterConfig(t *testing.T, config string) string {
	var sb strings.Builder
	err := testTemplate.Execute(&sb, testTemplateData{
		KubeConfig: testClusterKubeConfigPath,
		Name:       testClusterName,
		Config:     config,
	})
	if err != nil {
		t.Fatalf("error rendering resource template: %v", err)
		return ""
	}
	return sb.String()
}
