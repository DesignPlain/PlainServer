package monitoring

type Group struct {
	/*
	   The filter used to determine which monitored resources
	   belong to this group.


	   - - -
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	/*
	   If true, the members of this group are considered to be a
	   cluster. The system can perform additional analysis on
	   groups that are clusters.
	*/
	IsCluster bool `json:"isCluster,omitempty" yaml:"isCluster,omitempty"`

	/*
	   The name of the group's parent, if it has one. The format is
	   "projects/{project_id_or_number}/groups/{group_id}". For
	   groups with no parent, parentName is the empty string, "".
	*/
	ParentName string `json:"parentName,omitempty" yaml:"parentName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A user-assigned name for this group, used only for display
	   purposes.
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
