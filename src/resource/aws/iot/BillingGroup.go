package iot

import types "DesignSphere_Server/src/resource/aws/types"

type BillingGroup struct {
	// The name of the Billing Group.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The Billing Group properties. Defined below.
	Properties types.Iot_BillingGroupProperties `json:"properties,omitempty" yaml:"properties,omitempty"`

	// Key-value mapping of resource tags
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
