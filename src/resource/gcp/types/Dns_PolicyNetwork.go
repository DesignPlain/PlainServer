package types

type Dns_PolicyNetwork struct {
	/*
	   The id or fully qualified URL of the VPC network to forward queries to.
	   This should be formatted like `projects/{project}/global/networks/{network}` or
	   `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`
	*/
	NetworkUrl string `json:"networkUrl,omitempty" yaml:"networkUrl,omitempty"`
}
