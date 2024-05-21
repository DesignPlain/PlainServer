package types

type Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModule struct {
	// The Skaffold Config modules to use from the specified source.
	Configs []string `json:"configs,omitempty" yaml:"configs,omitempty"`

	/*
	   Remote git repository containing the Skaffold Config modules.
	   Structure is documented below.
	*/
	Git Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModuleGit `json:"git,omitempty" yaml:"git,omitempty"`

	/*
	   Cloud Storage bucket containing Skaffold Config modules.
	   Structure is documented below.
	*/
	GoogleCloudStorage Clouddeploy_CustomTargetTypeCustomActionsIncludeSkaffoldModuleGoogleCloudStorage `json:"googleCloudStorage,omitempty" yaml:"googleCloudStorage,omitempty"`
}
