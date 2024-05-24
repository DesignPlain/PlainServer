package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerParquetSerDe struct {
	// The Hadoop Distributed File System (HDFS) block size. This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Kinesis Data Firehose uses this value for padding calculations.
	BlockSizeBytes int `json:"blockSizeBytes,omitempty" yaml:"blockSizeBytes,omitempty"`

	// The compression code to use over data blocks. The possible values are `UNCOMPRESSED`, `SNAPPY`, and `GZIP`, with the default being `SNAPPY`. Use `SNAPPY` for higher decompression speed. Use `GZIP` if the compression ratio is more important than speed.
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// Indicates whether to enable dictionary compression.
	EnableDictionaryCompression bool `json:"enableDictionaryCompression,omitempty" yaml:"enableDictionaryCompression,omitempty"`

	// The maximum amount of padding to apply. This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is `0`.
	MaxPaddingBytes int `json:"maxPaddingBytes,omitempty" yaml:"maxPaddingBytes,omitempty"`

	// The Parquet page size. Column chunks are divided into pages. A page is conceptually an indivisible unit (in terms of compression and encoding). The minimum value is 64 KiB and the default is 1 MiB.
	PageSizeBytes int `json:"pageSizeBytes,omitempty" yaml:"pageSizeBytes,omitempty"`

	// Indicates the version of row format to output. The possible values are `V1` and `V2`. The default is `V1`.
	WriterVersion string `json:"writerVersion,omitempty" yaml:"writerVersion,omitempty"`
}
