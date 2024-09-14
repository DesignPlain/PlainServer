package lightsail

import types "libds/aws/types"

type InstancePublicPorts struct {
	// Name of the Lightsail Instance.
	InstanceName string `json:"instanceName,omitempty" yaml:"instanceName,omitempty"`

	// Configuration block with port information. AWS closes all currently open ports that are not included in the `port_info`. Detailed below.
	PortInfos []types.Lightsail_InstancePublicPortsPortInfo `json:"portInfos,omitempty" yaml:"portInfos,omitempty"`
}
