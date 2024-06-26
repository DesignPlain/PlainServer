package types

type Certificatemanager_getCertificateMapGclbTarget struct {
	// An IP configuration where this Certificate Map is serving
	IpConfigs []Certificatemanager_getCertificateMapGclbTargetIpConfig `json:"ipConfigs,omitempty" yaml:"ipConfigs,omitempty"`

	/*
	   Proxy name must be in the format projects/-/locations/-/targetHttpsProxies/-.
	   This field is part of a union field 'target_proxy': Only one of 'targetHttpsProxy' or
	   'targetSslProxy' may be set.
	*/
	TargetHttpsProxy string `json:"targetHttpsProxy,omitempty" yaml:"targetHttpsProxy,omitempty"`

	/*
	   Proxy name must be in the format projects/-/locations/-/targetSslProxies/-.
	   This field is part of a union field 'target_proxy': Only one of 'targetHttpsProxy' or
	   'targetSslProxy' may be set.
	*/
	TargetSslProxy string `json:"targetSslProxy,omitempty" yaml:"targetSslProxy,omitempty"`
}
