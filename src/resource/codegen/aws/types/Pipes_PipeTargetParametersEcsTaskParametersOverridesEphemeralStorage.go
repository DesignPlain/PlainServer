package types

type Pipes_PipeTargetParametersEcsTaskParametersOverridesEphemeralStorage struct {
	// The total amount, in GiB, of ephemeral storage to set for the task. The minimum supported value is 21 GiB and the maximum supported value is 200 GiB.
	SizeInGib int `json:"sizeInGib,omitempty" yaml:"sizeInGib,omitempty"`
}
