package types

type Beyondcorp_getAppConnectorPrincipalInfo struct {
	// ServiceAccount represents a GCP service account.
	ServiceAccounts []Beyondcorp_getAppConnectorPrincipalInfoServiceAccount `json:"serviceAccounts,omitempty" yaml:"serviceAccounts,omitempty"`
}
