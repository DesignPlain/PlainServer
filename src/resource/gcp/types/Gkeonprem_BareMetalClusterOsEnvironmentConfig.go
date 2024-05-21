package types

type Gkeonprem_BareMetalClusterOsEnvironmentConfig struct {
	/*
	   Whether the package repo should not be included when initializing
	   bare metal machines.
	*/
	PackageRepoExcluded bool `json:"packageRepoExcluded,omitempty" yaml:"packageRepoExcluded,omitempty"`
}
