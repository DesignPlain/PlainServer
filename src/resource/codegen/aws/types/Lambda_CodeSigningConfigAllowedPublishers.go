package types

type Lambda_CodeSigningConfigAllowedPublishers struct {
	// The Amazon Resource Name (ARN) for each of the signing profiles. A signing profile defines a trusted user who can sign a code package.
	SigningProfileVersionArns []string `json:"signingProfileVersionArns,omitempty" yaml:"signingProfileVersionArns,omitempty"`
}
