package medialive

import types "libds/aws/types"

type MultiplexProgram struct {
	// Multiplex ID.
	MultiplexId string `json:"multiplexId,omitempty" yaml:"multiplexId,omitempty"`

	/*
	   MultiplexProgram settings. See Multiplex Program Settings for more details.

	   The following arguments are optional:
	*/
	MultiplexProgramSettings types.Medialive_MultiplexProgramMultiplexProgramSettings `json:"multiplexProgramSettings,omitempty" yaml:"multiplexProgramSettings,omitempty"`

	// Unique program name.
	ProgramName string `json:"programName,omitempty" yaml:"programName,omitempty"`
}
