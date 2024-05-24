package synthetics

import types "DesignSphere_Server/src/resource/aws/types"

type Canary struct {
	// S3 key of your script. --Conflicts with `zip_file`.--
	S3Key string `json:"s3Key,omitempty" yaml:"s3Key,omitempty"`

	// Whether to run or stop the canary.
	StartCanary bool `json:"startCanary,omitempty" yaml:"startCanary,omitempty"`

	// Number of days to retain data about successful runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	SuccessRetentionPeriod int `json:"successRetentionPeriod,omitempty" yaml:"successRetentionPeriod,omitempty"`

	// Name for this canary. Has a maximum length of 21 characters. Valid characters are lowercase alphanumeric, hyphen, or underscore.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// ARN of the IAM role to be used to run the canary. see [AWS Docs](https://docs.aws.amazon.com/AmazonSynthetics/latest/APIReference/API_CreateCanary.html#API_CreateCanary_RequestSyntax) for permissions needs for IAM Role.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Configuration block for individual canary runs. Detailed below.
	RunConfig types.Synthetics_CanaryRunConfig `json:"runConfig,omitempty" yaml:"runConfig,omitempty"`

	// Runtime version to use for the canary. Versions change often so consult the [Amazon CloudWatch documentation](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html) for the latest valid versions. Values include `syn-python-selenium-1.0`, `syn-nodejs-puppeteer-3.0`, `syn-nodejs-2.2`, `syn-nodejs-2.1`, `syn-nodejs-2.0`, and `syn-1.0`.
	RuntimeVersion string `json:"runtimeVersion,omitempty" yaml:"runtimeVersion,omitempty"`

	// Full bucket name which is used if your canary script is located in S3. The bucket must already exist. --Conflicts with `zip_file`.--
	S3Bucket string `json:"s3Bucket,omitempty" yaml:"s3Bucket,omitempty"`

	// S3 version ID of your script. --Conflicts with `zip_file`.--
	S3Version string `json:"s3Version,omitempty" yaml:"s3Version,omitempty"`

	/*
	   Configuration block providing how often the canary is to run and when these test runs are to stop. Detailed below.

	   The following arguments are optional:
	*/
	Schedule types.Synthetics_CanarySchedule `json:"schedule,omitempty" yaml:"schedule,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// configuration for canary artifacts, including the encryption-at-rest settings for artifacts that the canary uploads to Amazon S3. See Artifact Config.
	ArtifactConfig types.Synthetics_CanaryArtifactConfig `json:"artifactConfig,omitempty" yaml:"artifactConfig,omitempty"`

	// Entry point to use for the source code when running the canary. This value must end with the string `.handler` .
	Handler string `json:"handler,omitempty" yaml:"handler,omitempty"`

	// Configuration block. Detailed below.
	VpcConfig types.Synthetics_CanaryVpcConfig `json:"vpcConfig,omitempty" yaml:"vpcConfig,omitempty"`

	// Specifies whether to also delete the Lambda functions and layers used by this canary. The default is `false`.
	DeleteLambda bool `json:"deleteLambda,omitempty" yaml:"deleteLambda,omitempty"`

	// Number of days to retain data about failed runs of this canary. If you omit this field, the default of 31 days is used. The valid range is 1 to 455 days.
	FailureRetentionPeriod int `json:"failureRetentionPeriod,omitempty" yaml:"failureRetentionPeriod,omitempty"`

	// ZIP file that contains the script, if you input your canary script directly into the canary instead of referring to an S3 location. It can be up to 225KB. --Conflicts with `s3_bucket`, `s3_key`, and `s3_version`.--
	ZipFile string `json:"zipFile,omitempty" yaml:"zipFile,omitempty"`

	// Location in Amazon S3 where Synthetics stores artifacts from the test runs of this canary.
	ArtifactS3Location string `json:"artifactS3Location,omitempty" yaml:"artifactS3Location,omitempty"`
}
