package types

type Cloudbuild_getTriggerBuildArtifactNpmPackage struct {
	// Path to the package.json. e.g. workspace/path/to/package
	PackagePath string `json:"packagePath,omitempty" yaml:"packagePath,omitempty"`

	/*
	   Artifact Registry repository, in the form "https://$REGION-npm.pkg.dev/$PROJECT/$REPOSITORY"

	   Npm package in the workspace specified by path will be zipped and uploaded to Artifact Registry with this location as a prefix.
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`
}
