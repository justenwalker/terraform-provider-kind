---
layout: ""
page_title: "Resource: Cluster"
description: |-
    Manage a Kind Cluster
---

# Kind Provider

This resource creates a cluster.

~> This resource does not support update. Any change will result in a cluster re-build.

## Example Usage

```terraform
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

## Schema

### Required

- **name** (String, Required) the name of the cluster. corresponds to the --name flag on the kind cli.

### Optional

- **config** (String, Optional) the cluster config as documented on https://kind.sigs.k8s.io/docs/user/configuration/
- **id** (String, Optional) The ID of this resource.
- **image** (String, Optional) The image to use for the kind nodes. corresponds to the --image flag on the cli.
- **image_version** (String, Optional) Kubernetes major.minor version, which chooses the correct node image from the published SHAs matching this version of KIND

### Read-only

- **ca_certificate_data** (String, Read-only) The base64-encoded CA Certificate used by the API Server
- **client_certificate_data** (String, Read-only) The base64-encoded client certificate data for connecting the cluster
- **client_key_data** (String, Read-only) The base64-encoded client private key data for connecting the cluster
- **context** (String, Read-only) The name of the context in KubeConfig
- **control_plane_containers** (List of String, Read-only) The list of control-plane node container names
- **kubeconfig** (String, Read-only) The full text of the kubeconfig that can be used to connect to this cluster
- **kubeconfig_internal** (String, Read-only) The full text of the kubeconfig that can be used to connect to this cluster from inside the container network
- **nodes** (List of Object, Read-only) The list of nodes that were provisioned for this cluster (see [below for nested schema](#nestedatt--nodes))
- **server** (String, Read-only) Kubernetes API Server URL

<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

- **ipv4_address** (String)
- **ipv6_address** (String)
- **name** (String)
- **role** (String)