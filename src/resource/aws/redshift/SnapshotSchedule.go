package redshift

type SnapshotSchedule struct {
	// The definition of the snapshot schedule. The definition is made up of schedule expressions, for example `cron(30 12 -)` or `rate(12 hours)`.
	Definitions []string `json:"definitions,omitempty" yaml:"definitions,omitempty"`

	// The description of the snapshot schedule.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Whether to destroy all associated clusters with this snapshot schedule on deletion. Must be enabled and applied before attempting deletion.
	ForceDestroy bool `json:"forceDestroy,omitempty" yaml:"forceDestroy,omitempty"`

	// The snapshot schedule identifier. If omitted, this provider will assign a random, unique identifier.
	Identifier string `json:"identifier,omitempty" yaml:"identifier,omitempty"`

	/*
	   Creates a unique
	   identifier beginning with the specified prefix. Conflicts with `identifier`.
	*/
	IdentifierPrefix string `json:"identifierPrefix,omitempty" yaml:"identifierPrefix,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
