package glue

import types "DesignSphere_Server/src/resource/aws/types"

type Connection struct {
	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The ID of the Data Catalog in which to create the connection. If none is supplied, the AWS account ID is used by default.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// A map of key-value pairs used as parameters for this connection.
	ConnectionProperties map[string]string `json:"connectionProperties,omitempty" yaml:"connectionProperties,omitempty"`

	// The type of the connection. Supported are: `CUSTOM`, `JDBC`, `KAFKA`, `MARKETPLACE`, `MONGODB`, and `NETWORK`. Defaults to `JDBC`.
	ConnectionType string `json:"connectionType,omitempty" yaml:"connectionType,omitempty"`

	// Description of the connection.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A list of criteria that can be used in selecting this connection.
	MatchCriterias []string `json:"matchCriterias,omitempty" yaml:"matchCriterias,omitempty"`

	// The name of the connection.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of physical connection requirements, such as VPC and SecurityGroup. Defined below.
	PhysicalConnectionRequirements types.Glue_ConnectionPhysicalConnectionRequirements `json:"physicalConnectionRequirements,omitempty" yaml:"physicalConnectionRequirements,omitempty"`
}
