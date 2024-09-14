package types

type Drs_ReplicationConfigurationTemplatePitPolicy struct {
	// Units used to measure the `interval` and `retention_duration`. Valid values are `MINUTE`, `HOUR`, and `DAY`.
	Units string `json:"units,omitempty" yaml:"units,omitempty"`

	// Whether this rule is enabled or not.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// How often, in the chosen units, a snapshot should be taken.
	Interval int `json:"interval,omitempty" yaml:"interval,omitempty"`

	// Duration to retain a snapshot for, in the chosen `units`.
	RetentionDuration int `json:"retentionDuration,omitempty" yaml:"retentionDuration,omitempty"`

	// ID of the rule. Valid values are integers.
	RuleId int `json:"ruleId,omitempty" yaml:"ruleId,omitempty"`
}
