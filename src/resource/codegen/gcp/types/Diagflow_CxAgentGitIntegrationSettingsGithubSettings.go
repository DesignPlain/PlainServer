package types

type Diagflow_CxAgentGitIntegrationSettingsGithubSettings struct {
	// The unique repository display name for the GitHub repository.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// The GitHub repository URI related to the agent.
	RepositoryUri string `json:"repositoryUri,omitempty" yaml:"repositoryUri,omitempty"`

	// The branch of the GitHub repository tracked for this agent.
	TrackingBranch string `json:"trackingBranch,omitempty" yaml:"trackingBranch,omitempty"`

	/*
	   The access token used to authenticate the access to the GitHub repository.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	AccessToken string `json:"accessToken,omitempty" yaml:"accessToken,omitempty"`

	// A list of branches configured to be used from Dialogflow.
	Branches []string `json:"branches,omitempty" yaml:"branches,omitempty"`
}
