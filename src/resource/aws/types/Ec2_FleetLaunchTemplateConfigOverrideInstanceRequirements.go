package types

type Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirements struct {
	// Indicates whether burstable performance T instance types are `included`, `excluded`, or `required`. Default is `excluded`.
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	/*
	   The price protection threshold for On-Demand Instances. This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 20.

	   If you set `target_capacity_unit_type` to `vcpu` or `memory-mib`, the price protection threshold is applied based on the per-vCPU or per-memory price instead of the per-instance price.
	*/
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// Block describing the minimum and maximum number of vCPUs. Default is no maximum.
	VcpuCount Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsVcpuCount `json:"vcpuCount,omitempty" yaml:"vcpuCount,omitempty"`

	// Block describing the minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips). Default is no minimum or maximum limits.
	AcceleratorCount Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorCount `json:"acceleratorCount,omitempty" yaml:"acceleratorCount,omitempty"`

	// Indicate whether instance types with local storage volumes are `included`, `excluded`, or `required`. Default is `included`.
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	// The minimum and maximum amount of network bandwidth, in gigabits per second (Gbps). Default is No minimum or maximum.
	NetworkBandwidthGbps Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkBandwidthGbps `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`

	// Block describing the minimum and maximum total local storage (GB). Default is no minimum or maximum.
	TotalLocalStorageGb Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsTotalLocalStorageGb `json:"totalLocalStorageGb,omitempty" yaml:"totalLocalStorageGb,omitempty"`

	// Block describing the minimum and maximum baseline EBS bandwidth, in Mbps. Default is no minimum or maximum.
	BaselineEbsBandwidthMbps Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsBaselineEbsBandwidthMbps `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	/*
	   The CPU manufacturers to include. Default is any manufacturer.
	   > --NOTE:-- Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.
	*/
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	/*
	   The instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (\-). The following are examples: `c5-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are excluding the entire C5 instance family, which includes all C5a and C5n instance  If you specify `m5a.-`, you are excluding all the M5a instance types, but not the M5n instance  Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance

	   If you specify `AllowedInstanceTypes`, you can't specify `ExcludedInstanceTypes`.
	*/
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	/*
	   The price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 100.

	   If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	*/
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	// Block describing the minimum and maximum total memory of the accelerators. Default is no minimum or maximum.
	AcceleratorTotalMemoryMib Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMib,omitempty" yaml:"acceleratorTotalMemoryMib,omitempty"`

	// The accelerator types that must be on the instance type. Default is any accelerator type.
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	/*
	   The instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes. You can use strings with one or more wild cards,represented by an asterisk (\-). The following are examples: `c5-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are excluding the entire C5 instance family, which includes all C5a and C5n instance  If you specify `m5a.-`, you are excluding all the M5a instance types, but not the M5n instance  Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance  Default is any instance type.

	   If you specify `AllowedInstanceTypes`, you can't specify `ExcludedInstanceTypes`.
	*/
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	// Indicate whether bare metal instace types should be `included`, `excluded`, or `required`. Default is `excluded`.
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	// Block describing the minimum and maximum amount of memory (GiB) per vCPU. Default is no minimum or maximum.
	MemoryGibPerVcpu Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryGibPerVcpu `json:"memoryGibPerVcpu,omitempty" yaml:"memoryGibPerVcpu,omitempty"`

	// The minimum and maximum amount of memory per vCPU, in GiB. Default is no minimum or maximum limits.
	MemoryMib Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsMemoryMib `json:"memoryMib,omitempty" yaml:"memoryMib,omitempty"`

	// Block describing the minimum and maximum number of network interfaces. Default is no minimum or maximum.
	NetworkInterfaceCount Ec2_FleetLaunchTemplateConfigOverrideInstanceRequirementsNetworkInterfaceCount `json:"networkInterfaceCount,omitempty" yaml:"networkInterfaceCount,omitempty"`

	// Indicate whether instance types must support On-Demand Instance Hibernation, either `true` or `false`. Default is `false`.
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`

	// List of accelerator manufacturer names. Default is any manufacturer.
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	// List of accelerator names. Default is any acclerator.
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	// Indicates whether current or previous generation instance types are included. The current generation instance types are recommended for use. Valid values are `current` and `previous`. Default is `current` and `previous` generation instance
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	// List of local storage type names. Valid values are `hdd` and `ssd`. Default any storage type.
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`
}
