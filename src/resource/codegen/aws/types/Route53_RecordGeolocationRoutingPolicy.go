package types

type Route53_RecordGeolocationRoutingPolicy struct {
	// A two-letter continent code. See http://docs.aws.amazon.com/Route53/latest/APIReference/API_GetGeoLocation.html for code details. Either `continent` or `country` must be specified.
	Continent string `json:"continent,omitempty" yaml:"continent,omitempty"`

	// A two-character country code or `-` to indicate a default resource record set.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`

	// A subdivision code for a country.
	Subdivision string `json:"subdivision,omitempty" yaml:"subdivision,omitempty"`
}
