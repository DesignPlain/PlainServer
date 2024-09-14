package types

type Dns_ResponsePolicyNetwork struct {
	/*
	   The fully qualified URL of the VPC network to bind to.
	   This should be formatted like
	   `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
	*/
	NetworkUrl string `json:"networkUrl,omitempty" yaml:"networkUrl,omitempty"`
}
