package types

type Container_AwsClusterControlPlaneInstancePlacement struct {
	// The tenancy for the instance. Possible values: TENANCY_UNSPECIFIED, DEFAULT, DEDICATED, HOST
	Tenancy string `json:"tenancy,omitempty" yaml:"tenancy,omitempty"`
}
