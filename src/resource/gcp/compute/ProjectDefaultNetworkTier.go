package compute

type ProjectDefaultNetworkTier struct {
	/*
	   The default network tier to be configured for the project.
	   This field can take the following values: `PREMIUM` or `STANDARD`.

	   - - -
	*/
	NetworkTier string `json:"networkTier,omitempty" yaml:"networkTier,omitempty"`

	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
