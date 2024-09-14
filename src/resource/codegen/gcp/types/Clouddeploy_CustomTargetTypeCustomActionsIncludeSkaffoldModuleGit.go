package types

type Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModuleGit struct {
	// Relative path from the repository root to the Skaffold file.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// Git ref the package should be cloned from.
	Ref string `json:"ref,omitempty" yaml:"ref,omitempty"`

	// Git repository the package should be cloned from.
	Repo string `json:"repo,omitempty" yaml:"repo,omitempty"`
}
