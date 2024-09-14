package types

type Compute_getMachineTypesMachineTypeBundledLocalSsd struct {
	// The number of partitions.
	PartitionCount int `json:"partitionCount,omitempty" yaml:"partitionCount,omitempty"`

	// The default disk interface if the interface is not specified.
	DefaultInterface string `json:"defaultInterface,omitempty" yaml:"defaultInterface,omitempty"`
}
