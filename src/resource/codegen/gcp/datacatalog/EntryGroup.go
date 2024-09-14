package datacatalog

type EntryGroup struct {
	/*
	   The id of the entry group to create. The id must begin with a letter or underscore,
	   contain only English letters, numbers and underscores, and be at most 64 characters.


	   - - -
	*/
	EntryGroupId string `json:"entryGroupId,omitempty" yaml:"entryGroupId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// EntryGroup location region.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Entry group description, which can consist of several sentences or paragraphs that describe entry group contents.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A short name to identify the entry group, for example, "analytics data - jan 2011".
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
