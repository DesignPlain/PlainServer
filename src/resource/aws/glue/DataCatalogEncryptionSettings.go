package glue

import types "DesignSphere_Server/src/resource/aws/types"

type DataCatalogEncryptionSettings struct {
	// The ID of the Data Catalog to set the security configuration for. If none is provided, the AWS account ID is used by default.
	CatalogId string `json:"catalogId,omitempty" yaml:"catalogId,omitempty"`

	// The security configuration to set. see Data Catalog Encryption Settings.
	DataCatalogEncryptionSettings types.Glue_DataCatalogEncryptionSettingsDataCatalogEncryptionSettings `json:"dataCatalogEncryptionSettings,omitempty" yaml:"dataCatalogEncryptionSettings,omitempty"`
}
