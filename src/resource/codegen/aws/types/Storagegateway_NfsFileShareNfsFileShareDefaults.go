package types

type Storagegateway_NfsFileShareNfsFileShareDefaults struct {
	// The Unix directory mode in the string form "nnnn". Defaults to `"0777"`.
	DirectoryMode string `json:"directoryMode,omitempty" yaml:"directoryMode,omitempty"`

	// The Unix file mode in the string form "nnnn". Defaults to `"0666"`.
	FileMode string `json:"fileMode,omitempty" yaml:"fileMode,omitempty"`

	// The default group ID for the file share (unless the files have another group ID specified). Defaults to `65534` (`nfsnobody`). Valid values: `0` through `4294967294`.
	GroupId string `json:"groupId,omitempty" yaml:"groupId,omitempty"`

	// The default owner ID for the file share (unless the files have another owner ID specified). Defaults to `65534` (`nfsnobody`). Valid values: `0` through `4294967294`.
	OwnerId string `json:"ownerId,omitempty" yaml:"ownerId,omitempty"`
}
