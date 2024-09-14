package amplify

type Webhook struct {
	// Name for a branch that is part of the Amplify app.
	BranchName string `json:"branchName,omitempty" yaml:"branchName,omitempty"`

	// Description for a webhook.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Unique ID for an Amplify app.
	AppId string `json:"appId,omitempty" yaml:"appId,omitempty"`
}
