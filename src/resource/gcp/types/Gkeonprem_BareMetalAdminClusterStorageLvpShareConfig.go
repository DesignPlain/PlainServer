package types

type Gkeonprem_BareMetalAdminClusterStorageLvpShareConfig struct {
	// The number of subdirectories to create under path.
	SharedPathPvCount int `json:"sharedPathPvCount,omitempty" yaml:"sharedPathPvCount,omitempty"`

	/*
	   Defines the machine path and storage class for the LVP Share.
	   Structure is documented below.
	*/
	LvpConfig Gkeonprem_BareMetalAdminClusterStorageLvpShareConfigLvpConfig `json:"lvpConfig,omitempty" yaml:"lvpConfig,omitempty"`
}
