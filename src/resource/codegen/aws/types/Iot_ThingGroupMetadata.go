package types

type Iot_ThingGroupMetadata struct {
	//
	CreationDate string `json:"creationDate,omitempty" yaml:"creationDate,omitempty"`

	// The name of the parent Thing Group.
	ParentGroupName string `json:"parentGroupName,omitempty" yaml:"parentGroupName,omitempty"`

	//
	RootToParentGroups []Iot_ThingGroupMetadataRootToParentGroup `json:"rootToParentGroups,omitempty" yaml:"rootToParentGroups,omitempty"`
}
