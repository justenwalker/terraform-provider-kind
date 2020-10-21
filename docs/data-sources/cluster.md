# Cluster Data Source

This data source provides kind cluster details such as nodes and kubconfig.

## Example Usage

```hcl
# Get attributes of an existing cluster named 'kind'
data "kind_cluster" "cluster" {
	name = "kind"
}
```

## Argument Reference

* `name` - (Required) The name of the cluster.

## Attribute Reference

* `kubeconfig` - yaml string contianing the kubeconfig connection details for connecting to this cluster
* `nodes` - A list of cluster nodes. See the [Nodes](#nodes) section

## Nodes

* `name` - Name of the cluster node. Usually corresponds with the container name (in docker/podman)
* `role` - The node's role, `worker` or `control-plane`
* `ipv4_address` - The internal ipv4 address of the node.
* `ipv6_address` - The internal ipv6 address of the node.

