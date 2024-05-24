package types

type Servicecatalog_getProvisioningArtifactsProvisioningArtifactDetail struct {
	// The identifier of the provisioning artifact.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// The name of the provisioning artifact.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of provisioning artifact.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Indicates whether the product version is active.
	Active bool `json:"active,omitempty" yaml:"active,omitempty"`

	// The UTC time stamp of the creation time.
	CreatedTime string `json:"createdTime,omitempty" yaml:"createdTime,omitempty"`

	// The description of the provisioning artifact.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Information set by the administrator to provide guidance to end users about which provisioning artifacts to use.
	Guidance string `json:"guidance,omitempty" yaml:"guidance,omitempty"`
}
