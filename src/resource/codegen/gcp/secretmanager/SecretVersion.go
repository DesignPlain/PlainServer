package secretmanager

type SecretVersion struct {
	/*
	   The secret data. Must be no larger than 64KiB.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	SecretData string `json:"secretData,omitempty" yaml:"secretData,omitempty"`

	/*
	   The deletion policy for the secret version. Setting `ABANDON` allows the resource
	   to be abandoned rather than deleted. Setting `DISABLE` allows the resource to be
	   disabled rather than deleted. Default is `DELETE`. Possible values are:
	   - DELETE
	   - DISABLE
	   - ABANDON
	*/
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// The current state of the SecretVersion.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// If set to 'true', the secret data is expected to be base64-encoded string and would be sent as is.
	IsSecretDataBase64 bool `json:"isSecretDataBase64,omitempty" yaml:"isSecretDataBase64,omitempty"`

	/*
	   Secret Manager secret resource


	   - - -
	*/
	Secret string `json:"secret,omitempty" yaml:"secret,omitempty"`
}
