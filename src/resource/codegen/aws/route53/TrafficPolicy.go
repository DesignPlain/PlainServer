package route53

type TrafficPolicy struct {
	// Comment for the traffic policy.
	Comment string `json:"comment,omitempty" yaml:"comment,omitempty"`

	/*
	   Policy document. This is a JSON formatted string. For more information about building Route53 traffic policy documents, see the [AWS Route53 Traffic Policy document format](https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html)

	   The following arguments are optional:
	*/
	Document string `json:"document,omitempty" yaml:"document,omitempty"`

	// Name of the traffic policy.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
