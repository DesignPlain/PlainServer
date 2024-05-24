package types

type Elb_LoadBalancerListener struct {
	// The port on the instance to route to
	InstancePort int `json:"instancePort,omitempty" yaml:"instancePort,omitempty"`

	/*
	   The protocol to use to the instance. Valid
	   values are `HTTP`, `HTTPS`, `TCP`, or `SSL`
	*/
	InstanceProtocol string `json:"instanceProtocol,omitempty" yaml:"instanceProtocol,omitempty"`

	// The port to listen on for the load balancer
	LbPort int `json:"lbPort,omitempty" yaml:"lbPort,omitempty"`

	/*
	   The protocol to listen on. Valid values are `HTTP`,
	   `HTTPS`, `TCP`, or `SSL`
	*/
	LbProtocol string `json:"lbProtocol,omitempty" yaml:"lbProtocol,omitempty"`

	/*
	   The ARN of an SSL certificate you have
	   uploaded to AWS IAM. --Note ECDSA-specific restrictions below.  Only valid when `lb_protocol` is either HTTPS or SSL--
	*/
	SslCertificateId string `json:"sslCertificateId,omitempty" yaml:"sslCertificateId,omitempty"`
}
