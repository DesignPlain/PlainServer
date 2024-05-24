package glue

import types "DesignSphere_Server/src/resource/aws/types"

type UserDefinedFunction struct {
	// ID of the Glue Catalog to create the function in. If omitted, this defaults to the AWS Account ID.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// The Java class that contains the function code.
	ClassName string `json:"className,omitempty" yaml:"className,omitempty"`

	// The name of the Database to create the Function.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// The name of the function.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The owner of the function.
	OwnerName string `json:"ownerName,omitempty" yaml:"ownerName,omitempty"`

	// The owner type. can be one of `USER`, `ROLE`, and `GROUP`.
	OwnerType string `json:"ownerType,omitempty" yaml:"ownerType,omitempty"`

	// The configuration block for Resource URIs. See resource uris below for more details.
	ResourceUris []types.Glue_UserDefinedFunctionResourceUri `json:"resourceUris,omitempty" yaml:"resourceUris,omitempty"`
}
