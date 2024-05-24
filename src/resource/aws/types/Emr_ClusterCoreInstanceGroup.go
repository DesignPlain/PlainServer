package types

type Emr_ClusterCoreInstanceGroup struct {
	// ID of the cluster.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Target number of instances for the instance group. Must be at least 1. Defaults to 1.
	InstanceCount int `json:"instanceCount,omitempty" yaml:"instanceCount,omitempty"`

	// EC2 instance type for all instances in the instance group.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Friendly name given to the instance group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// String containing the [EMR Auto Scaling Policy](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-automatic-scaling.html) JSON.
	AutoscalingPolicy string `json:"autoscalingPolicy,omitempty" yaml:"autoscalingPolicy,omitempty"`

	// Bid price for each EC2 instance in the instance group, expressed in USD. By setting this attribute, the instance group is being declared as a Spot Instance, and will implicitly create a Spot request. Leave this blank to use On-Demand Instances.
	BidPrice string `json:"bidPrice,omitempty" yaml:"bidPrice,omitempty"`

	// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
	EbsConfigs []Emr_ClusterCoreInstanceGroupEbsConfig `json:"ebsConfigs,omitempty" yaml:"ebsConfigs,omitempty"`
}
