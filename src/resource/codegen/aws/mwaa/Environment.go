package mwaa

import types "libds/aws/types"

type Environment struct {
	// The minimum number of web servers that you want to run in your environment. Value need to be between `2` and `5`. Will be `2` by default.
	MinWebservers int `json:"minWebservers,omitempty" yaml:"minWebservers,omitempty"`

	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion string `json:"airflowVersion,omitempty" yaml:"airflowVersion,omitempty"`

	// Defines whether the VPC endpoints configured for the environment are created and managed by the customer or by AWS. If set to `SERVICE`, Amazon MWAA will create and manage the required VPC endpoints in your VPC. If set to `CUSTOMER`, you must create, and manage, the VPC endpoints for your VPC. Defaults to `SERVICE` if not set.
	EndpointManagement string `json:"endpointManagement,omitempty" yaml:"endpointManagement,omitempty"`

	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass string `json:"environmentClass,omitempty" yaml:"environmentClass,omitempty"`

	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey string `json:"kmsKey,omitempty" yaml:"kmsKey,omitempty"`

	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs. See `logging_configuration` Block for details.
	LoggingConfiguration types.Mwaa_EnvironmentLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`

	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers int `json:"maxWorkers,omitempty" yaml:"maxWorkers,omitempty"`

	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then plugins_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path string `json:"pluginsS3Path,omitempty" yaml:"pluginsS3Path,omitempty"`

	// The version of the startup shell script you want to use. You must specify the version ID that Amazon S3 assigns to the file every time you update the script.
	StartupScriptS3ObjectVersion string `json:"startupScriptS3ObjectVersion,omitempty" yaml:"startupScriptS3ObjectVersion,omitempty"`

	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers int `json:"minWorkers,omitempty" yaml:"minWorkers,omitempty"`

	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See `network_configuration` Block for details.
	NetworkConfiguration types.Mwaa_EnvironmentNetworkConfiguration `json:"networkConfiguration,omitempty" yaml:"networkConfiguration,omitempty"`

	// The number of schedulers that you want to run in your environment. v2.0.2 and above accepts `2` - `5`, default `2`. v1.10.12 accepts `1`.
	Schedulers int `json:"schedulers,omitempty" yaml:"schedulers,omitempty"`

	// The relative path to the script hosted in your bucket. The script runs as your environment starts before starting the Apache Airflow process. Use this script to install dependencies, modify configuration options, and set environment variables. See [Using a startup script](https://docs.aws.amazon.com/mwaa/latest/userguide/using-startup-script.html). Supported for environment versions 2.x and later.
	StartupScriptS3Path string `json:"startupScriptS3Path,omitempty" yaml:"startupScriptS3Path,omitempty"`

	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path string `json:"dagS3Path,omitempty" yaml:"dagS3Path,omitempty"`

	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion string `json:"requirementsS3ObjectVersion,omitempty" yaml:"requirementsS3ObjectVersion,omitempty"`

	// A map of resource tags to associate with the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart string `json:"weeklyMaintenanceWindowStart,omitempty" yaml:"weeklyMaintenanceWindowStart,omitempty"`

	// The `airflow_configuration_options` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions map[string]string `json:"airflowConfigurationOptions,omitempty" yaml:"airflowConfigurationOptions,omitempty"`

	// The maximum number of web servers that you want to run in your environment. Value need to be between `2` and `5`. Will be `2` by default.
	MaxWebservers int `json:"maxWebservers,omitempty" yaml:"maxWebservers,omitempty"`

	// The name of the Apache Airflow Environment
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion string `json:"pluginsS3ObjectVersion,omitempty" yaml:"pluginsS3ObjectVersion,omitempty"`

	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirements_s3_object_version is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path string `json:"requirementsS3Path,omitempty" yaml:"requirementsS3Path,omitempty"`

	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn string `json:"sourceBucketArn,omitempty" yaml:"sourceBucketArn,omitempty"`

	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode string `json:"webserverAccessMode,omitempty" yaml:"webserverAccessMode,omitempty"`
}
