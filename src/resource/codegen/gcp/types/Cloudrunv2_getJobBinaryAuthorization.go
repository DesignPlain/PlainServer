package types

type Cloudrunv2_getJobBinaryAuthorization struct {
	// If present, indicates to use Breakglass using this justification. If useDefault is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass
	BreakglassJustification string `json:"breakglassJustification,omitempty" yaml:"breakglassJustification,omitempty"`

	// If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.
	UseDefault bool `json:"useDefault,omitempty" yaml:"useDefault,omitempty"`
}
