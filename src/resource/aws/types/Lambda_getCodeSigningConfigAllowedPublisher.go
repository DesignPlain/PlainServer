package types

type Lambda_getCodeSigningConfigAllowedPublisher struct {
	// The ARN for each of the signing profiles. A signing profile defines a trusted user who can sign a code package.
	SigningProfileVersionArns []string `json:"signingProfileVersionArns,omitempty" yaml:"signingProfileVersionArns,omitempty"`
}
