package types

type Dataloss_PreventionJobTriggerInspectJobStorageConfigCloudStorageOptions struct {
	/*
	   Max number of bytes to scan from a file. If a scanned file's size is bigger than this value
	   then the rest of the bytes are omitted.
	*/
	BytesLimitPerFile int `json:"bytesLimitPerFile,omitempty" yaml:"bytesLimitPerFile,omitempty"`

	/*
	   Max percentage of bytes to scan from a file. The rest are omitted. The number of bytes scanned is rounded down.
	   Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	*/
	BytesLimitPerFilePercent int `json:"bytesLimitPerFilePercent,omitempty" yaml:"bytesLimitPerFilePercent,omitempty"`

	/*
	   Set of files to scan.
	   Structure is documented below.
	*/
	FileSet Dataloss_PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSet `json:"fileSet,omitempty" yaml:"fileSet,omitempty"`

	/*
	   List of file type groups to include in the scan. If empty, all files are scanned and available data
	   format processors are applied. In addition, the binary content of the selected files is always scanned as well.
	   Images are scanned only as binary if the specified region does not support image inspection and no fileTypes were specified.
	   Each value may be one of: `BINARY_FILE`, `TEXT_FILE`, `IMAGE`, `WORD`, `PDF`, `AVRO`, `CSV`, `TSV`, `POWERPOINT`, `EXCEL`.
	*/
	FileTypes []string `json:"fileTypes,omitempty" yaml:"fileTypes,omitempty"`

	/*
	   Limits the number of files to scan to this percentage of the input FileSet. Number of files scanned is rounded down.
	   Must be between 0 and 100, inclusively. Both 0 and 100 means no limit.
	*/
	FilesLimitPercent int `json:"filesLimitPercent,omitempty" yaml:"filesLimitPercent,omitempty"`

	/*
	   How to sample bytes if not all bytes are scanned. Meaningful only when used in conjunction with bytesLimitPerFile.
	   If not specified, scanning would start from the top.
	   Possible values are: `TOP`, `RANDOM_START`.
	*/
	SampleMethod string `json:"sampleMethod,omitempty" yaml:"sampleMethod,omitempty"`
}
