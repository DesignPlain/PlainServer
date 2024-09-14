package types

type Sql_DatabaseInstanceSettingsInsightsConfig struct {
	// True if Query Insights will record client address when enabled.
	RecordClientAddress bool `json:"recordClientAddress,omitempty" yaml:"recordClientAddress,omitempty"`

	// True if Query Insights feature is enabled.
	QueryInsightsEnabled bool `json:"queryInsightsEnabled,omitempty" yaml:"queryInsightsEnabled,omitempty"`

	/*
	   Number of query execution plans captured by Insights per minute for all queries combined. Between 0 and 20. Default to 5.

	   The optional `settings.password_validation_policy` subblock for instances declares [Password Validation Policy](https://cloud.google.com/sql/docs/postgres/built-in-authentication) configuration. It contains:
	*/
	QueryPlansPerMinute int `json:"queryPlansPerMinute,omitempty" yaml:"queryPlansPerMinute,omitempty"`

	// Maximum query length stored in bytes. Between 256 and 4500. Default to 1024. Higher query lengths are more useful for analytical queries, but they also require more memory. Changing the query length requires you to restart the instance. You can still add tags to queries that exceed the length limit.
	QueryStringLength int `json:"queryStringLength,omitempty" yaml:"queryStringLength,omitempty"`

	// True if Query Insights will record application tags from query when enabled.
	RecordApplicationTags bool `json:"recordApplicationTags,omitempty" yaml:"recordApplicationTags,omitempty"`
}
