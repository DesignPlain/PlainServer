package types

type Bigquery_getDatasetAccess struct {
	/*
	   A view from a different dataset to grant access to. Queries
	   executed against that view will have read access to tables in
	   this dataset. The role field is not required when this field is
	   set. If that view is updated by any user, access to the view
	   needs to be granted again via an update operation.
	*/
	Views []Bigquery_getDatasetAccessView `json:"views,omitempty" yaml:"views,omitempty"`

	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Datasets []Bigquery_getDatasetAccessDataset `json:"datasets,omitempty" yaml:"datasets,omitempty"`

	/*
	   A domain to grant access to. Any users signed in with the
	   domain specified will be granted the specified access
	*/
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`

	// An email address of a Google Group to grant access to.
	GroupByEmail string `json:"groupByEmail,omitempty" yaml:"groupByEmail,omitempty"`

	/*
	   Describes the rights granted to the user specified by the other
	   member of the access object. Basic, predefined, and custom roles
	   are supported. Predefined roles that have equivalent basic roles
	   are swapped by the API to their basic counterparts. See
	   [official docs](https://cloud.google.com/bigquery/docs/access-control).
	*/
	Role string `json:"role,omitempty" yaml:"role,omitempty"`

	/*
	   A routine from a different dataset to grant access to. Queries
	   executed against that routine will have read access to tables in
	   this dataset. The role field is not required when this field is
	   set. If that routine is updated by any user, access to the routine
	   needs to be granted again via an update operation.
	*/
	Routines []Bigquery_getDatasetAccessRoutine `json:"routines,omitempty" yaml:"routines,omitempty"`

	/*
	   Some other type of member that appears in the IAM Policy but isn't a user,
	   group, domain, or special group. For example: 'allUsers'
	*/
	IamMember string `json:"iamMember,omitempty" yaml:"iamMember,omitempty"`

	/*
	   A special group to grant access to. Possible values include:


	   - 'projectOwners': Owners of the enclosing project.


	   - 'projectReaders': Readers of the enclosing project.


	   - 'projectWriters': Writers of the enclosing project.


	   - 'allAuthenticatedUsers': All authenticated BigQuery users.
	*/
	SpecialGroup string `json:"specialGroup,omitempty" yaml:"specialGroup,omitempty"`

	/*
	   An email address of a user to grant access to. For example:
	   fred@example.com
	*/
	UserByEmail string `json:"userByEmail,omitempty" yaml:"userByEmail,omitempty"`
}
