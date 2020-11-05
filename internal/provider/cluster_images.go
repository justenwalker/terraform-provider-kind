package provider

// https://github.com/kubernetes-sigs/kind/releases/tag/v0.9.0
var clusterImage = map[string]string{
	"1.19": "kindest/node:v1.19.1@sha256:98cf5288864662e37115e362b23e4369c8c4a408f99cbc06e58ac30ddc721600",
	"1.18": "kindest/node:v1.18.8@sha256:f4bcc97a0ad6e7abaf3f643d890add7efe6ee4ab90baeb374b4f41a4c95567eb",
	"1.17": "kindest/node:v1.17.11@sha256:5240a7a2c34bf241afb54ac05669f8a46661912eab05705d660971eeb12f6555",
	"1.16": "kindest/node:v1.16.15@sha256:a89c771f7de234e6547d43695c7ab047809ffc71a0c3b65aa54eda051c45ed20",
	"1.15": "kindest/node:v1.15.12@sha256:d9b939055c1e852fe3d86955ee24976cab46cba518abcb8b13ba70917e6547a6",
	"1.14": "kindest/node:v1.14.10@sha256:ce4355398a704fca68006f8a29f37aafb49f8fc2f64ede3ccd0d9198da910146",
	"1.13": "kindest/node:v1.13.12@sha256:1c1a48c2bfcbae4d5f4fa4310b5ed10756facad0b7a2ca93c7a4b5bae5db29f5",
}
