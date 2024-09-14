package types

type Workspaces_getBundleComputeType struct {
	// Name of the bundle. You cannot combine this parameter with `bundle_id`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
