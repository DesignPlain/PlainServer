package types

type Dataproc_getMetastoreServiceHiveMetastoreConfig struct {
	/*
	   A mapping of Hive metastore version to the auxiliary version configuration.
	   When specified, a secondary Hive metastore service is created along with the primary service.
	   All auxiliary versions must be less than the service's primary version.
	   The key is the auxiliary service name and it must match the regular expression a-z?.
	   This means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
	*/
	AuxiliaryVersions []Dataproc_getMetastoreServiceHiveMetastoreConfigAuxiliaryVersion `json:"auxiliaryVersions,omitempty" yaml:"auxiliaryVersions,omitempty"`

	/*
	   A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml).
	   The mappings override system defaults (some keys cannot be overridden)
	*/
	ConfigOverrides map[string]string `json:"configOverrides,omitempty" yaml:"configOverrides,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified, defaults to 'THRIFT'. Default value: "THRIFT" Possible values: ["THRIFT", "GRPC"]
	EndpointProtocol string `json:"endpointProtocol,omitempty" yaml:"endpointProtocol,omitempty"`

	// Information used to configure the Hive metastore service as a service principal in a Kerberos realm.
	KerberosConfigs []Dataproc_getMetastoreServiceHiveMetastoreConfigKerberosConfig `json:"kerberosConfigs,omitempty" yaml:"kerberosConfigs,omitempty"`

	// The Hive metastore schema version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
