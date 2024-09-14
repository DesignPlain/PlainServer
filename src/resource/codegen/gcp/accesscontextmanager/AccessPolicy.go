package accesscontextmanager

type AccessPolicy struct {
	/*
	   The parent of this AccessPolicy in the Cloud Resource Hierarchy.
	   Format: organizations/{organization_id}
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`

	/*
	   Folder or project on which this policy is applicable.
	   Format: folders/{{folder_id}} or projects/{{project_id}}
	*/
	Scopes string `json:"scopes,omitempty" yaml:"scopes,omitempty"`

	/*
	   Human readable title. Does not affect behavior.


	   - - -
	*/
	Title string `json:"title,omitempty" yaml:"title,omitempty"`
}
