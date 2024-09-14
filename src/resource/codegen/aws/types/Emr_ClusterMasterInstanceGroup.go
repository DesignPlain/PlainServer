package types

type Emr_ClusterMasterInstanceGroup struct {
	// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
	EbsConfigs []Emr_ClusterMasterInstanceGroupEbsConfig `json:"ebsConfigs,omitempty" yaml:"ebsConfigs,omitempty"`

	// Master node type Instance Group ID, if using Instance Group for this node type.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Target number of instances for the instance group. Must be 1 or 3. Defaults to 1. Launching with multiple master nodes is only supported in EMR version 5.23.0+, and requires this resource's `core_instance_group` to be configured. Public (Internet accessible) instances must be created in VPC subnets that have map public IP on launch enabled. Termination protection is automatically enabled when launched with multiple master nodes and this provider must have the `termination_protection = false` configuration applied before destroying this resource.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// EC2 instance type for all instances in the instance group.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Friendly name given to the instance group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice string `json:"bidPrice,omitempty" yaml:"bidPrice,omitempty"`
}
