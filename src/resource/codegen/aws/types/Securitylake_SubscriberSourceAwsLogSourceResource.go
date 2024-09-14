package types

type Securitylake_SubscriberSourceAwsLogSourceResource struct {
	// Provides data expiration details of Amazon Security Lake object.
	SourceName string `json:"sourceName,omitempty" yaml:"sourceName,omitempty"`

	// Provides data storage transition details of Amazon Security Lake object.
	SourceVersion string `json:"sourceVersion,omitempty" yaml:"sourceVersion,omitempty"`
}
