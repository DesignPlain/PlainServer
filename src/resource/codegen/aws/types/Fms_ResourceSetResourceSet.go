package types

type Fms_ResourceSetResourceSet struct {
	// Description of the resource set.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique identifier for the resource set. It's returned in the responses to create and list commands. You provide it to operations like update and delete.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Last time that the reosurce set was changed.
	LastUpdateTime string `json:"lastUpdateTime,omitempty" yaml:"lastUpdateTime,omitempty"`

	// Descriptive name of the resource set. You can't change the name of a resource set after you create it.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Indicates whether the resource set is in or out of the admin's Region scope. Valid values are `ACTIVE` (Admin can manage and delete the resource set) or `OUT_OF_ADMIN_SCOPE` (Admin can view the resource set, but theyy can't edit or delete the resource set.)
	ResourceSetStatus string `json:"resourceSetStatus,omitempty" yaml:"resourceSetStatus,omitempty"`

	// Determines the resources that can be associated to the resource set. Depending on your setting for max results and the number of resource sets, a single call might not return the full list.
	ResourceTypeLists []string `json:"resourceTypeLists,omitempty" yaml:"resourceTypeLists,omitempty"`

	//
	UpdateToken string `json:"updateToken,omitempty" yaml:"updateToken,omitempty"`
}
