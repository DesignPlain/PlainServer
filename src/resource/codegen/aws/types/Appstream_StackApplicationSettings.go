package types

type Appstream_StackApplicationSettings struct {
	// Whether application settings should be persisted.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	/*
	   Name of the settings group.
	   Required when `enabled` is `true`.
	   Can be up to 100 characters.
	*/
	SettingsGroup string `json:"settingsGroup,omitempty" yaml:"settingsGroup,omitempty"`
}
