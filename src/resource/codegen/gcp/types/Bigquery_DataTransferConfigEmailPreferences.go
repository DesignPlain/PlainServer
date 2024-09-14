package types

type Bigquery_DataTransferConfigEmailPreferences struct {
	// If true, email notifications will be sent on transfer run failures.
	EnableFailureEmail bool `json:"enableFailureEmail,omitempty" yaml:"enableFailureEmail,omitempty"`
}
