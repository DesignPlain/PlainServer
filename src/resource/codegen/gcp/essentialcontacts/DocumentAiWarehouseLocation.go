package essentialcontacts

type DocumentAiWarehouseLocation struct {
	/*
	   The access control mode for accessing the customer data.
	   Possible values are: `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_GCI`, `ACL_MODE_DOCUMENT_LEVEL_ACCESS_CONTROL_BYOID`, `ACL_MODE_UNIVERSAL_ACCESS`.
	*/
	AccessControlMode string `json:"accessControlMode,omitempty" yaml:"accessControlMode,omitempty"`

	/*
	   The type of database used to store customer data.
	   Possible values are: `DB_INFRA_SPANNER`, `DB_CLOUD_SQL_POSTGRES`.
	*/
	DatabaseType string `json:"databaseType,omitempty" yaml:"databaseType,omitempty"`

	/*
	   The default role for the person who create a document.
	   Possible values are: `DOCUMENT_ADMIN`, `DOCUMENT_EDITOR`, `DOCUMENT_VIEWER`.
	*/
	DocumentCreatorDefaultRole string `json:"documentCreatorDefaultRole,omitempty" yaml:"documentCreatorDefaultRole,omitempty"`

	/*
	   The KMS key used for CMEK encryption. It is required that
	   the kms key is in the same region as the endpoint. The
	   same key will be used for all provisioned resources, if
	   encryption is available. If the kmsKey is left empty, no
	   encryption will be enforced.
	*/
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	/*
	   The location in which the instance is to be provisioned. It takes the form projects/{projectNumber}/locations/{location}.


	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The unique identifier of the project.
	ProjectNumber string `json:"projectNumber,omitempty" yaml:"projectNumber,omitempty"`
}
