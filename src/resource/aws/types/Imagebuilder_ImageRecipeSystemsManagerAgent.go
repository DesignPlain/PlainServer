package types

type Imagebuilder_ImageRecipeSystemsManagerAgent struct {
	// Whether to remove the Systems Manager Agent after the image has been built. Defaults to `false`.
	UninstallAfterBuild bool `json:"uninstallAfterBuild,omitempty" yaml:"uninstallAfterBuild,omitempty"`
}
