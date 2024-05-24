package types

type Kinesis_FirehoseDeliveryStreamExtendedS3ConfigurationDataFormatConversionConfigurationOutputFormatConfigurationSerializerOrcSerDe struct {
	// The Bloom filter false positive probability (FPP). The lower the FPP, the bigger the Bloom filter. The default value is `0.05`, the minimum is `0`, and the maximum is `1`.
	BloomFilterFalsePositiveProbability float64 `json:"bloomFilterFalsePositiveProbability,omitempty" yaml:"bloomFilterFalsePositiveProbability,omitempty"`

	// The compression code to use over data blocks. The default is `SNAPPY`.
	Compression string `json:"compression,omitempty" yaml:"compression,omitempty"`

	// Set this to `true` to indicate that you want stripes to be padded to the HDFS block boundaries. This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is `false`.
	EnablePadding bool `json:"enablePadding,omitempty" yaml:"enablePadding,omitempty"`

	// The version of the file to write. The possible values are `V0_11` and `V0_12`. The default is `V0_12`.
	FormatVersion string `json:"formatVersion,omitempty" yaml:"formatVersion,omitempty"`

	// The number of bytes in each stripe. The default is 64 MiB and the minimum is 8 MiB.
	StripeSizeBytes int `json:"stripeSizeBytes,omitempty" yaml:"stripeSizeBytes,omitempty"`

	// The Hadoop Distributed File System (HDFS) block size. This is useful if you intend to copy the data from Amazon S3 to HDFS before querying. The default is 256 MiB and the minimum is 64 MiB. Kinesis Data Firehose uses this value for padding calculations.
	BlockSizeBytes int `json:"blockSizeBytes,omitempty" yaml:"blockSizeBytes,omitempty"`

	// A float that represents the fraction of the total number of non-null rows. To turn off dictionary encoding, set this fraction to a number that is less than the number of distinct keys in a dictionary. To always use dictionary encoding, set this threshold to `1`.
	DictionaryKeyThreshold float64 `json:"dictionaryKeyThreshold,omitempty" yaml:"dictionaryKeyThreshold,omitempty"`

	// A float between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size. The default value is `0.05`, which means 5 percent of stripe size. For the default values of 64 MiB ORC stripes and 256 MiB HDFS blocks, the default block padding tolerance of 5 percent reserves a maximum of 3.2 MiB for padding within the 256 MiB block. In such a case, if the available size within the block is more than 3.2 MiB, a new, smaller stripe is inserted to fit within that space. This ensures that no stripe crosses block boundaries and causes remote reads within a node-local task. Kinesis Data Firehose ignores this parameter when `enable_padding` is `false`.
	PaddingTolerance float64 `json:"paddingTolerance,omitempty" yaml:"paddingTolerance,omitempty"`

	// The number of rows between index entries. The default is `10000` and the minimum is `1000`.
	RowIndexStride int `json:"rowIndexStride,omitempty" yaml:"rowIndexStride,omitempty"`

	// A list of column names for which you want Kinesis Data Firehose to create bloom filters.
	BloomFilterColumns []string `json:"bloomFilterColumns,omitempty" yaml:"bloomFilterColumns,omitempty"`
}
