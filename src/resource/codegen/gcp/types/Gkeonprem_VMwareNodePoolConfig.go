package types

type Gkeonprem_VMwareNodePoolConfig struct {
	// VMware disk size to be used during creation.
	BootDiskSizeGb int `json:"bootDiskSizeGb,omitempty" yaml:"bootDiskSizeGb,omitempty"`

	// The number of CPUs for each node in the node pool.
	Cpus int `json:"cpus,omitempty" yaml:"cpus,omitempty"`

	/*
	   Allow node pool traffic to be load balanced. Only works for clusters with
	   MetalLB load balancers.
	*/
	EnableLoadBalancer bool `json:"enableLoadBalancer,omitempty" yaml:"enableLoadBalancer,omitempty"`

	// The megabytes of memory for each node in the node pool.
	MemoryMb int `json:"memoryMb,omitempty" yaml:"memoryMb,omitempty"`

	// The number of nodes in the node pool.
	Replicas int `json:"replicas,omitempty" yaml:"replicas,omitempty"`

	// The OS image name in vCenter, only valid when using Windows.
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	/*
	   The OS image to be used for each node in a node pool.
	   Currently `cos`, `ubuntu`, `ubuntu_containerd` and `windows` are supported.
	*/
	ImageType string `json:"imageType,omitempty" yaml:"imageType,omitempty"`

	/*
	   The map of Kubernetes labels (key/value pairs) to be applied to each node.
	   These will added in addition to any default label(s) that
	   Kubernetes may apply to the node.
	   In case of conflict in label keys, the applied set may differ depending on
	   the Kubernetes version -- it's best to assume the behavior is undefined
	   and conflicts should be avoided.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The initial taints assigned to nodes of this node pool.
	   Structure is documented below.
	*/
	Taints []Gkeonprem_VMwareNodePoolConfigTaint `json:"taints,omitempty" yaml:"taints,omitempty"`

	/*
	   Specifies the vSphere config for node pool.
	   Structure is documented below.
	*/
	VsphereConfigs []Gkeonprem_VMwareNodePoolConfigVsphereConfig `json:"vsphereConfigs,omitempty" yaml:"vsphereConfigs,omitempty"`
}
