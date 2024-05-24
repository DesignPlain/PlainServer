package types

type Ec2_getLaunchTemplateInstanceRequirement struct {
	//
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	//
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	//
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	//
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	//
	BaselineEbsBandwidthMbps []Ec2_getLaunchTemplateInstanceRequirementBaselineEbsBandwidthMbp `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	//
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	//
	MemoryGibPerVcpus []Ec2_getLaunchTemplateInstanceRequirementMemoryGibPerVcpus `json:"memoryGibPerVcpus,omitempty" yaml:"memoryGibPerVcpus,omitempty"`

	//
	MemoryMibs []Ec2_getLaunchTemplateInstanceRequirementMemoryMib `json:"memoryMibs,omitempty" yaml:"memoryMibs,omitempty"`

	//
	NetworkInterfaceCounts []Ec2_getLaunchTemplateInstanceRequirementNetworkInterfaceCount `json:"networkInterfaceCounts,omitempty" yaml:"networkInterfaceCounts,omitempty"`

	//
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`

	//
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	//
	VcpuCounts []Ec2_getLaunchTemplateInstanceRequirementVcpuCount `json:"vcpuCounts,omitempty" yaml:"vcpuCounts,omitempty"`

	//
	AcceleratorCounts []Ec2_getLaunchTemplateInstanceRequirementAcceleratorCount `json:"acceleratorCounts,omitempty" yaml:"acceleratorCounts,omitempty"`

	//
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	//
	AcceleratorTotalMemoryMibs []Ec2_getLaunchTemplateInstanceRequirementAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMibs,omitempty" yaml:"acceleratorTotalMemoryMibs,omitempty"`

	//
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	//
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	//
	TotalLocalStorageGbs []Ec2_getLaunchTemplateInstanceRequirementTotalLocalStorageGb `json:"totalLocalStorageGbs,omitempty" yaml:"totalLocalStorageGbs,omitempty"`

	//
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	//
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	//
	NetworkBandwidthGbps []Ec2_getLaunchTemplateInstanceRequirementNetworkBandwidthGbp `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`

	//
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	//
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`
}
