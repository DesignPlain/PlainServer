package types

type Codebuild_ProjectEnvironment struct {
	// Whether to enable running the Docker daemon inside a Docker container. Defaults to `false`.
	PrivilegedMode bool `json:"privilegedMode,omitempty" yaml:"privilegedMode,omitempty"`

	// Configuration block. Detailed below.
	RegistryCredential Codebuild_ProjectEnvironmentRegistryCredential `json:"registryCredential,omitempty" yaml:"registryCredential,omitempty"`

	// Type of environment variable. Valid values: `PARAMETER_STORE`, `PLAINTEXT`, `SECRETS_MANAGER`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// ARN of the S3 bucket, path prefix and object key that contains the PEM-encoded certificate.
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`

	// Information about the compute resources the build project will use. Valid values: `BUILD_GENERAL1_SMALL`, `BUILD_GENERAL1_MEDIUM`, `BUILD_GENERAL1_LARGE`, `BUILD_GENERAL1_2XLARGE`, `BUILD_LAMBDA_1GB`, `BUILD_LAMBDA_2GB`, `BUILD_LAMBDA_4GB`, `BUILD_LAMBDA_8GB`, `BUILD_LAMBDA_10GB`. `BUILD_GENERAL1_SMALL` is only valid if `type` is set to `LINUX_CONTAINER`. When `type` is set to `LINUX_GPU_CONTAINER`, `compute_type` must be `BUILD_GENERAL1_LARGE`. When `type` is set to `LINUX_LAMBDA_CONTAINER` or `ARM_LAMBDA_CONTAINER`, `compute_type` must be `BUILD_LAMBDA_XGB`.`
	ComputeType string `json:"computeType,omitempty" yaml:"computeType,omitempty"`

	// Configuration block. Detailed below.
	EnvironmentVariables []Codebuild_ProjectEnvironmentEnvironmentVariable `json:"environmentVariables,omitempty" yaml:"environmentVariables,omitempty"`

	// Docker image to use for this build project. Valid values include Docker images provided by CodeBuild, and full Docker repository URIs such as those for ECR (e.g., `137112412989.dkr.ecr.us-west-2.amazonaws.com/amazonlinux:latest`).
	Image string `json:"image,omitempty" yaml:"image,omitempty"`

	// Type of credentials AWS CodeBuild uses to pull images in your build. Valid values: `CODEBUILD`, `SERVICE_ROLE`. When you use a cross-account or private registry image, you must use SERVICE_ROLE credentials. When you use an AWS CodeBuild curated image, you must use CodeBuild credentials. Defaults to `CODEBUILD`.
	ImagePullCredentialsType string `json:"imagePullCredentialsType,omitempty" yaml:"imagePullCredentialsType,omitempty"`
}
