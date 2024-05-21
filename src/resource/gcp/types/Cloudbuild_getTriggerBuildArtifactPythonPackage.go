package types

type Cloudbuild_getTriggerBuildArtifactPythonPackage struct {
	// Path globs used to match files in the build's workspace. For Python/ Twine, this is usually dist/-, and sometimes additionally an .asc file.
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`

	/*
	   Artifact Registry repository, in the form "https://$REGION-python.pkg.dev/$PROJECT/$REPOSITORY"

	   Files in the workspace matching any path pattern will be uploaded to Artifact Registry with this location as a prefix.
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
