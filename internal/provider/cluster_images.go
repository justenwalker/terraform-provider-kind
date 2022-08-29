package provider

// https://github.com/kubernetes-sigs/kind/releases/tag/v0.14.0
var clusterImage = map[string]string{
	"1.24": "kindest/node:v1.24.0@sha256:0866296e693efe1fed79d5e6c7af8df71fc73ae45e3679af05342239cdc5bc8e",
	"1.23": "kindest/node:v1.23.6@sha256:b1fa224cc6c7ff32455e0b1fd9cbfd3d3bc87ecaa8fcb06961ed1afb3db0f9ae",
	"1.22": "kindest/node:v1.22.9@sha256:8135260b959dfe320206eb36b3aeda9cffcb262f4b44cda6b33f7bb73f453105",
	"1.21": "kindest/node:v1.21.12@sha256:f316b33dd88f8196379f38feb80545ef3ed44d9197dca1bfd48bcb1583210207",
	"1.20": "kindest/node:v1.20.15@sha256:6f2d011dffe182bad80b85f6c00e8ca9d86b5b8922cdf433d53575c4c5212248",
	"1.19": "kindest/node:v1.19.16@sha256:d9c819e8668de8d5030708e484a9fdff44d95ec4675d136ef0a0a584e587f65c",
	"1.18": "kindest/node:v1.18.20@sha256:738cdc23ed4be6cc0b7ea277a2ebcc454c8373d7d8fb991a7fcdbd126188e6d7",
}
