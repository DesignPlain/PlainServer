package types

type Ec2_getLaunchTemplateInstanceRequirement struct {
	//
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`

	//
	NetworkBandwidthGbps []Ec2_getLaunchTemplateInstanceRequirementNetworkBandwidthGbp `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`

	//
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	//
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	//
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	//
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	//
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	//
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	//
	MaxSpotPriceAsPercentageOfOptimalOnDemandPrice int `json:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty" yaml:"maxSpotPriceAsPercentageOfOptimalOnDemandPrice,omitempty"`

	//
	MemoryMibs []Ec2_getLaunchTemplateInstanceRequirementMemoryMib `json:"memoryMibs,omitempty" yaml:"memoryMibs,omitempty"`

	//
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	//
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	//
	TotalLocalStorageGbs []Ec2_getLaunchTemplateInstanceRequirementTotalLocalStorageGb `json:"totalLocalStorageGbs,omitempty" yaml:"totalLocalStorageGbs,omitempty"`

	//
	VcpuCounts []Ec2_getLaunchTemplateInstanceRequirementVcpuCount `json:"vcpuCounts,omitempty" yaml:"vcpuCounts,omitempty"`

	//
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`

	//
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	//
	NetworkInterfaceCounts []Ec2_getLaunchTemplateInstanceRequirementNetworkInterfaceCount `json:"networkInterfaceCounts,omitempty" yaml:"networkInterfaceCounts,omitempty"`

	//
	BaselineEbsBandwidthMbps []Ec2_getLaunchTemplateInstanceRequirementBaselineEbsBandwidthMbp `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	//
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	//
	MemoryGibPerVcpus []Ec2_getLaunchTemplateInstanceRequirementMemoryGibPerVcpus `json:"memoryGibPerVcpus,omitempty" yaml:"memoryGibPerVcpus,omitempty"`

	//
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	//
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	//
	AcceleratorCounts []Ec2_getLaunchTemplateInstanceRequirementAcceleratorCount `json:"acceleratorCounts,omitempty" yaml:"acceleratorCounts,omitempty"`

	//
	AcceleratorTotalMemoryMibs []Ec2_getLaunchTemplateInstanceRequirementAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMibs,omitempty" yaml:"acceleratorTotalMemoryMibs,omitempty"`
}
