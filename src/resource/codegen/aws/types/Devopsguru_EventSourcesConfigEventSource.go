package types

type Devopsguru_EventSourcesConfigEventSource struct {
	// Stores whether DevOps Guru is configured to consume recommendations which are generated from AWS CodeGuru Profiler. See `amazon_code_guru_profiler` below.
	AmazonCodeGuruProfilers []Devopsguru_EventSourcesConfigEventSourceAmazonCodeGuruProfiler `json:"amazonCodeGuruProfilers,omitempty" yaml:"amazonCodeGuruProfilers,omitempty"`
}
