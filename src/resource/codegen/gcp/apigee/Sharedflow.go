package apigee

type Sharedflow struct {
	/*
	   Path to the config zip bundle.

	   - - -
	*/
	ConfigBundle string `json:"configBundle,omitempty" yaml:"configBundle,omitempty"`

	/*
	   A hash of local config bundle in string, user needs to use a Terraform Hash function of their choice. A change in hash
	   will trigger an update.
	*/
	DetectMd5hash string `json:"detectMd5hash,omitempty" yaml:"detectMd5hash,omitempty"`

	// The ID of the shared flow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Apigee Organization name associated with the Apigee instance.
	OrgId string `json:"orgId,omitempty" yaml:"orgId,omitempty"`
}
