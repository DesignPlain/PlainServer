package types

type Sql_getDatabaseInstanceSettingLocationPreference struct {
	// A Google App Engine application whose zone to remain in. Must be in the same region as this instance.
	FollowGaeApplication string `json:"followGaeApplication,omitempty" yaml:"followGaeApplication,omitempty"`

	// The preferred Compute Engine zone for the secondary/failover
	SecondaryZone string `json:"secondaryZone,omitempty" yaml:"secondaryZone,omitempty"`

	// The preferred compute engine zone.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
