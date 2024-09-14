package ec2

type SpotDatafeedSubscription struct {
	// The Amazon S3 bucket in which to store the Spot instance data feed.
	Bucket string `json:"bucket,omitempty" yaml:"bucket,omitempty"`

	// Path of folder inside bucket to place spot pricing data.
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`
}
