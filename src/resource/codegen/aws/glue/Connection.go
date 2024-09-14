package glue

import types "libds/aws/types"

type Connection struct {
	// ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	/*
	   Map of key-value pairs used as parameters for this connection. For more information, see the [AWS Documentation](https://docs.aws.amazon.com/glue/latest/dg/connection-properties.html).

	   --Note:-- Some connection types require the `SparkProperties` property with a JSON document that contains the actual connection properties. For specific examples, refer to Example Usage.
	*/
	ConnectionProperties map[string]string `json:"connectionProperties,omitempty" yaml:"connectionProperties,omitempty"`

	// Type of the connection. Valid values: `AZURECOSMOS`, `AZURESQL`, `BIGQUERY`, `CUSTOM`, `JDBC`, `KAFKA`, `MARKETPLACE`, `MONGODB`, `NETWORK`, `OPENSEARCH`, `SNOWFLAKE`. Defaults to `JDBC`.
	ConnectionType string `json:"connectionType,omitempty" yaml:"connectionType,omitempty"`

	// Description of the connection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// List of criteria that can be used in selecting this connection.
	MatchCriterias []string `json:"matchCriterias,omitempty" yaml:"matchCriterias,omitempty"`

	/*
	   Name of the connection.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Map of physical connection requirements, such as VPC and SecurityGroup. See `physical_connection_requirements` Block for details.
	PhysicalConnectionRequirements types.Glue_ConnectionPhysicalConnectionRequirements `json:"physicalConnectionRequirements,omitempty" yaml:"physicalConnectionRequirements,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
