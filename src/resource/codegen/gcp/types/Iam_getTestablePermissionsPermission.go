package types

type Iam_getTestablePermissionsPermission struct {
	// Human readable title of the permission.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Whether the corresponding API has been enabled for the resource.
	ApiDisabled bool `json:"apiDisabled,omitempty" yaml:"apiDisabled,omitempty"`

	// The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
	CustomSupportLevel string `json:"customSupportLevel,omitempty" yaml:"customSupportLevel,omitempty"`

	// Name of the permission.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Release stage of the permission.
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`
}
