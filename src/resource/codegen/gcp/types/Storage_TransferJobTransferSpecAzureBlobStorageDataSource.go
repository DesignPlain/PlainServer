package types

type Storage_TransferJobTransferSpecAzureBlobStorageDataSource struct {
	// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The name of the Azure Storage account.
	StorageAccount string `json:"storageAccount,omitempty" yaml:"storageAccount,omitempty"`

	// Credentials used to authenticate API requests to Azure block.
	AzureCredentials Storage_TransferJobTransferSpecAzureBlobStorageDataSourceAzureCredentials `json:"azureCredentials,omitempty" yaml:"azureCredentials,omitempty"`

	// The container to transfer from the Azure Storage account.`
	Container string `json:"container,omitempty" yaml:"container,omitempty"`
}
