package directoryservice

type SharedDirectoryAccepter struct {
	// Identifier of the directory that is stored in the directory consumer account that corresponds to the shared directory in the owner account.
	SharedDirectoryId string `json:"sharedDirectoryId,omitempty" yaml:"sharedDirectoryId,omitempty"`
}
