package types

type S3_BucketOwnershipControlsRule struct {
	// Object ownership. Valid values: `BucketOwnerPreferred`, `ObjectWriter` or `BucketOwnerEnforced`
	ObjectOwnership string `json:"objectOwnership,omitempty" yaml:"objectOwnership,omitempty"`
}
