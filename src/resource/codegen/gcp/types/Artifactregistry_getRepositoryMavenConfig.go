package types

type Artifactregistry_getRepositoryMavenConfig struct {
	/*
	   The repository with this flag will allow publishing the same
	   snapshot versions.
	*/
	AllowSnapshotOverwrites bool `json:"allowSnapshotOverwrites,omitempty" yaml:"allowSnapshotOverwrites,omitempty"`

	// Version policy defines the versions that the registry will accept. Default value: "VERSION_POLICY_UNSPECIFIED" Possible values: ["VERSION_POLICY_UNSPECIFIED", "RELEASE", "SNAPSHOT"]
	VersionPolicy string `json:"versionPolicy,omitempty" yaml:"versionPolicy,omitempty"`
}
