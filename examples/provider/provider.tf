# Initialize the kind provider using 'docker' backend
# and overriding which kubeconfig will be updated with
# configuration for connecting to the provisioned clusers
provider "kind" {
  provider   = "docker"
  kubeconfig = pathexpand("~/.kube/kind-config")
}