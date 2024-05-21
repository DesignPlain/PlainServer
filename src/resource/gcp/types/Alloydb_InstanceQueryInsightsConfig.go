package types

type Alloydb_InstanceQueryInsightsConfig struct {
	// Number of query execution plans captured by Insights per minute for all queries combined. The default value is 5. Any integer between 0 and 20 is considered valid.
	QueryPlansPerMinute int `json:"queryPlansPerMinute,omitempty" yaml:"queryPlansPerMinute,omitempty"`

	// Query string length. The default value is 1024. Any integer between 256 and 4500 is considered valid.
	QueryStringLength int `json:"queryStringLength,omitempty" yaml:"queryStringLength,omitempty"`

	// Record application tags for an instance. This flag is turned "on" by default.
	RecordApplicationTags bool `json:"recordApplicationTags,omitempty" yaml:"recordApplicationTags,omitempty"`

	// Record client address for an instance. Client address is PII information. This flag is turned "on" by default.
	RecordClientAddress bool `json:"recordClientAddress,omitempty" yaml:"recordClientAddress,omitempty"`
}
