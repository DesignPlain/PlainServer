package types

type Networkmanager_getSiteLocation struct {
	// Address of the location.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Latitude of the location.
	Latitude string `json:"latitude,omitempty" yaml:"latitude,omitempty"`

	// Longitude of the location.
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty"`
}
