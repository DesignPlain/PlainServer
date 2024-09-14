package types

type Route53_getTrafficPolicyDocumentRule struct {
	// ID of a rule you want to assign.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Configuration block for when you add a multivalue answer rule, you configure your traffic policy to route traffic approximately randomly to your healthy resources.  Only valid for `multivalue` type. See below
	Items []Route53_getTrafficPolicyDocumentRuleItem `json:"items,omitempty" yaml:"items,omitempty"`

	// Configuration block for when you add a geolocation rule, you configure your traffic policy to route your traffic based on the geographic location of your users.  Only valid for `geo` type. See below
	Locations []Route53_getTrafficPolicyDocumentRuleLocation `json:"locations,omitempty" yaml:"locations,omitempty"`

	// Configuration block for the settings for the rule or endpoint that you want to route traffic to whenever the corresponding resources are available. Only valid for `failover` type. See below
	Primary Route53_getTrafficPolicyDocumentRulePrimary `json:"primary,omitempty" yaml:"primary,omitempty"`

	//
	Regions []Route53_getTrafficPolicyDocumentRuleRegion `json:"regions,omitempty" yaml:"regions,omitempty"`

	// Configuration block for the rule or endpoint that you want to route traffic to whenever the primary resources are not available. Only valid for `failover` type. See below
	Secondary Route53_getTrafficPolicyDocumentRuleSecondary `json:"secondary,omitempty" yaml:"secondary,omitempty"`

	// Type of the rule.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Configuration block for when you add a geoproximity rule, you configure Amazon Route 53 to route traffic to your resources based on the geographic location of your resources. Only valid for `geoproximity` type. See below
	GeoProximityLocations []Route53_getTrafficPolicyDocumentRuleGeoProximityLocation `json:"geoProximityLocations,omitempty" yaml:"geoProximityLocations,omitempty"`
}
