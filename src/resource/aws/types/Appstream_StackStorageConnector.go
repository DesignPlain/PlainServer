package types

type Appstream_StackStorageConnector struct {
	/*
	   Type of storage connector.
	   Valid values are `HOMEFOLDERS`, `GOOGLE_DRIVE`, or `ONE_DRIVE`.
	*/
	ConnectorType string `json:"connectorType,omitempty" yaml:"connectorType,omitempty"`

	// Names of the domains for the account.
	Domains []string `json:"domains,omitempty" yaml:"domains,omitempty"`

	// ARN of the storage connector.
	ResourceIdentifier string `json:"resourceIdentifier,omitempty" yaml:"resourceIdentifier,omitempty"`
}
