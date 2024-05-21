package types

type Sql_getDatabaseInstancesInstanceSettingInsightsConfig struct {
	// True if Query Insights feature is enabled.
	QueryInsightsEnabled bool `json:"queryInsightsEnabled,omitempty" yaml:"queryInsightsEnabled,omitempty"`

	// Number of query execution plans captured by Insights per minute for all queries combined. Between 0 and 20. Default to 5.
	QueryPlansPerMinute int `json:"queryPlansPerMinute,omitempty" yaml:"queryPlansPerMinute,omitempty"`

	// Maximum query length stored in bytes. Between 256 and 4500. Default to 1024.
	QueryStringLength int `json:"queryStringLength,omitempty" yaml:"queryStringLength,omitempty"`

	// True if Query Insights will record application tags from query when enabled.
	RecordApplicationTags bool `json:"recordApplicationTags,omitempty" yaml:"recordApplicationTags,omitempty"`

	// True if Query Insights will record client address when enabled.
	RecordClientAddress bool `json:"recordClientAddress,omitempty" yaml:"recordClientAddress,omitempty"`
}
