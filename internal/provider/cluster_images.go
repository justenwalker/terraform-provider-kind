package provider

// https://github.com/kubernetes-sigs/kind/releases/tag/v0.17.0
var clusterImage = map[string]string{
	"1.25": "kindest/node:v1.25.3@sha256:f52781bc0d7a19fb6c405c2af83abfeb311f130707a0e219175677e366cc45d1",
	"1.24": "kindest/node:v1.24.7@sha256:577c630ce8e509131eab1aea12c022190978dd2f745aac5eb1fe65c0807eb315",
	"1.23": "kindest/node:v1.23.13@sha256:ef453bb7c79f0e3caba88d2067d4196f427794086a7d0df8df4f019d5e336b61",
	"1.22": "kindest/node:v1.22.15@sha256:7d9708c4b0873f0fe2e171e2b1b7f45ae89482617778c1c875f1053d4cef2e41",
	"1.21": "kindest/node:v1.21.14@sha256:9d9eb5fb26b4fbc0c6d95fa8c790414f9750dd583f5d7cee45d92e8c26670aa1",
	"1.20": "kindest/node:v1.20.15@sha256:a32bf55309294120616886b5338f95dd98a2f7231519c7dedcec32ba29699394",
	"1.19": "kindest/node:v1.19.16@sha256:476cb3269232888437b61deca013832fee41f9f074f9bed79f57e4280f7c48b7",
}
