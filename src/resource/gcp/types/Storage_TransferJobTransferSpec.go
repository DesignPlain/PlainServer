package types

type Storage_TransferJobTransferSpec struct {
	// A Google Cloud Storage data source. Structure documented below.
	GcsDataSource Storage_TransferJobTransferSpecGcsDataSource `json:"gcsDataSource,omitempty" yaml:"gcsDataSource,omitempty"`

	// A HTTP URL data source. Structure documented below.
	HttpDataSource Storage_TransferJobTransferSpecHttpDataSource `json:"httpDataSource,omitempty" yaml:"httpDataSource,omitempty"`

	// Only objects that satisfy these object conditions are included in the set of data source and data sink objects. Object conditions based on objects' `last_modification_time` do not exclude objects in a data sink. Structure documented below.
	ObjectConditions Storage_TransferJobTransferSpecObjectConditions `json:"objectConditions,omitempty" yaml:"objectConditions,omitempty"`

	// A POSIX filesystem data source. Structure documented below.
	PosixDataSource Storage_TransferJobTransferSpecPosixDataSource `json:"posixDataSource,omitempty" yaml:"posixDataSource,omitempty"`

	// Specifies the agent pool name associated with the posix data source. When unspecified, the default name is used.
	SourceAgentPoolName string `json:"sourceAgentPoolName,omitempty" yaml:"sourceAgentPoolName,omitempty"`

	// An AWS S3 data source. Structure documented below.
	AwsS3DataSource Storage_TransferJobTransferSpecAwsS3DataSource `json:"awsS3DataSource,omitempty" yaml:"awsS3DataSource,omitempty"`

	// A Google Cloud Storage data sink. Structure documented below.
	GcsDataSink Storage_TransferJobTransferSpecGcsDataSink `json:"gcsDataSink,omitempty" yaml:"gcsDataSink,omitempty"`

	// A POSIX data sink. Structure documented below.
	PosixDataSink Storage_TransferJobTransferSpecPosixDataSink `json:"posixDataSink,omitempty" yaml:"posixDataSink,omitempty"`

	// Specifies the agent pool name associated with the posix data sink. When unspecified, the default name is used.
	SinkAgentPoolName string `json:"sinkAgentPoolName,omitempty" yaml:"sinkAgentPoolName,omitempty"`

	// Characteristics of how to treat files from datasource and sink during job. If the option `delete_objects_unique_in_sink` is true, object conditions based on objects' `last_modification_time` are ignored and do not exclude objects in a data source or a data sink. Structure documented below.
	TransferOptions Storage_TransferJobTransferSpecTransferOptions `json:"transferOptions,omitempty" yaml:"transferOptions,omitempty"`

	// An Azure Blob Storage data source. Structure documented below.
	AzureBlobStorageDataSource Storage_TransferJobTransferSpecAzureBlobStorageDataSource `json:"azureBlobStorageDataSource,omitempty" yaml:"azureBlobStorageDataSource,omitempty"`
}
