package types

type Evidently_ProjectDataDeliveryS3Destination struct {
	// The name of the bucket in which Evidently stores evaluation events.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// The bucket prefix in which Evidently stores evaluation events.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
