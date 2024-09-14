package types

type Cloudscheduler_JobHttpTargetOidcToken struct {
	/*
	   Audience to be used when generating OIDC token. If not specified,
	   the URI specified in target will be used.
	*/
	Audience string `json:"audience,omitempty" yaml:"audience,omitempty"`

	/*
	   Service account email to be used for generating OAuth token.
	   The service account must be within the same project as the job.
	*/
	ServiceAccountEmail string `json:"serviceAccountEmail,omitempty" yaml:"serviceAccountEmail,omitempty"`
}
