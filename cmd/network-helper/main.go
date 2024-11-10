package main

import (
	"fmt"
	"os"

	ctlcni "github.com/cloudweav/cloudweav/pkg/generated/controllers/k8s.cni.cncf.io"
	"github.com/urfave/cli"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/manager/nad"
	"github.com/cloudweav/cloudweav-network-controller/pkg/helper"
	"github.com/cloudweav/cloudweav-network-controller/pkg/utils"
)

func main() {
	app := cli.NewApp()
	app.Name = "network-helper"
	app.Usage = "network-helper is to help get the network information through DHCP protocol from the pod within the VLAN network"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "kubeconfig, k",
			EnvVar: "KUBECONFIG",
			Value:  "",
			Usage:  "Kubernetes config files, e.g. $HOME/.kube/config",
		},
		cli.StringFlag{
			Name:   "master, m",
			EnvVar: "MASTERURL",
			Value:  "",
			Usage:  "Kubernetes cluster master URL.",
		},
		// example: [{"interface":"net1","name":"vlan178","namespace":"default"}]
		cli.StringFlag{
			Name:   "nadnetworks, n",
			EnvVar: nad.JobEnvNadNetwork,
			Value:  "",
			Usage:  "NAD network information",
		},
		cli.StringFlag{
			Name:   "dhcpserver, d",
			EnvVar: nad.JobEnvDHCPServer,
			Value:  "",
			Usage:  "DHCP server IP address",
		},
	}
	app.Action = func(c *cli.Context) {
		if err := run(c); err != nil {
			panic(err)
		}
	}

	if err := app.Run(os.Args); err != nil {
		klog.Error(err)
	}
}

func run(c *cli.Context) error {
	masterURL := c.String("master")
	kubeconfig := c.String("kubeconfig")
	networks := c.String("nadnetworks")
	dhcpServerIPAddr := c.String("dhcpserver")

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		return fmt.Errorf("error building config from flags: %w", err)
	}
	cni, err := ctlcni.NewFactoryFromConfig(cfg)
	if err != nil {
		return err
	}

	selectedNetworks, err := utils.NewNADSelectedNetworks(networks)
	if err != nil {
		return fmt.Errorf("failed to create nad selected network, networks: %s, error: %w", networks, err)
	}
	netHelper := helper.New(cni)

	for _, selectedNetwork := range selectedNetworks {
		networkConf := netHelper.GetVLANLayer3Network(&selectedNetwork, dhcpServerIPAddr)

		if err := netHelper.RecordToNad(&selectedNetwork, networkConf); err != nil {
			return fmt.Errorf("failed to record to nad cr, error: %w", err)
		}
	}

	return nil
}
