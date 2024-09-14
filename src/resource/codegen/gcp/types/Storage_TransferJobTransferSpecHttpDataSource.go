package types

type Storage_TransferJobTransferSpecHttpDataSource struct {
	// The URL that points to the file that stores the object list entries. This file must allow public access. Currently, only URLs with HTTP and HTTPS schemes are supported.
	ListUrl string `json:"listUrl,omitempty" yaml:"listUrl,omitempty"`
}
