package tags

type TagValue struct {
	/*
	   Input only. User-assigned short name for TagValue. The short name should be unique for TagValues within the same parent TagKey.
	   The short name must be 63 characters or less, beginning and ending with an alphanumeric character ([a-z0-9A-Z]) with dashes (-), underscores (_), dots (.), and alphanumerics between.


	   - - -
	*/
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty"`

	// User-assigned description of the TagValue. Must not exceed 256 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Input only. The resource name of the new TagValue's parent. Must be of the form tagKeys/{tag_key_id}.
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
