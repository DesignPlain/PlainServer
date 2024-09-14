package types

type Glue_CatalogDatabaseFederatedDatabase struct {
	// Unique identifier for the federated database.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	// Name of the connection to the external metastore.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`
}
