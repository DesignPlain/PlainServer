package types

type Medialive_ChannelDestinationMultiplexSettings struct {
	// The ID of the Multiplex that the encoder is providing output to.
	MultiplexId string `json:"multiplexId,omitempty" yaml:"multiplexId,omitempty"`

	// The program name of the Multiplex program that the encoder is providing output to.
	ProgramName string `json:"programName,omitempty" yaml:"programName,omitempty"`
}
