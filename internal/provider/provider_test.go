package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	testClusterName = "test"
	testClusterKubeConfigPath = "./testdata/kind-config"
)

var testAccProvider *schema.Provider

var testAccProviderFactories = map[string]func() (*schema.Provider, error){
	"kind": func() (*schema.Provider, error) {
		p := Provider()
		testAccProvider = p
		return p, nil
	},
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {

}
