package types

type Securitylake_DataLakeConfigurationLifecycleConfiguration struct {
	// Provides data expiration details of Amazon Security Lake object.
	Expiration Securitylake_DataLakeConfigurationLifecycleConfigurationExpiration `json:"expiration,omitempty" yaml:"expiration,omitempty"`

	// Provides data storage transition details of Amazon Security Lake object.
	Transitions []Securitylake_DataLakeConfigurationLifecycleConfigurationTransition `json:"transitions,omitempty" yaml:"transitions,omitempty"`
}
