package types

type Container_ClusterNodePoolNodeConfigLinuxNodeConfig struct {
	/*
	   The Linux kernel parameters to be applied to the nodes
	   and all pods running on the nodes. Specified as a map from the key, such as
	   `net.core.wmem_max`, to a string value. Currently supported attributes can be found [here](https://cloud.google.com/sdk/gcloud/reference/beta/container/node-pools/create#--system-config-from-file).
	   Note that validations happen all server side. All attributes are optional.
	*/
	Sysctls map[string]string `json:"sysctls,omitempty" yaml:"sysctls,omitempty"`

	/*
	   Possible cgroup modes that can be used.
	   Accepted values are:
	*/
	CgroupMode string `json:"cgroupMode,omitempty" yaml:"cgroupMode,omitempty"`
}
