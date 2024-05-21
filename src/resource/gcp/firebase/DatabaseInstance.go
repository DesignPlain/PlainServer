package firebase

type DatabaseInstance struct {
	/*
	   The globally unique identifier of the Firebase Realtime Database instance.
	   Instance IDs cannot be reused after deletion.


	   - - -
	*/
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   A reference to the region where the Firebase Realtime database resides.
	   Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The database type.
	   Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	   Creating user Databases is only available for projects on the Blaze plan.
	   Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo.
	   Default value is `USER_DATABASE`.
	   Possible values are: `DEFAULT_DATABASE`, `USER_DATABASE`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The intended database state.
	DesiredState string `json:"desiredState,omitempty" yaml:"desiredState,omitempty"`
}
