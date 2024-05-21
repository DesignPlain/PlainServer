package types

type Netapp_VolumeReplicationDestinationVolumeParameters struct {
	// Description for the destination volume.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Share name for destination volume. If not specified, name of source volume's share name will be used.
	ShareName string `json:"shareName,omitempty" yaml:"shareName,omitempty"`

	// Name of an existing storage pool for the destination volume with format: `projects/{{project}}/locations/{{location}}/storagePools/{{poolId}}`
	StoragePool string `json:"storagePool,omitempty" yaml:"storagePool,omitempty"`

	// Name for the destination volume to be created. If not specified, the name of the source volume will be used.
	VolumeId string `json:"volumeId,omitempty" yaml:"volumeId,omitempty"`
}
