package types

type Sql_getDatabaseInstanceClone struct {
	// The name of the instance from which the point in time should be restored.
	SourceInstanceName string `json:"sourceInstanceName,omitempty" yaml:"sourceInstanceName,omitempty"`

	// The name of the allocated ip range for the private ip CloudSQL instance. For example: "google-managed-services-default". If set, the cloned instance ip will be created in the allocated range. The range name must comply with [RFC 1035](https://tools.ietf.org/html/rfc1035). Specifically, the name must be 1-63 characters long and match the regular expression a-z?.
	AllocatedIpRange string `json:"allocatedIpRange,omitempty" yaml:"allocatedIpRange,omitempty"`

	// (SQL Server only, use with point_in_time) clone only the specified databases from the source instance. Clone all databases if empty.
	DatabaseNames []string `json:"databaseNames,omitempty" yaml:"databaseNames,omitempty"`

	// The timestamp of the point in time that should be restored.
	PointInTime string `json:"pointInTime,omitempty" yaml:"pointInTime,omitempty"`

	// (Point-in-time recovery for PostgreSQL only) Clone to an instance in the specified zone. If no zone is specified, clone to the same zone as the source instance.
	PreferredZone string `json:"preferredZone,omitempty" yaml:"preferredZone,omitempty"`
}
