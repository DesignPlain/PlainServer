package types

type Osconfig_GuestPoliciesRecipeArtifact struct {
	/*
	   Defaults to false. When false, recipes are subject to validations based on the artifact type:
	   Remote: A checksum must be specified, and only protocols with transport-layer security are permitted.
	   GCS: An object generation number must be specified.
	*/
	AllowInsecure bool `json:"allowInsecure,omitempty" yaml:"allowInsecure,omitempty"`

	/*
	   A Google Cloud Storage artifact.
	   Structure is documented below.
	*/
	Gcs Osconfig_GuestPoliciesRecipeArtifactGcs `json:"gcs,omitempty" yaml:"gcs,omitempty"`

	/*
	   Id of the artifact, which the installation and update steps of this recipe can reference.
	   Artifacts in a recipe cannot have the same id.
	*/
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	/*
	   A generic remote artifact.
	   Structure is documented below.
	*/
	Remote Osconfig_GuestPoliciesRecipeArtifactRemote `json:"remote,omitempty" yaml:"remote,omitempty"`
}
