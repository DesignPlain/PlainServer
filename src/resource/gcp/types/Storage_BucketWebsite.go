package types

type Storage_BucketWebsite struct {
	/*
	   Behaves as the bucket's directory index where
	   missing objects are treated as potential directories.
	*/
	MainPageSuffix string `json:"mainPageSuffix,omitempty" yaml:"mainPageSuffix,omitempty"`

	/*
	   The custom object to return when a requested
	   resource is not found.
	*/
	NotFoundPage string `json:"notFoundPage,omitempty" yaml:"notFoundPage,omitempty"`
}
