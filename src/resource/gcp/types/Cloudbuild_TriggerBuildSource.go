package types

type Cloudbuild_TriggerBuildSource struct {
	/*
	   Location of the source in a Google Cloud Source Repository.
	   Structure is documented below.
	*/
	RepoSource Cloudbuild_TriggerBuildSourceRepoSource `json:"repoSource,omitempty" yaml:"repoSource,omitempty"`

	/*
	   Location of the source in an archive file in Google Cloud Storage.
	   Structure is documented below.
	*/
	StorageSource Cloudbuild_TriggerBuildSourceStorageSource `json:"storageSource,omitempty" yaml:"storageSource,omitempty"`
}
