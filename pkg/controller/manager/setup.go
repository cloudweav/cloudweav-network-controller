package manager

import (
	"github.com/cloudweav/cloudweav-network-controller/pkg/config"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/manager/clusternetwork"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/manager/nad"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/manager/node"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/manager/vlanconfig"
)

var RegisterFuncList = []config.RegisterFunc{
	nad.Register,
	vlanconfig.Register,
	node.Register,
	clusternetwork.Register,
}
