package types

type Datasync_LocationAzureBlobSasConfiguration struct {
	// A SAS token that provides permissions to access your Azure Blob Storage.
	Token string `json:"token,omitempty" yaml:"token,omitempty"`
}
