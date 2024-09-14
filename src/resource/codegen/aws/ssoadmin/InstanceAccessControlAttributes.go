package ssoadmin

import types "libds/aws/types"

type InstanceAccessControlAttributes struct {
	// See AccessControlAttribute for more details.
	Attributes []types.Ssoadmin_InstanceAccessControlAttributesAttribute `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	// The Amazon Resource Name (ARN) of the SSO Instance.
	InstanceArn string `json:"instanceArn,omitempty" yaml:"instanceArn,omitempty"`
}
