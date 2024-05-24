package types

type Ebs_getSnapshotIdsFilter struct {
	//
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
