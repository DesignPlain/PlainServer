package organizations

type Folder struct {
	/*
	   The folder’s display name.
	   A folder’s display name must be unique amongst its siblings, e.g. no two folders with the same parent can share the same display name. The display name must start and end with a letter or digit, may contain letters, digits, spaces, hyphens and underscores and can be no longer than 30 characters.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The resource name of the parent Folder or Organization.
	   Must be of the form `folders/{folder_id}` or `organizations/{org_id}`.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
