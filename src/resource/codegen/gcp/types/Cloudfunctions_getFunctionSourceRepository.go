package types

type Cloudfunctions_getFunctionSourceRepository struct {
	// The URL pointing to the hosted repository where the function was defined at the time of deployment.
	DeployedUrl string `json:"deployedUrl,omitempty" yaml:"deployedUrl,omitempty"`

	// The URL pointing to the hosted repository where the function is defined.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
