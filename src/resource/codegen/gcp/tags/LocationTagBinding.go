package tags

type LocationTagBinding struct {
	/*
	   Location of the target resource.

	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The full resource name of the resource the TagValue is bound to. E.g. //cloudresourcemanager.googleapis.com/projects/123
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	// The TagValue of the TagBinding. Must be of the form tagValues/456.
	TagValue string `json:"tagValue,omitempty" yaml:"tagValue,omitempty"`
}
