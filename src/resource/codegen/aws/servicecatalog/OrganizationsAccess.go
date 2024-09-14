package servicecatalog

type OrganizationsAccess struct {
	// Whether to enable AWS Organizations access.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
