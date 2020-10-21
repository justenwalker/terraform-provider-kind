terraform {
	required_providers {
		kind = {
			versions = ["0.1"]
			source = "registry.terraform.io/justenwalker/kind"
		}
	}
}

# Initialize the kind provider using 'docker' backend
# and overriding which kubeconfig will be updated with
# configuration for connecting to the provisioned clusers
provider "kind" {
	provider   = "docker"
	kubeconfig = pathexpand("~/.kube/kind-config")
}

# Get attributes of an existing cluster named 'kind'
data "kind_cluster" "cluster" {
	name = "kind"
}

# Create a test cluster named "test" with a control-plane and
# two workers using the custom config option
resource "kind_cluster" "new" {
	name = "test"
	config = <<-EOF
        apiVersion: kind.x-k8s.io/v1alpha4
        kind: Cluster
        nodes:
        - role: control-plane
        - role: worker
        - role: worker
	EOF
}