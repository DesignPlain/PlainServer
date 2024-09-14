package types

type Timestreamwrite_getTableMagneticStoreWriteProperty struct {
	// Flag that is set based on if magnetic store writes are enabled.
	EnableMagneticStoreWrites bool `json:"enableMagneticStoreWrites,omitempty" yaml:"enableMagneticStoreWrites,omitempty"`

	// Object containing the following attributes to describe error reports for records rejected during magnetic store writes.
	MagneticStoreRejectedDataLocations []Timestreamwrite_getTableMagneticStoreWritePropertyMagneticStoreRejectedDataLocation `json:"magneticStoreRejectedDataLocations,omitempty" yaml:"magneticStoreRejectedDataLocations,omitempty"`
}
