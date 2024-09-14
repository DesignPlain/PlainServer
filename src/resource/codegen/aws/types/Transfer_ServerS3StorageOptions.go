package types

type Transfer_ServerS3StorageOptions struct {
	/*
	   Specifies whether or not performance for your Amazon S3 directories is optimized. Valid values are `DISABLED`, `ENABLED`.

	   By default, home directory mappings have a `TYPE` of `DIRECTORY`. If you enable this option, you would then need to explicitly set the `HomeDirectoryMapEntry` Type to `FILE` if you want a mapping to have a file target. See [Using logical directories to simplify your Transfer Family directory structures](https://docs.aws.amazon.com/transfer/latest/userguide/logical-dir-mappings.html) for details.
	*/
	DirectoryListingOptimization string `json:"directoryListingOptimization,omitempty" yaml:"directoryListingOptimization,omitempty"`
}
