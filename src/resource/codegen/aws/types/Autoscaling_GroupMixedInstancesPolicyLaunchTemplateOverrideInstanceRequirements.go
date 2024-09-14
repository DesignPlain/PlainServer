package types

type Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirements struct {
	// Block describing the minimum and maximum number of vCPUs. Default is no maximum.
	VcpuCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsVcpuCount `json:"vcpuCount,omitempty" yaml:"vcpuCount,omitempty"`

	// Indicate whether burstable performance instance types should be `included`, `excluded`, or `required`. Default is `excluded`.
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	/*
	   List of CPU manufacturer names. Default is any manufacturer.

	   > --NOTE:-- Don't confuse the CPU hardware manufacturer with the CPU hardware architecture. Instances will be launched with a compatible CPU architecture based on the Amazon Machine Image (AMI) that you specify in your launch template.

	   ```
	   Valid names:
	   - amazon-web-services
	   - amd
	   - intel
	   ```
	*/
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	/*
	   List of instance types to exclude. You can use strings with one or more wild cards, represented by an asterisk (\-), to exclude an instance type, size, or generation. The following are examples: `m5.8xlarge`, `c5-.-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are excluding the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.-`, you are excluding all the M5a instance types, but not the M5n instance types. Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is no excluded instance types.

	   > --NOTE:-- If you specify `excluded_instance_types`, you can't specify `allowed_instance_types`.
	*/
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	// Block describing the minimum and maximum amount of memory (MiB). Default is no maximum.
	MemoryMib Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryMib `json:"memoryMib,omitempty" yaml:"memoryMib,omitempty"`

	// Indicate whether instance types must support On-Demand Instance Hibernation, either `true` or `false`. Default is `false`.
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`

	// Block describing the minimum and maximum total local storage (GB). Default is no minimum or maximum.
	TotalLocalStorageGb Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsTotalLocalStorageGb `json:"totalLocalStorageGb,omitempty" yaml:"totalLocalStorageGb,omitempty"`

	/*
	   List of accelerator manufacturer names. Default is any manufacturer.

	   ```
	   Valid names:
	   - amazon-web-services
	   - amd
	   - nvidia
	   - xilinx
	   ```
	*/
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	/*
	   List of accelerator names. Default is any acclerator.

	   ```
	   Valid names:
	   - a100            - NVIDIA A100 GPUs
	   - v100            - NVIDIA V100 GPUs
	   - k80             - NVIDIA K80 GPUs
	   - t4              - NVIDIA T4 GPUs
	   - m60             - NVIDIA M60 GPUs
	   - radeon-pro-v520 - AMD Radeon Pro V520 GPUs
	   - vu9p            - Xilinx VU9P FPGAs
	   ```
	*/
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	/*
	   List of accelerator types. Default is any accelerator type.

	   ```
	   Valid types:
	   - fpga
	   - gpu
	   - inference
	   ```
	*/
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	// Block describing the minimum and maximum baseline EBS bandwidth, in Mbps. Default is no minimum or maximum.
	BaselineEbsBandwidthMbps Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsBaselineEbsBandwidthMbps `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	// Block describing the minimum and maximum number of network interfaces. Default is no minimum or maximum.
	NetworkInterfaceCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkInterfaceCount `json:"networkInterfaceCount,omitempty" yaml:"networkInterfaceCount,omitempty"`

	/*
	   Price protection threshold for On-Demand Instances. This is the maximum you’ll pay for an On-Demand Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 20.

	   If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	*/
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// Block describing the minimum and maximum number of accelerators (GPUs, FPGAs, or AWS Inferentia chips). Default is no minimum or maximum.
	AcceleratorCount Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorCount `json:"acceleratorCount,omitempty" yaml:"acceleratorCount,omitempty"`

	/*
	   List of instance types to apply your specified attributes against. All other instance types are ignored, even if they match your specified attributes. You can use strings with one or more wild cards, represented by an asterisk (\-), to allow an instance type, size, or generation. The following are examples: `m5.8xlarge`, `c5-.-`, `m5a.-`, `r-`, `-3-`. For example, if you specify `c5-`, you are allowing the entire C5 instance family, which includes all C5a and C5n instance types. If you specify `m5a.-`, you are allowing all the M5a instance types, but not the M5n instance types. Maximum of 400 entries in the list; each entry is limited to 30 characters. Default is all instance types.

	   > --NOTE:-- If you specify `allowed_instance_types`, you can't specify `excluded_instance_types`.
	*/
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	// Indicate whether instance types with local storage volumes are `included`, `excluded`, or `required`. Default is `included`.
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	// The price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Conflicts with `spot_max_price_percentage_over_lowest_price`
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice int `json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty"`

	/*
	   Price protection threshold for Spot Instances. This is the maximum you’ll pay for a Spot Instance, expressed as a percentage higher than the cheapest M, C, or R instance type with your specified attributes. When Amazon EC2 Auto Scaling selects instance types with your attributes, we will exclude instance types whose price is higher than your threshold. The parameter accepts an integer, which Amazon EC2 Auto Scaling interprets as a percentage. To turn off price protection, specify a high value, such as 999999. Default is 100. Conflicts with `max_spot_price_as_percentage_of_optimal_on_demand_price`

	   If you set DesiredCapacityType to vcpu or memory-mib, the price protection threshold is applied based on the per vCPU or per memory price instead of the per instance price.
	*/
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	// Block describing the minimum and maximum total memory of the accelerators. Default is no minimum or maximum.
	AcceleratorTotalMemoryMib Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMib,omitempty" yaml:"acceleratorTotalMemoryMib,omitempty"`

	// Indicate whether bare metal instace types should be `included`, `excluded`, or `required`. Default is `excluded`.
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	/*
	   List of instance generation names. Default is any generation.

	   ```
	   Valid names:
	   - current  - Recommended for best performance.
	   - previous - For existing applications optimized for older instance types.
	   ```
	*/
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	/*
	   List of local storage type names. Default any storage type.

	   ```
	   Value names:
	   - hdd - hard disk drive
	   - ssd - solid state drive
	   ```
	*/
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`

	// Block describing the minimum and maximum amount of memory (GiB) per vCPU. Default is no minimum or maximum.
	MemoryGibPerVcpu Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsMemoryGibPerVcpu `json:"memoryGibPerVcpu,omitempty" yaml:"memoryGibPerVcpu,omitempty"`

	// Block describing the minimum and maximum amount of network bandwidth, in gigabits per second (Gbps). Default is no minimum or maximum.
	NetworkBandwidthGbps Autoscaling_GroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementsNetworkBandwidthGbps `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`
}
