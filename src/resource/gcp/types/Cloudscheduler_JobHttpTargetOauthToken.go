package types

type Cloudscheduler_JobHttpTargetOauthToken struct {
	/*
	   Service account email to be used for generating OAuth token.
	   The service account must be within the same project as the job.
	*/
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`

	/*
	   OAuth scope to be used for generating OAuth access token. If not specified,
	   "https://www.googleapis.com/auth/cloud-platform" will be used.
	*/
	Scope string `json:"scope,omitempty" yaml:"scope,omitempty"`
}
