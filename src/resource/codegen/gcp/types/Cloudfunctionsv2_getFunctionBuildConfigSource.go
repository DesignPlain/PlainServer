package types

type Cloudfunctionsv2_getFunctionBuildConfigSource struct {
	// If provided, get the source from this location in Google Cloud Storage.
	StorageSources []Cloudfunctionsv2_getFunctionBuildConfigSourceStorageSource `json:"storageSources,omitempty" yaml:"storageSources,omitempty"`

	// If provided, get the source from this location in a Cloud Source Repository.
	RepoSources []Cloudfunctionsv2_getFunctionBuildConfigSourceRepoSource `json:"repoSources,omitempty" yaml:"repoSources,omitempty"`
}
