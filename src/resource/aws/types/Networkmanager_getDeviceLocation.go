package types

type Networkmanager_getDeviceLocation struct {
	// Physical address.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// Latitude.
	Latitude string `json:"latitude,omitempty" yaml:"latitude,omitempty"`

	// Longitude.
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty"`
}
