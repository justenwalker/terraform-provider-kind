---
layout: ""
page_title: "Data Source: Cluster"
description: |-
    Get information about a Kind Cluster
---

# Kind Provider

Get information about a Kind Cluster

## Example Usage

```terraform
# Get attributes of an existing cluster named 'kind'
data "kind_cluster" "cluster" {
  name = "kind"
}
```

## Schema

### Required

- **name** (String, Required)

### Optional

- **id** (String, Optional) The ID of this resource.

### Read-only

- **ca_certificate_data** (String, Read-only) The base64-encoded CA Certificate used by the API Server
- **client_certificate_data** (String, Read-only) The base64-encoded client certificate data for connecting the cluster
- **client_key_data** (String, Read-only) The base64-encoded client private key data for connecting the cluster
- **context** (String, Read-only) The name of the context in KubeConfig
- **kubeconfig** (String, Read-only) The full text of the kubeconfig that can be used to connect to this cluster
- **nodes** (List of Object, Read-only) The list of nodes that were provisioned for this cluster (see [below for nested schema](#nestedatt--nodes))
- **server** (String, Read-only) Kubernetes API Server URL

<a id="nestedatt--nodes"></a>
### Nested Schema for `nodes`

- **ipv4_address** (String)
- **ipv6_address** (String)
- **name** (String)
- **role** (String)