package types

type Wafv2_WebAclAssociationConfigRequestBodyVerifiedAccessInstance struct {
	// Specifies the maximum size of the web request body component that an associated AWS Verified Access instances should send to AWS WAF for inspection. This applies to statements in the web ACL that inspect the body or JSON body. Valid values are `KB_16`, `KB_32`, `KB_48` and `KB_64`.
	DefaultSizeInspectionLimit string `json:"defaultSizeInspectionLimit,omitempty" yaml:"defaultSizeInspectionLimit,omitempty"`
}
