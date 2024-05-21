package types

type Cloudbuild_TriggerBuildArtifactsMavenArtifact struct {
	// Maven artifactId value used when uploading the artifact to Artifact Registry.
	ArtifactId string `json:"artifactId,omitempty" yaml:"artifactId,omitempty"`

	// Maven groupId value used when uploading the artifact to Artifact Registry.
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`

	// Path to an artifact in the build's workspace to be uploaded to Artifact Registry. This can be either an absolute path, e.g. /workspace/my-app/target/my-app-1.0.SNAPSHOT.jar or a relative path from /workspace, e.g. my-app/target/my-app-1.0.SNAPSHOT.jar.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	/*
	   Artifact Registry repository, in the form "https://$REGION-maven.pkg.dev/$PROJECT/$REPOSITORY"
	   Artifact in the workspace specified by path will be uploaded to Artifact Registry with this location as a prefix.
	*/
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Maven version value used when uploading the artifact to Artifact Registry.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
