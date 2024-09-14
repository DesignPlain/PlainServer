package types

type Compute_InstanceServiceAccount struct {
	/*
	   The service account e-mail address.
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	Email string `json:"email,omitempty" yaml:"email,omitempty"`

	/*
	   A list of service scopes. Both OAuth2 URLs and gcloud
	   short names are supported. To allow full access to all Cloud APIs, use the
	   `cloud-platform` scope. See a complete list of scopes [here](https://cloud.google.com/sdk/gcloud/reference/alpha/compute/instances/set-scopes#--scopes).
	   --Note--: `allow_stopping_for_update` must be set to true or your instance must have a `desired_status` of `TERMINATED` in order to update this field.
	*/
	Scopes []string `json:"scopes,omitempty" yaml:"scopes,omitempty"`
}
