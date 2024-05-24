package amplify

type Webhook struct {
	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`

	// Name for a branch that is part of the Amplify app.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Description for a webhook.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
