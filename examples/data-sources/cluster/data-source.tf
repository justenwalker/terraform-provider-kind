# Get attributes of an existing cluster named 'kind'
data "kind_cluster" "cluster" {
  name = "kind"
}