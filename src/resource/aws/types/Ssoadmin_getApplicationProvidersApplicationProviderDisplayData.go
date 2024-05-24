package types

type Ssoadmin_getApplicationProvidersApplicationProviderDisplayData struct {
	// URL that points to an icon that represents the application provider.
	IconUrl string `json:"iconUrl,omitempty" yaml:"iconUrl,omitempty"`

	// Description of the application provider.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Name of the application provider.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
