package serviceaccount

type Key struct {
	/*
	   The algorithm used to generate the key. KEY_ALG_RSA_2048 is the default algorithm.
	   Valid values are listed at
	   [ServiceAccountPrivateKeyType](https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKeyAlgorithm)
	   (only used on create)
	*/
	KeyAlgorithm string `json:"keyAlgorithm,omitempty" yaml:"keyAlgorithm,omitempty"`

	// The output format of the private key. TYPE_GOOGLE_CREDENTIALS_FILE is the default output format.
	PrivateKeyType string `json:"privateKeyType,omitempty" yaml:"privateKeyType,omitempty"`

	// Public key data to create a service account key for given service account. The expected format for this field is a base64 encoded X509_PEM and it conflicts with `public_key_type` and `private_key_type`.
	PublicKeyData string `json:"publicKeyData,omitempty" yaml:"publicKeyData,omitempty"`

	// The output format of the public key requested. TYPE_X509_PEM_FILE is the default output format.
	PublicKeyType string `json:"publicKeyType,omitempty" yaml:"publicKeyType,omitempty"`

	/*
	   The Service account id of the Key. This can be a string in the format
	   `{ACCOUNT}` or `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`. If the `{ACCOUNT}`-only syntax is used, either
	   the --full-- email address of the service account or its name can be specified as a value, in which case the project will
	   automatically be inferred from the account. Otherwise, if the `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}`
	   syntax is used, the `{ACCOUNT}` specified can be the full email address of the service account or the service account's
	   unique id. Substituting `-` as a wildcard for the `{PROJECT_ID}` will infer the project from the account.
	*/
	ServiceAccountId string `json:"serviceAccountId,omitempty" yaml:"serviceAccountId,omitempty"`

	// Arbitrary map of values that, when changed, will trigger a new key to be generated.
	Keepers map[string]string `json:"keepers,omitempty" yaml:"keepers,omitempty"`
}
