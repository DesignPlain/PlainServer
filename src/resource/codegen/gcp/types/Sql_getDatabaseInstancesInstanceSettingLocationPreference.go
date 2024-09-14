package types

type Sql_getDatabaseInstancesInstanceSettingLocationPreference struct {
	// A Google App Engine application whose zone to remain in. Must be in the same region as this instance.
	FollowGaeApplication string `json:"followGaeApplication,omitempty" yaml:"followGaeApplication,omitempty"`

	// The preferred Compute Engine zone for the secondary/failover
	SecondaryZone string `json:"secondaryZone,omitempty" yaml:"secondaryZone,omitempty"`

	// To filter out the Cloud SQL instances which are located in the specified zone. This zone refers to the Compute Engine zone that the instance is currently serving from.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
