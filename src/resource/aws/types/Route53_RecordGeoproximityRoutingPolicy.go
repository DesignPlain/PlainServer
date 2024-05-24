package types

type Route53_RecordGeoproximityRoutingPolicy struct {
	// Specify `latitude` and `longitude` for routing traffic to non-AWS resources.
	Coordinates []Route53_RecordGeoproximityRoutingPolicyCoordinate `json:"coordinates,omitempty" yaml:"coordinates,omitempty"`

	// A AWS local zone group where the resource is present. See https://docs.aws.amazon.com/local-zones/latest/ug/available-local-zones.html for local zone group list.
	LocalZoneGroup string `json:"localZoneGroup,omitempty" yaml:"localZoneGroup,omitempty"`

	// A AWS region where the resource is present.
	AwsRegion string `json:"awsRegion,omitempty" yaml:"awsRegion,omitempty"`

	// Route more traffic or less traffic to the resource by specifying a value ranges between -90 to 90. See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy-geoproximity.html for bias details.
	Bias int `json:"bias,omitempty" yaml:"bias,omitempty"`
}
