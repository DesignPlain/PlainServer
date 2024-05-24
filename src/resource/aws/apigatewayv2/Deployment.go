package apigatewayv2

type Deployment struct {
	// API identifier.
	ApiId string `json:"apiId,omitempty" yaml:"apiId,omitempty"`

	// Description for the deployment resource. Must be less than or equal to 1024 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Map of arbitrary keys and values that, when changed, will trigger a redeployment.
	Triggers map[string]string `json:"triggers,omitempty" yaml:"triggers,omitempty"`
}
