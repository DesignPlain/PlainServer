package oslogin

type SshPublicKey struct {
	/*
	   The user email.


	   - - -
	*/
	User string `json:"user,omitempty" yaml:"user,omitempty"`

	// An expiration time in microseconds since epoch.
	ExpirationTimeUsec string `json:"expirationTimeUsec,omitempty" yaml:"expirationTimeUsec,omitempty"`

	// Public key text in SSH format, defined by RFC4253 section 6.6.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The project ID of the Google Cloud Platform project.
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
