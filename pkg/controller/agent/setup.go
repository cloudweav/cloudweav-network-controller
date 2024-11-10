package agent

import (
	"github.com/cloudweav/cloudweav-network-controller/pkg/config"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/agent/linkmonitor"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/agent/nad"
	"github.com/cloudweav/cloudweav-network-controller/pkg/controller/agent/vlanconfig"
)

var RegisterFuncList = []config.RegisterFunc{
	nad.Register,
	vlanconfig.Register,
	linkmonitor.Register,
}
