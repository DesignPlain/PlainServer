package types

type Cloudrun_getServiceTraffic struct {
	/*
	   LatestRevision may be optionally provided to indicate that the latest ready
	   Revision of the Configuration should be used for this traffic target. When
	   provided LatestRevision must be true if RevisionName is empty; it must be
	   false when RevisionName is non-empty.
	*/
	LatestRevision bool `json:"latestRevision,omitempty" yaml:"latestRevision,omitempty"`

	// Percent specifies percent of the traffic to this Revision or Configuration.
	Percent int `json:"percent,omitempty" yaml:"percent,omitempty"`

	// RevisionName of a specific revision to which to send this portion of traffic.
	RevisionName string `json:"revisionName,omitempty" yaml:"revisionName,omitempty"`

	// Tag is optionally used to expose a dedicated url for referencing this target exclusively.
	Tag string `json:"tag,omitempty" yaml:"tag,omitempty"`

	/*
	   URL displays the URL for accessing tagged traffic targets. URL is displayed in status,
	   and is disallowed on spec. URL must contain a scheme (e.g. http://) and a hostname,
	   but may not contain anything else (e.g. basic auth, url path, etc.)
	*/
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
