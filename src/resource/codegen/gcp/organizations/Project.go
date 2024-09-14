package organizations

type Project struct {
	/*
	   The numeric ID of the organization this project belongs to.
	   Changing this forces a new project to be created.  Only one of
	   `org_id` or `folder_id` may be specified. If the `org_id` is
	   specified then the project is created at the top level. Changing
	   this forces the project to be migrated to the newly specified
	   organization.
	*/
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`

	// The project ID. Changing this forces a new project to be created.
	ProjectId string `json:"projectId,omitempty" yaml:"projectId,omitempty"`

	/*
	   If true, the resource can be deleted
	   without deleting the Project via the Google API.
	*/
	SkipDelete bool `json:"skipDelete,omitempty" yaml:"skipDelete,omitempty"`

	/*
	   Create the 'default' network automatically. Default true. If set to false, the default network will be deleted. Note
	   that, for quota purposes, you will still need to have 1 network slot available to create the project successfully, even
	   if you set auto_create_network to false, since the network will exist momentarily.
	*/
	AutoCreateNetwork bool `json:"autoCreateNetwork,omitempty" yaml:"autoCreateNetwork,omitempty"`

	/*
	   The alphanumeric ID of the billing account this project
	   belongs to. The user or service account performing this operation with the provider
	   must have at mininum Billing Account User privileges (`roles/billing.user`) on the billing account.
	   See [Google Cloud Billing API Access Control](https://cloud.google.com/billing/docs/how-to/billing-access)
	   for more details.
	*/
	BillingAccount string `json:"billingAccount,omitempty" yaml:"billingAccount,omitempty"`

	/*
	   The numeric ID of the folder this project should be
	   created under. Only one of `org_id` or `folder_id` may be
	   specified. If the `folder_id` is specified, then the project is
	   created under the specified folder. Changing this forces the
	   project to be migrated to the newly specified folder.
	*/
	FolderId string `json:"folderId,omitempty" yaml:"folderId,omitempty"`

	/*
	   A set of key/value label pairs to assign to the project.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// The display name of the project.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
