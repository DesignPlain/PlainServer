package types

type Medialive_ChannelDestination struct {
	// User-specified id. Ths is used in an output group or an output.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Destination settings for a MediaPackage output; one destination for both encoders. See Media Package Settings for more details.
	MediaPackageSettings []Medialive_ChannelDestinationMediaPackageSetting `json:"mediaPackageSettings,omitempty" yaml:"mediaPackageSettings,omitempty"`

	// Destination settings for a Multiplex output; one destination for both encoders. See Multiplex Settings for more details.
	MultiplexSettings Medialive_ChannelDestinationMultiplexSettings `json:"multiplexSettings,omitempty" yaml:"multiplexSettings,omitempty"`

	// Destination settings for a standard output; one destination for each redundant encoder. See Settings for more details.
	Settings []Medialive_ChannelDestinationSetting `json:"settings,omitempty" yaml:"settings,omitempty"`
}
