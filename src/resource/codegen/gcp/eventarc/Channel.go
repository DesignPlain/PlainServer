package eventarc

type Channel struct {
	// Optional. Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt their event data. It must match the pattern `projects/-/locations/-/keyRings/-/cryptoKeys/-`.
	CryptoKeyName string `json:"cryptoKeyName,omitempty" yaml:"cryptoKeyName,omitempty"`

	// The location for the resource
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Required. The resource name of the channel. Must be unique within the location on the project.



	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The project for the resource
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The name of the event provider (e.g. Eventarc SaaS partner) associated with the channel. This provider will be granted permissions to publish events to the channel. Format: `projects/{project}/locations/{location}/providers/{provider_id}`.
	ThirdPartyProvider string `json:"thirdPartyProvider,omitempty" yaml:"thirdPartyProvider,omitempty"`
}
