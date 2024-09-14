package iam

type AccessKey struct {
	// IAM user to associate with this access key.
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// Either a base-64 encoded PGP public key, or a keybase username in the form `keybase:some_person_that_exists`, for use in the `encrypted_secret` output attribute. If providing a base-64 encoded PGP public key, make sure to provide the "raw" version and not the "armored" one (e.g. avoid passing the `-a` option to `gpg --export`).
	PgpKey string `json:"pgpKey,omitempty" yaml:"pgpKey,omitempty"`

	// Access key status to apply. Defaults to `Active`. Valid values are `Active` and `Inactive`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
