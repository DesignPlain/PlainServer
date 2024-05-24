package types

type Dynamodb_TablePointInTimeRecovery struct {
	// Whether to enable point-in-time recovery. It can take 10 minutes to enable for new tables. If the `point_in_time_recovery` block is not provided, this defaults to `false`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
