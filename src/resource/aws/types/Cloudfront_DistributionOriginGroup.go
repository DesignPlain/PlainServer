package types

type Cloudfront_DistributionOriginGroup struct {
	// The failover criteria for when to failover to the secondary origin.
	FailoverCriteria Cloudfront_DistributionOriginGroupFailoverCriteria `json:"failoverCriteria,omitempty" yaml:"failoverCriteria,omitempty"`

	// Ordered member configuration blocks assigned to the origin group, where the first member is the primary origin. You must specify two members.
	Members []Cloudfront_DistributionOriginGroupMember `json:"members,omitempty" yaml:"members,omitempty"`

	// Unique identifier of the member origin.
	OriginId string `json:"originId,omitempty" yaml:"originId,omitempty"`
}
