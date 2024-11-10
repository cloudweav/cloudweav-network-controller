package main

import (
	"os"

	controllergen "github.com/rancher/wrangler/pkg/controller-gen"
	"github.com/rancher/wrangler/pkg/controller-gen/args"

	networkv1 "github.com/cloudweav/cloudweav-network-controller/pkg/apis/network.cloudweavhci.io/v1beta1"
)

func main() {
	os.Unsetenv("GOPATH")
	controllergen.Run(args.Options{
		OutputPackage: "github.com/cloudweav/cloudweav-network-controller/pkg/generated",
		Boilerplate:   "hack/boilerplate.go.txt",
		Groups: map[string]args.Group{
			"network.cloudweavhci.io": {
				Types: []interface{}{
					networkv1.ClusterNetwork{},
					networkv1.NodeNetwork{},
					networkv1.VlanConfig{},
					networkv1.VlanStatus{},
					networkv1.LinkMonitor{},
				},
				GenerateTypes:   true,
				GenerateClients: true,
			},
		},
	})
}
