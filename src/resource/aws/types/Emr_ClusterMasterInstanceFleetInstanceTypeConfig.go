package types

type Emr_ClusterMasterInstanceFleetInstanceTypeConfig struct {
	// Bid price for each EC2 Spot instance type as defined by `instance_type`. Expressed in USD. If neither `bid_price` nor `bid_price_as_percentage_of_on_demand_price` is provided, `bid_price_as_percentage_of_on_demand_price` defaults to 100%!.(MISSING)
	BidPrice string `json:"bidPrice,omitempty" yaml:"bidPrice,omitempty"`

	// Bid price, as a percentage of On-Demand price, for each EC2 Spot instance as defined by `instance_type`. Expressed as a number (for example, 20 specifies 20%!)(MISSING). If neither `bid_price` nor `bid_price_as_percentage_of_on_demand_price` is provided, `bid_price_as_percentage_of_on_demand_price` defaults to 100%!.(MISSING)
	BidPriceAsPercentageOfOnDemandPrice float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" yaml:"bidPriceAsPercentageOfOnDemandPrice,omitempty"`

	// Configuration classification that applies when provisioning cluster instances, which can include configurations for applications and software that run on the cluster. List of `configuration` blocks.
	Configurations []Emr_ClusterMasterInstanceFleetInstanceTypeConfigConfiguration `json:"configurations,omitempty" yaml:"configurations,omitempty"`

	// Configuration block(s) for EBS volumes attached to each instance in the instance group. Detailed below.
	EbsConfigs []Emr_ClusterMasterInstanceFleetInstanceTypeConfigEbsConfig `json:"ebsConfigs,omitempty" yaml:"ebsConfigs,omitempty"`

	// EC2 instance type, such as m4.xlarge.
	InstanceType string `json:"instanceType,omitempty" yaml:"instanceType,omitempty"`

	// Number of units that a provisioned instance of this type provides toward fulfilling the target capacities defined in `aws.emr.InstanceFleet`.
	WeightedCapacity int `json:"weightedCapacity,omitempty" yaml:"weightedCapacity,omitempty"`
}
