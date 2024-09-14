package types

type Elb_getLoadBalancerListener struct {
	//
	LbProtocol string `json:"lbProtocol,omitempty" yaml:"lbProtocol,omitempty"`

	//
	SslCertificateId string `json:"sslCertificateId,omitempty" yaml:"sslCertificateId,omitempty"`

	//
	InstancePort int `json:"instancePort,omitempty" yaml:"instancePort,omitempty"`

	//
	InstanceProtocol string `json:"instanceProtocol,omitempty" yaml:"instanceProtocol,omitempty"`

	//
	LbPort int `json:"lbPort,omitempty" yaml:"lbPort,omitempty"`
}
