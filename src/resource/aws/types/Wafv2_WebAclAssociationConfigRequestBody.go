package types

type Wafv2_WebAclAssociationConfigRequestBody struct {
	// Customizes the request body that your protected CloudFront distributions forward to AWS WAF for inspection. See `cloudfront` below for details.
	Cloudfronts []Wafv2_WebAclAssociationConfigRequestBodyCloudfront `json:"cloudfronts,omitempty" yaml:"cloudfronts,omitempty"`
}
