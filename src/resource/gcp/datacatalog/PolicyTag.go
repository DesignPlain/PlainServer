package datacatalog

type PolicyTag struct {
	/*
	   Taxonomy the policy tag is associated with


	   - - -
	*/
	Taxonomy string `json:"taxonomy,omitempty" yaml:"taxonomy,omitempty"`

	/*
	   Description of this policy tag. It must: contain only unicode characters, tabs,
	   newlines, carriage returns and page breaks; and be at most 2000 bytes long when
	   encoded in UTF-8. If not set, defaults to an empty description.
	   If not set, defaults to an empty description.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   User defined name of this policy tag. It must: be unique within the parent
	   taxonomy; contain only unicode letters, numbers, underscores, dashes and spaces;
	   not start or end with spaces; and be at most 200 bytes long when encoded in UTF-8.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Resource name of this policy tag's parent policy tag.
	   If empty, it means this policy tag is a top level policy tag.
	   If not set, defaults to an empty string.
	*/
	ParentPolicyTag string `json:"parentPolicyTag,omitempty" yaml:"parentPolicyTag,omitempty"`
}
