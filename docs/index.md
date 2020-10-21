# Kind Provider

Terraform provider for managing [Kind](https://kind.sigs.k8s.io/) clusters (Kubernetes clusters using Docker).

## Example Usage

```hcl
# Initialize the kind provider using 'docker' backend
# and overriding which kubeconfig will be updated with
# configuration for connecting to the provisioned clusers
provider "kind" {
	provider   = "docker"
	kubeconfig = pathexpand("~/.kube/kind-config")
}

```

## Argument Reference

- `provider` - (Optional) The provider used to run the containers. Can be either `docker` or `podman` (Default: `docker`)
- `kubeconfig` - (Optional) Path to the kubeconfig to add/update contexts. (Default: `KUBECONFIG` env)
- `verbosity` - (Optional) Set logging verbosity (Default: `0`)