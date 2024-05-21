package types

type Storage_BucketLogging struct {
	// The bucket that will receive log objects.
	LogBucket string `json:"logBucket,omitempty" yaml:"logBucket,omitempty"`

	/*
	   The object prefix for log objects. If it's not provided,
	   by default GCS sets this to this bucket's name.
	*/
	LogObjectPrefix string `json:"logObjectPrefix,omitempty" yaml:"logObjectPrefix,omitempty"`
}
