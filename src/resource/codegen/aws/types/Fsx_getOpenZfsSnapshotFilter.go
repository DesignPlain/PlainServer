package types

type Fsx_getOpenZfsSnapshotFilter struct {
	// Name of the snapshot.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
