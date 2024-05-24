package types

type Cloudfront_DistributionOriginGroupFailoverCriteria struct {
	// List of HTTP status codes for the origin group.
	StatusCodes []int `json:"statusCodes,omitempty" yaml:"statusCodes,omitempty"`
}
