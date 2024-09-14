package ivs

type PlaybackKeyPair struct {
	// Playback Key Pair name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Public portion of a customer-generated key pair. Must be an ECDSA public key in PEM format.

	   The following arguments are optional:
	*/
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
