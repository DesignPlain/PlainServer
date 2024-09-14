package types

type Dataplex_AssetDiscoveryStatusStat struct {
	// The count of data items within the referenced resource.
	DataItems int `json:"dataItems,omitempty" yaml:"dataItems,omitempty"`

	// The number of stored data bytes within the referenced resource.
	DataSize int `json:"dataSize,omitempty" yaml:"dataSize,omitempty"`

	// The count of fileset entities within the referenced resource.
	Filesets int `json:"filesets,omitempty" yaml:"filesets,omitempty"`

	// The count of table entities within the referenced resource.
	Tables int `json:"tables,omitempty" yaml:"tables,omitempty"`
}
