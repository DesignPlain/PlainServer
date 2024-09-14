package types

type Cloudbuild_TriggerBuildArtifactsObjects struct {
	/*
	   Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".
	   Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
	   this location as a prefix.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// Path globs used to match files in the build's workspace.
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`

	/*
	   (Output)
	   Output only. Stores timing information for pushing all artifact objects.
	   Structure is documented below.


	   <a name="nested_timing"></a>The `timing` block contains:
	*/
	Timings []Cloudbuild_TriggerBuildArtifactsObjectsTiming `json:"timings,omitempty" yaml:"timings,omitempty"`
}
