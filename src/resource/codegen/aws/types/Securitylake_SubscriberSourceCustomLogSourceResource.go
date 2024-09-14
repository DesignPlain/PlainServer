package types

type Securitylake_SubscriberSourceCustomLogSourceResource struct {
	// The details of the log provider for the third-party custom source. See `provider` Block below.
	Providers []Securitylake_SubscriberSourceCustomLogSourceResourceProvider `json:"providers,omitempty" yaml:"providers,omitempty"`

	// The name for a third-party custom source. This must be a Regionally unique value.
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`

	// The version for a third-party custom source. This must be a Regionally unique value.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`

	// The attributes of the third-party custom source. See `attributes` Block below.
	Attributes []Securitylake_SubscriberSourceCustomLogSourceResourceAttribute `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}
