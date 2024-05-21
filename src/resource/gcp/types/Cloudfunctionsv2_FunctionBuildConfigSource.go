package types

type Cloudfunctionsv2_FunctionBuildConfigSource struct {
	/*
	   If provided, get the source from this location in a Cloud Source Repository.
	   Structure is documented below.
	*/
	RepoSource Cloudfunctionsv2_FunctionBuildConfigSourceRepoSource `json:"repoSource,omitempty" yaml:"repoSource,omitempty"`

	/*
	   If provided, get the source from this location in Google Cloud Storage.
	   Structure is documented below.
	*/
	StorageSource Cloudfunctionsv2_FunctionBuildConfigSourceStorageSource `json:"storageSource,omitempty" yaml:"storageSource,omitempty"`
}
