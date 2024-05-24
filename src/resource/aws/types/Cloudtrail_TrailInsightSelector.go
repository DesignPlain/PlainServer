package types

type Cloudtrail_TrailInsightSelector struct {
	// Type of insights to log on a trail. Valid values are: `ApiCallRateInsight` and `ApiErrorRateInsight`.
	InsightType string `json:"insightType,omitempty" yaml:"insightType,omitempty"`
}
