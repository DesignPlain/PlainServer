package rolesanywhere

type Profile struct {
	// A list of IAM roles that this profile can assume
	RoleArns []string `json:"roleArns,omitempty" yaml:"roleArns,omitempty"`

	// A session policy that applies to the trust boundary of the vended session credentials.
	SessionPolicy string `json:"sessionPolicy,omitempty" yaml:"sessionPolicy,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The number of seconds the vended session credentials are valid for. Defaults to 3600.
	DurationSeconds int `json:"durationSeconds,omitempty" yaml:"durationSeconds,omitempty"`

	// Whether or not the Profile is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A list of managed policy ARNs that apply to the vended session credentials.
	ManagedPolicyArns []string `json:"managedPolicyArns,omitempty" yaml:"managedPolicyArns,omitempty"`

	// The name of the Profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Specifies whether instance properties are required in [CreateSession](https://docs.aws.amazon.com/rolesanywhere/latest/APIReference/API_CreateSession.html) requests with this profile.
	RequireInstanceProperties bool `json:"requireInstanceProperties,omitempty" yaml:"requireInstanceProperties,omitempty"`
}
