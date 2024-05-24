package backup

type GlobalSettings struct {
	// A list of resources along with the opt-in preferences for the account.
	GlobalSettings map[string]string `json:"globalSettings,omitempty" yaml:"globalSettings,omitempty"`
}
