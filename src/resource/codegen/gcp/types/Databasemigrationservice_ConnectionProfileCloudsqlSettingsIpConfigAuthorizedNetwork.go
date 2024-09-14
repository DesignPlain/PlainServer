package types

type Databasemigrationservice_ConnectionProfileCloudsqlSettingsIpConfigAuthorizedNetwork struct {
	// The time when this access control entry expires in RFC 3339 format.
	ExpireTime string `json:"expireTime,omitempty" yaml:"expireTime,omitempty"`

	// A label to identify this entry.
	Label string `json:"label,omitempty" yaml:"label,omitempty"`

	// Input only. The time-to-leave of this access control entry.
	Ttl string `json:"ttl,omitempty" yaml:"ttl,omitempty"`

	// The allowlisted value for the access control list.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
