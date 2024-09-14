package tags

type TagKey struct {
	/*
	   Optional. A purpose cannot be changed once set.
	   A purpose denotes that this Tag is intended for use in policies of a specific policy engine, and will involve that policy engine in management operations involving this Tag.
	   Possible values are: `GCE_FIREWALL`.
	*/
	Purpose string `json:"purpose,omitempty" yaml:"purpose,omitempty"`

	/*
	   Optional. Purpose data cannot be changed once set.
	   Purpose data corresponds to the policy system that the tag is intended for. For example, the GCE_FIREWALL purpose expects data in the following format: `network = "<project-name>/<vpc-name>"`.
	*/
	PurposeData map[string]string `json:"purposeData,omitempty" yaml:"purposeData,omitempty"`

	/*
	   Input only. The user friendly name for a TagKey. The short name should be unique for TagKeys within the same tag namespace.
	   The short name must be 1-63 characters, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.


	   - - -
	*/
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`

	// User-assigned description of the TagKey. Must not exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Input only. The resource name of the new TagKey's parent. Must be of the form organizations/{org_id} or projects/{project_id_or_number}.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
