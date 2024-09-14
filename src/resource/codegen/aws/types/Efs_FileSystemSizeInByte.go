package types

type Efs_FileSystemSizeInByte struct {
	// The latest known metered size (in bytes) of data stored in the file system.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`

	// The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.
	ValueInIa int `json:"valueInIa,omitempty" yaml:"valueInIa,omitempty"`

	// The latest known metered size (in bytes) of data stored in the Standard storage class.
	ValueInStandard int `json:"valueInStandard,omitempty" yaml:"valueInStandard,omitempty"`
}
