package glue

import types "DesignSphere_Server/src/resource/aws/types"

type CatalogDatabase struct {
	// ID of the Glue Catalog to create the database in. If omitted, this defaults to the AWS Account ID.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// Creates a set of default permissions on the table for principals. See `create_table_default_permission` below.
	CreateTableDefaultPermissions []types.Glue_CatalogDatabaseCreateTableDefaultPermission `json:"createTableDefaultPermissions,omitempty" yaml:"createTableDefaultPermissions,omitempty"`

	// Location of the database (for example, an HDFS path).
	LocationUri string `json:"locationUri,omitempty" yaml:"locationUri,omitempty"`

	// List of key-value pairs that define parameters and properties of the database.
	Parameters map[string]string `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Configuration block for a target database for resource linking. See `target_database` below.
	TargetDatabase types.Glue_CatalogDatabaseTargetDatabase `json:"targetDatabase,omitempty" yaml:"targetDatabase,omitempty"`

	// Description of the database.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block that references an entity outside the AWS Glue Data Catalog. See `federated_database` below.
	FederatedDatabase types.Glue_CatalogDatabaseFederatedDatabase `json:"federatedDatabase,omitempty" yaml:"federatedDatabase,omitempty"`

	// Name of the database. The acceptable characters are lowercase letters, numbers, and the underscore character.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
