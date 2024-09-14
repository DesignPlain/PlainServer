package types

type Dataproc_MetastoreServiceScalingConfig struct {
	/*
	   Metastore instance sizes.
	   Possible values are: `EXTRA_SMALL`, `SMALL`, `MEDIUM`, `LARGE`, `EXTRA_LARGE`.
	*/
	InstanceSize string `json:"instanceSize,omitempty" yaml:"instanceSize,omitempty"`

	// Scaling factor, in increments of 0.1 for values less than 1.0, and increments of 1.0 for values greater than 1.0.
	ScalingFactor float64 `json:"scalingFactor,omitempty" yaml:"scalingFactor,omitempty"`
}
