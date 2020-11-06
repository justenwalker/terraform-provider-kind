---
layout: ""
page_title: "Provider: Kind"
description: |-
  Terraform provider for managing Kind clusters
---

# Kind Provider

Terraform provider for managing [Kind](https://kind.sigs.k8s.io/) clusters (Kubernetes clusters using Docker).

## Example Usage

```terraform
# Initialize the kind provider using 'docker' backend
# and overriding which kubeconfig will be updated with
# configuration for connecting to the provisioned clusers
provider "kind" {
  provider   = "docker"
  kubeconfig = pathexpand("~/.kube/kind-config")
}
```

## Schema

### Optional

- **http_proxy** (String, Optional) Override the HTTPS proxy used when provisioning kind clusters.
- **https_proxy** (String, Optional) Override the HTTP proxy used when provisioning kind clusters.
- **kubeconfig** (String, Optional) Sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config
- **no_proxy** (String, Optional) Override the NO_PROXY list used when provisioning kind clusters.
- **provider** (String, Optional) The provider used to run the containers. Can be either `docker` or `podman`
- **verbosity** (Number, Optional) Sets the logging verbosity. larger number means more logs.