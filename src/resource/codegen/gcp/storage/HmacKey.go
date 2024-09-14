package storage

type HmacKey struct {
	/*
	   The email address of the key's associated service account.


	   - - -
	*/
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   The state of the key. Can be set to one of ACTIVE, INACTIVE.
	   Default value is `ACTIVE`.
	   Possible values are: `ACTIVE`, `INACTIVE`.
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
