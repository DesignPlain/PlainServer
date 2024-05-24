package types

type Cloudfront_DistributionRestrictionsGeoRestriction struct {
	// Method that you want to use to restrict distribution of your content by country: `none`, `whitelist`, or `blacklist`.
	RestrictionType string `json:"restrictionType,omitempty" yaml:"restrictionType,omitempty"`

	// [ISO 3166-1-alpha-2 codes][4] for which you want CloudFront either to distribute your content (`whitelist`) or not distribute your content (`blacklist`). If the type is specified as `none` an empty array can be used.
	Locations []string `json:"locations,omitempty" yaml:"locations,omitempty"`
}
