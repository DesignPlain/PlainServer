package types

type Organizations_PolicyRestorePolicy struct {
	// May only be set to true. If set, then the default Policy is restored.
	Default bool `json:"default,omitempty" yaml:"default,omitempty"`
}
