package types

type Finspace_KxDataviewSegmentConfiguration struct {
	// The database path of the data that you want to place on each selected volume. Each segment must have a unique database path for each volume.
	DbPaths []string `json:"dbPaths,omitempty" yaml:"dbPaths,omitempty"`

	// Enables on-demand caching on the selected database path when a particular file or a column of the database is accessed. When on demand caching is --True--, dataviews perform minimal loading of files on the filesystem as needed. When it is set to --False--, everything is cached. The default value is --False--.
	OnDemand bool `json:"onDemand,omitempty" yaml:"onDemand,omitempty"`

	// The name of the volume that you want to attach to a dataview. This volume must be in the same availability zone as the dataview that you are attaching to.
	VolumeName string `json:"volumeName,omitempty" yaml:"volumeName,omitempty"`
}
