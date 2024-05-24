package types

type Lightsail_ContainerServicePublicDomainNamesCertificate struct {
	//
	CertificateName string `json:"certificateName,omitempty" yaml:"certificateName,omitempty"`

	//
	DomainNames []string `json:"domainNames,omitempty" yaml:"domainNames,omitempty"`
}
