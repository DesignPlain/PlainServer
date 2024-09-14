package types

type Glacier_VaultNotification struct {
	// You can configure a vault to publish a notification for `ArchiveRetrievalCompleted` and `InventoryRetrievalCompleted` events.
	Events []string `json:"events,omitempty" yaml:"events,omitempty"`

	// The SNS Topic ARN.
	SnsTopic string `json:"snsTopic,omitempty" yaml:"snsTopic,omitempty"`
}
