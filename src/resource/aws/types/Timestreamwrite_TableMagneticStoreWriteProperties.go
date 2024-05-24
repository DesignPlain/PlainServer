package types

type Timestreamwrite_TableMagneticStoreWriteProperties struct {
	// A flag to enable magnetic store writes.
	EnableMagneticStoreWrites bool `json:"enableMagneticStoreWrites,omitempty" yaml:"enableMagneticStoreWrites,omitempty"`

	// The location to write error reports for records rejected asynchronously during magnetic store writes. See Magnetic Store Rejected Data Location below for more details.
	MagneticStoreRejectedDataLocation Timestreamwrite_TableMagneticStoreWritePropertiesMagneticStoreRejectedDataLocation `json:"magneticStoreRejectedDataLocation,omitempty" yaml:"magneticStoreRejectedDataLocation,omitempty"`
}
