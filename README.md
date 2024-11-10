Cloudweav Network Controller
========
[![Build Status](https://drone-publish.rancher.io/api/badges/cloudweav/network-controller-cloudweav/status.svg)](https://drone-publish.rancher.io/cloudweav/network-controller-cloudweav)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudweav/network-controller-cloudweav)](https://goreportcard.com/report/github.com/cloudweav/network-controller-cloudweav)
[![Releases](https://img.shields.io/github/release/cloudweav/network-controller-cloudweav/all.svg)](https://github.com/cloudweav/network-controller-cloudweav/releases)

A network controller helps to manage the host network configuration of the [Cloudweav](https://github.com/cloudweav/cloudweav) cluster.

## Manifests and Deploying
The `./manifests` folder contains useful YAML manifests to use for deploying and developing the Cloudweav network controller. 
This simply YAML deployment creates a Daemonset using the `rancher/cloudweav-network-controller` container.

## License
Copyright (c) 2020 [Rancher Labs, Inc.](http://rancher.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.