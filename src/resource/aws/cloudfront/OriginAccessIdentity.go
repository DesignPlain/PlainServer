package cloudfront

type OriginAccessIdentity struct {
	// An optional comment for the origin access identity.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`
}
