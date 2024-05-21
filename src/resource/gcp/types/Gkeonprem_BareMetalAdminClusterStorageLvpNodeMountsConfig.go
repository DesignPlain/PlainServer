package types

type Gkeonprem_BareMetalAdminClusterStorageLvpNodeMountsConfig struct {
	// The host machine path.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The StorageClass name that PVs will be created with.
	StorageClass string `json:"storageClass,omitempty" yaml:"storageClass,omitempty"`
}
