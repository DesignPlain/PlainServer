package types

type Osconfig_GuestPoliciesRecipeArtifactRemote struct {
	/*
	   Must be provided if allowInsecure is false. SHA256 checksum in hex format, to compare to the checksum of the artifact.
	   If the checksum is not empty and it doesn't match the artifact then the recipe installation fails before running any
	   of the steps.
	*/
	CheckSum string `json:"checkSum,omitempty" yaml:"checkSum,omitempty"`

	// URI from which to fetch the object. It should contain both the protocol and path following the format {protocol}://{location}.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
