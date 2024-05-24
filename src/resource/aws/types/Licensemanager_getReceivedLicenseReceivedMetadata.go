package types

type Licensemanager_getReceivedLicenseReceivedMetadata struct {
	// A list of allowed operations.
	AllowedOperations []string `json:"allowedOperations,omitempty" yaml:"allowedOperations,omitempty"`

	// Received status.
	ReceivedStatus string `json:"receivedStatus,omitempty" yaml:"receivedStatus,omitempty"`

	// Received status reason.
	ReceivedStatusReason string `json:"receivedStatusReason,omitempty" yaml:"receivedStatusReason,omitempty"`
}
