package types

type Compute_ImageRawDisk struct {
	/*
	   The full Google Cloud Storage URL where disk storage is stored
	   You must provide either this property or the sourceDisk property
	   but not both.
	*/
	Source string `json:"source,omitempty" yaml:"source,omitempty"`

	/*
	   The format used to encode and transmit the block device, which
	   should be TAR. This is just a container and transmission format
	   and not a runtime format. Provided by the client when the disk
	   image is created.
	   Default value is `TAR`.
	   Possible values are: `TAR`.
	*/
	ContainerType string `json:"containerType,omitempty" yaml:"containerType,omitempty"`

	/*
	   An optional SHA1 checksum of the disk image before unpackaging.
	   This is provided by the client when the disk image is created.
	*/
	Sha1 string `json:"sha1,omitempty" yaml:"sha1,omitempty"`
}
