package types

type S3control_MultiRegionAccessPointDetailsPublicAccessBlock struct {
	//
	BlockPublicAcls bool `json:"blockPublicAcls,omitempty" yaml:"blockPublicAcls,omitempty"`

	//
	BlockPublicPolicy bool `json:"blockPublicPolicy,omitempty" yaml:"blockPublicPolicy,omitempty"`

	//
	IgnorePublicAcls bool `json:"ignorePublicAcls,omitempty" yaml:"ignorePublicAcls,omitempty"`

	//
	RestrictPublicBuckets bool `json:"restrictPublicBuckets,omitempty" yaml:"restrictPublicBuckets,omitempty"`
}
