package ssm

type DefaultPatchBaseline struct {
	/*
	   ID of the patch baseline.
	   Can be an ID or an ARN.
	   When specifying an AWS-provided patch baseline, must be the ARN.
	*/
	BaselineId string `json:"baselineId,omitempty" yaml:"baselineId,omitempty"`

	/*
	   The operating system the patch baseline applies to.
	   Valid values are
	   `AMAZON_LINUX`,
	   `AMAZON_LINUX_2`,
	   `AMAZON_LINUX_2022`,
	   `CENTOS`,
	   `DEBIAN`,
	   `MACOS`,
	   `ORACLE_LINUX`,
	   `RASPBIAN`,
	   `REDHAT_ENTERPRISE_LINUX`,
	   `ROCKY_LINUX`,
	   `SUSE`,
	   `UBUNTU`, and
	   `WINDOWS`.
	*/
	OperatingSystem string `json:"operatingSystem,omitempty" yaml:"operatingSystem,omitempty"`
}
