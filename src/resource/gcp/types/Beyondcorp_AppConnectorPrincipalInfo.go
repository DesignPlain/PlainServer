package types

type Beyondcorp_AppConnectorPrincipalInfo struct {
	/*
	   ServiceAccount represents a GCP service account.
	   Structure is documented below.
	*/
	ServiceAccount Beyondcorp_AppConnectorPrincipalInfoServiceAccount `json:"serviceAccount,omitempty" yaml:"serviceAccount,omitempty"`
}
