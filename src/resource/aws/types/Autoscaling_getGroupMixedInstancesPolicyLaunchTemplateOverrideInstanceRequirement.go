package types

type Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirement struct {
	// List of excluded instance
	ExcludedInstanceTypes []string `json:"excludedInstanceTypes,omitempty" yaml:"excludedInstanceTypes,omitempty"`

	// List of instance types to apply the specified attributes against.
	AllowedInstanceTypes []string `json:"allowedInstanceTypes,omitempty" yaml:"allowedInstanceTypes,omitempty"`

	// Indicates whether bare metal instances are included, excluded, or required.
	BareMetal string `json:"bareMetal,omitempty" yaml:"bareMetal,omitempty"`

	// List of CPU manufacturer names.
	CpuManufacturers []string `json:"cpuManufacturers,omitempty" yaml:"cpuManufacturers,omitempty"`

	// Indicates whether instance types must support On-Demand Instance Hibernation.
	RequireHibernateSupport bool `json:"requireHibernateSupport,omitempty" yaml:"requireHibernateSupport,omitempty"`

	// Indicates whether burstable performance instance types are included, excluded, or required.
	BurstablePerformance string `json:"burstablePerformance,omitempty" yaml:"burstablePerformance,omitempty"`

	// List of local storage type names.
	LocalStorageTypes []string `json:"localStorageTypes,omitempty" yaml:"localStorageTypes,omitempty"`

	// Price protection threshold for On-Demand Instances.
	OnDemandMaxPricePercentageOverLowestPrice int `json:"onDemandMaxPricePercentageOverLowestPrice,omitempty" yaml:"onDemandMaxPricePercentageOverLowestPrice,omitempty"`

	// List of objects describing the minimum and maximum total memory of the accelerators.
	AcceleratorTotalMemoryMibs []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementAcceleratorTotalMemoryMib `json:"acceleratorTotalMemoryMibs,omitempty" yaml:"acceleratorTotalMemoryMibs,omitempty"`

	// List of objects describing the minimum and maximum baseline EBS bandwidth (Mbps).
	BaselineEbsBandwidthMbps []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementBaselineEbsBandwidthMbp `json:"baselineEbsBandwidthMbps,omitempty" yaml:"baselineEbsBandwidthMbps,omitempty"`

	// Indicates whether instance types with instance store volumes are included, excluded, or required.
	LocalStorage string `json:"localStorage,omitempty" yaml:"localStorage,omitempty"`

	// List of objects describing the minimum and maximum amount of memory (MiB).
	MemoryMibs []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementMemoryMib `json:"memoryMibs,omitempty" yaml:"memoryMibs,omitempty"`

	// List of objects describing the minimum and maximum amount of network interfaces.
	NetworkInterfaceCounts []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementNetworkInterfaceCount `json:"networkInterfaceCounts,omitempty" yaml:"networkInterfaceCounts,omitempty"`

	//
	AcceleratorCounts []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementAcceleratorCount `json:"acceleratorCounts,omitempty" yaml:"acceleratorCounts,omitempty"`

	// List of accelerator manufacturer names.
	AcceleratorManufacturers []string `json:"acceleratorManufacturers,omitempty" yaml:"acceleratorManufacturers,omitempty"`

	// List of accelerator names.
	AcceleratorNames []string `json:"acceleratorNames,omitempty" yaml:"acceleratorNames,omitempty"`

	// Price protection threshold for Spot Instances.
	SpotMaxPricePercentageOverLowestPrice int `json:"spotMaxPricePercentageOverLowestPrice,omitempty" yaml:"spotMaxPricePercentageOverLowestPrice,omitempty"`

	// List of objects describing the minimum and maximum total storage (GB).
	TotalLocalStorageGbs []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementTotalLocalStorageGb `json:"totalLocalStorageGbs,omitempty" yaml:"totalLocalStorageGbs,omitempty"`

	// List of objects describing the minimum and maximum amount of network bandwidth (Gbps).
	NetworkBandwidthGbps []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementNetworkBandwidthGbp `json:"networkBandwidthGbps,omitempty" yaml:"networkBandwidthGbps,omitempty"`

	// List of objects describing the minimum and maximum number of vCPUs.
	VcpuCounts []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementVcpuCount `json:"vcpuCounts,omitempty" yaml:"vcpuCounts,omitempty"`

	// List of accelerator
	AcceleratorTypes []string `json:"acceleratorTypes,omitempty" yaml:"acceleratorTypes,omitempty"`

	// List of instance generation names.
	InstanceGenerations []string `json:"instanceGenerations,omitempty" yaml:"instanceGenerations,omitempty"`

	// List of objects describing the minimum and maximum amount of memory (GiB) per vCPU.
	MemoryGibPerVcpus []Autoscaling_getGroupMixedInstancesPolicyLaunchTemplateOverrideInstanceRequirementMemoryGibPerVcpus `json:"memoryGibPerVcpus,omitempty" yaml:"memoryGibPerVcpus,omitempty"`
}
