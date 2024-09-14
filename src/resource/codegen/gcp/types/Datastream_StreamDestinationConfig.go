package types

type Datastream_StreamDestinationConfig struct {
	/*
	   A configuration for how data should be loaded to Cloud Storage.
	   Structure is documented below.
	*/
	BigqueryDestinationConfig Datastream_StreamDestinationConfigBigqueryDestinationConfig `json:"bigqueryDestinationConfig,omitempty" yaml:"bigqueryDestinationConfig,omitempty"`

	// Destination connection profile resource. Format: projects/{project}/locations/{location}/connectionProfiles/{name}
	DestinationConnectionProfile string `json:"destinationConnectionProfile,omitempty" yaml:"destinationConnectionProfile,omitempty"`

	/*
	   A configuration for how data should be loaded to Cloud Storage.
	   Structure is documented below.
	*/
	GcsDestinationConfig Datastream_StreamDestinationConfigGcsDestinationConfig `json:"gcsDestinationConfig,omitempty" yaml:"gcsDestinationConfig,omitempty"`
}
