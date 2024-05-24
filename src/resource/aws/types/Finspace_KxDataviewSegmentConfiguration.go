package types

type Finspace_KxDataviewSegmentConfiguration struct {
	// The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
	DbPaths []string `json:"dbPaths,omitempty" yaml:"dbPaths,omitempty"`

	// The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
	VolumeName string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
