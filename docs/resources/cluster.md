# Cluster Resource

This resource creates a cluster.

~> This resource does not support update. Any change will result in a cluster re-build.

## Example Usage

```hcl
# Create a test cluster named "test" with a control-plane and
# two workers using the inline config argument
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
```

## Argument Reference

* `name` - (Required) The name of the cluster.
* `config` - (Optional) [cluster config](https://www.terraform.io/docs/registry/providers/docs.html) used for customizing cluster creation.
* `image` - (Optional) The image to use for the kind nodes. corresponds to the `--image` flag on the cli.

## Attribute Reference

* `kubeconfig` - yaml string contianing the kubeconfig connection details for connecting to this cluster
* `nodes` - A list of cluster nodes. See the [Nodes](#nodes) section

## Nodes

* `name` - Name of the cluster node. Usually corresponds with the container name (in docker/podman)
* `role` - The node's role, `worker` or `control-plane`
* `ipv4_address` - The internal ipv4 address of the node.
* `ipv6_address` - The internal ipv6 address of the node.

