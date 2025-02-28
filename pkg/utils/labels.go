package utils

import "github.com/cloudweav/cloudweav-network-controller/pkg/apis/network.cloudweavhci.io"

const (
	KeyVlanLabel = network.GroupName + "/vlan-id"
	// KeyLastVlanLabel is used to record the last VLAN id to support changing the VLAN id of the VLAN networks
	KeyLastVlanLabel       = network.GroupName + "/last-vlan-id"
	KeyVlanConfigLabel     = network.GroupName + "/vlanconfig"
	KeyClusterNetworkLabel = network.GroupName + "/clusternetwork"
	// KeyLastClusterNetworkLabel is used to record the last cluster network to support changing the cluster network of NADs
	KeyLastClusterNetworkLabel = network.GroupName + "/last-clusternetwork"
	KeyNodeLabel               = network.GroupName + "/node"
	KeyNetworkType             = network.GroupName + "/type"
	KeyLastNetworkType         = network.GroupName + "/last-type"
	KeyNetworkReady            = network.GroupName + "/ready"
	KeyNetworkRoute            = network.GroupName + "/route"

	KeyMatchedNodes = network.GroupName + "/matched-nodes"

	ValueTrue  = "true"
	ValueFalse = "false"
)
