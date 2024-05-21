package eventarc

type GoogleChannelConfig struct {
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/-/locations/-/keyRings/-/cryptoKeys/-`.
	CryptoKeyName string `json:"cryptoKeyName,omitempty" yaml:"cryptoKeyName,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Required. The resource name of the config. Must be in the format of, `projects/{project}/locations/{location}/googleChannelConfig`.



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
