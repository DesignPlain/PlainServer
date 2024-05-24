package types

type Networkmanager_SiteLocation struct {
	// Longitude of the location.
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty"`

	// Address of the location.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Latitude of the location.
	Latitude string `json:"latitude,omitempty" yaml:"latitude,omitempty"`
}
