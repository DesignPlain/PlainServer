package ssm

type PatchGroup struct {
	// The ID of the patch baseline to register the patch group with.
	BaselineId string `json:"baselineId,omitempty" yaml:"baselineId,omitempty"`

	// The name of the patch group that should be registered with the patch baseline.
	PatchGroup string `json:"patchGroup,omitempty" yaml:"patchGroup,omitempty"`
}
