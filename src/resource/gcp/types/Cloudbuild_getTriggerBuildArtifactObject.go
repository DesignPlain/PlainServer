package types

type Cloudbuild_getTriggerBuildArtifactObject struct {
	/*
	   The Cloud Build location for the trigger.

	   - - -
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Path globs used to match files in the build's workspace.
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`

	// Output only. Stores timing information for pushing all artifact objects.
	Timings []Cloudbuild_getTriggerBuildArtifactObjectTiming `json:"timings,omitempty" yaml:"timings,omitempty"`
}
