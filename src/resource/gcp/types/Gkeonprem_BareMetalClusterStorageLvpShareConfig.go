package types

type Gkeonprem_BareMetalClusterStorageLvpShareConfig struct {
	/*
	   Defines the machine path and storage class for the LVP Share.
	   Structure is documented below.
	*/
	LvpConfig Gkeonprem_BareMetalClusterStorageLvpShareConfigLvpConfig `json:"lvpConfig,omitempty" yaml:"lvpConfig,omitempty"`

	// The number of subdirectories to create under path.
	SharedPathPvCount int `json:"sharedPathPvCount,omitempty" yaml:"sharedPathPvCount,omitempty"`
}
