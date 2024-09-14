package types

type Clouddeploy_CustomTargetTypeCustomActions struct {
	/*
	   List of Skaffold modules Cloud Deploy will include in the Skaffold Config as required before performing diagnose.
	   Structure is documented below.
	*/
	IncludeSkaffoldModules []Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModule `json:"includeSkaffoldModules,omitempty" yaml:"includeSkaffoldModules,omitempty"`

	// The Skaffold custom action responsible for render operations. If not provided then Cloud Deploy will perform the render operations via `skaffold render`.
	RenderAction string `json:"renderAction,omitempty" yaml:"renderAction,omitempty"`

	// The Skaffold custom action responsible for deploy operations.
	DeployAction string `json:"deployAction,omitempty" yaml:"deployAction,omitempty"`
}
