package types

type Connect_getUserHierarchyStructureHierarchyStructureLevelOne struct {
	// ARN of the hierarchy level.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`

	// The identifier of the hierarchy level.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Name of the user hierarchy level. Must not be more than 50 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
