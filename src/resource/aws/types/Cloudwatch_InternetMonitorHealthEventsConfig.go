package types

type Cloudwatch_InternetMonitorHealthEventsConfig struct {
	// The health event threshold percentage set for performance scores.
	PerformanceScoreThreshold float64 `json:"performanceScoreThreshold,omitempty" yaml:"performanceScoreThreshold,omitempty"`

	// The health event threshold percentage set for availability scores.
	AvailabilityScoreThreshold float64 `json:"availabilityScoreThreshold,omitempty" yaml:"availabilityScoreThreshold,omitempty"`
}
