package types

type Rolesanywhere_TrustAnchorSource struct {
	// The data denoting the source of trust, documented below
	SourceData Rolesanywhere_TrustAnchorSourceSourceData `json:"sourceData,omitempty" yaml:"sourceData,omitempty"`

	// The type of the source of trust. Must be either `AWS_ACM_PCA` or `CERTIFICATE_BUNDLE`.
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`
}
