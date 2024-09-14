package types

type Connect_UserHierarchyGroupHierarchyPathLevelThree struct {
	// The identifier of the hierarchy group.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name of the user hierarchy group. Must not be more than 100 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Amazon Resource Name (ARN) of the hierarchy group.
	Arn string `json:"arn,omitempty" yaml:"arn,omitempty"`
}
