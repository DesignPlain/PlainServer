package types

type Fsx_getOntapStorageVirtualMachinesFilter struct {
	// Name of the field to filter by, as defined by [the underlying AWS API](https://docs.aws.amazon.com/fsx/latest/APIReference/API_StorageVirtualMachineFilter.html).
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Set of values that are accepted for the given field. An SVM will be selected if any one of the given values matches.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
