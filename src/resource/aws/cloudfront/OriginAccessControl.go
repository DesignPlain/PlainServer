package cloudfront

type OriginAccessControl struct {
	// A name that identifies the Origin Access Control.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of origin that this Origin Access Control is for. Valid values are `s3`, and `mediastore`.
	OriginAccessControlOriginType string `json:"originAccessControlOriginType,omitempty" yaml:"originAccessControlOriginType,omitempty"`

	// Specifies which requests CloudFront signs. Specify `always` for the most common use case. Allowed values: `always`, `never`, and `no-override`.
	SigningBehavior string `json:"signingBehavior,omitempty" yaml:"signingBehavior,omitempty"`

	// Determines how CloudFront signs (authenticates) requests. The only valid value is `sigv4`.
	SigningProtocol string `json:"signingProtocol,omitempty" yaml:"signingProtocol,omitempty"`

	// The description of the Origin Access Control. Defaults to "Managed by Pulumi" if omitted.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
