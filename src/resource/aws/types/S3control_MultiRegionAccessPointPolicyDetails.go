package types

type S3control_MultiRegionAccessPointPolicyDetails struct {
	// The name of the Multi-Region Access Point.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A valid JSON document that specifies the policy that you want to associate with this Multi-Region Access Point. Once applied, the policy can be edited, but not deleted. For more information, see the documentation on [Multi-Region Access Point Permissions](https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointPermissions.html).

	   > --NOTE:-- When you update the `policy`, the update is first listed as the proposed policy. After the update is finished and all Regions have been updated, the proposed policy is listed as the established policy. If both policies have the same version number, the proposed policy is the established policy.
	*/
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
