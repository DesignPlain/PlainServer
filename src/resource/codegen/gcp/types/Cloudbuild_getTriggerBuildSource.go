package types

type Cloudbuild_getTriggerBuildSource struct {
	// Location of the source in a Google Cloud Source Repository.
	RepoSources []Cloudbuild_getTriggerBuildSourceRepoSource `json:"repoSources,omitempty" yaml:"repoSources,omitempty"`

	// Location of the source in an archive file in Google Cloud Storage.
	StorageSources []Cloudbuild_getTriggerBuildSourceStorageSource `json:"storageSources,omitempty" yaml:"storageSources,omitempty"`
}
