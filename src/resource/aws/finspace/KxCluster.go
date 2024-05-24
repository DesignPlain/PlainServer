package finspace

import types "DesignSphere_Server/src/resource/aws/types"

type KxCluster struct {
	// List of key-value pairs to make available inside the cluster.
	CommandLineArguments map[string]string `json:"commandLineArguments,omitempty" yaml:"commandLineArguments,omitempty"`

	// Size and type of the temporary storage that is used to hold data during the savedown process. This parameter is required when you choose `type` as RDB. All the data written to this storage space is lost when the cluster node is restarted. See savedown_storage_configuration.
	SavedownStorageConfiguration types.Finspace_KxClusterSavedownStorageConfiguration `json:"savedownStorageConfiguration,omitempty" yaml:"savedownStorageConfiguration,omitempty"`

	/*
	   Configuration details about the network where the Privatelink endpoint of the cluster resides. See vpc_configuration.

	   The following arguments are optional:
	*/
	VpcConfiguration types.Finspace_KxClusterVpcConfiguration `json:"vpcConfiguration,omitempty" yaml:"vpcConfiguration,omitempty"`

	// Configuration based on which FinSpace will scale in or scale out nodes in your cluster. See auto_scaling_configuration.
	AutoScalingConfiguration types.Finspace_KxClusterAutoScalingConfiguration `json:"autoScalingConfiguration,omitempty" yaml:"autoScalingConfiguration,omitempty"`

	// KX database that will be available for querying. Defined below.
	Databases []types.Finspace_KxClusterDatabase `json:"databases,omitempty" yaml:"databases,omitempty"`

	// Version of FinSpace Managed kdb to run.
	ReleaseLabel string `json:"releaseLabel,omitempty" yaml:"releaseLabel,omitempty"`

	// The structure that stores the configuration details of a scaling group.
	ScalingGroupConfiguration types.Finspace_KxClusterScalingGroupConfiguration `json:"scalingGroupConfiguration,omitempty" yaml:"scalingGroupConfiguration,omitempty"`

	// The availability zone identifiers for the requested regions. Required when `az_mode` is set to SINGLE.
	AvailabilityZoneId string `json:"availabilityZoneId,omitempty" yaml:"availabilityZoneId,omitempty"`

	// Structure for the metadata of a cluster. Includes information like the CPUs needed, memory of instances, and number of instances. See capacity_configuration.
	CapacityConfiguration types.Finspace_KxClusterCapacityConfiguration `json:"capacityConfiguration,omitempty" yaml:"capacityConfiguration,omitempty"`

	// Details of the custom code that you want to use inside a cluster when analyzing data. Consists of the S3 source bucket, location, object version, and the relative path from where the custom code is loaded into the cluster. See code.
	Code types.Finspace_KxClusterCode `json:"code,omitempty" yaml:"code,omitempty"`

	// Unique identifier for the KX environment.
	EnvironmentId string `json:"environmentId,omitempty" yaml:"environmentId,omitempty"`

	// An IAM role that defines a set of permissions associated with a cluster. These permissions are assumed when a cluster attempts to access another cluster.
	ExecutionRole string `json:"executionRole,omitempty" yaml:"executionRole,omitempty"`

	// Path to Q program that will be run at launch of a cluster. This is a relative path within .zip file that contains the custom code, which will be loaded on the cluster. It must include the file name itself. For example, somedir/init.q.
	InitializationScript string `json:"initializationScript,omitempty" yaml:"initializationScript,omitempty"`

	// Unique name for the cluster that you want to create.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Key-value mapping of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The number of availability zones you want to assign per cluster. This can be one of the following:
	   - SINGLE - Assigns one availability zone per cluster.
	   - MULTI - Assigns all the availability zones per cluster.
	*/
	AzMode string `json:"azMode,omitempty" yaml:"azMode,omitempty"`

	// Configurations for a read only cache storage associated with a cluster. This cache will be stored as an FSx Lustre that reads from the S3 store. See cache_storage_configuration.
	CacheStorageConfigurations []types.Finspace_KxClusterCacheStorageConfiguration `json:"cacheStorageConfigurations,omitempty" yaml:"cacheStorageConfigurations,omitempty"`

	// Description of the cluster.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A configuration to store Tickerplant logs. It consists of a list of volumes that will be mounted to your cluster. For the cluster type Tickerplant , the location of the TP volume on the cluster will be available by using the global variable .aws.tp_log_path.
	TickerplantLogConfigurations []types.Finspace_KxClusterTickerplantLogConfiguration `json:"tickerplantLogConfigurations,omitempty" yaml:"tickerplantLogConfigurations,omitempty"`

	/*
	   Type of KDB database. The following types are available:
	   - HDB - Historical Database. The data is only accessible with read-only permissions from one of the FinSpace managed KX databases mounted to the cluster.
	   - RDB - Realtime Database. This type of database captures all the data from a ticker plant and stores it in memory until the end of day, after which it writes all of its data to a disk and reloads the HDB. This cluster type requires local storage for temporary storage of data during the savedown process. If you specify this field in your request, you must provide the `savedownStorageConfiguration` parameter.
	   - GATEWAY - A gateway cluster allows you to access data across processes in kdb systems. It allows you to create your own routing logic using the initialization scripts and custom code. This type of cluster does not require a  writable local storage.
	   - GP - A general purpose cluster allows you to quickly iterate on code during development by granting greater access to system commands and enabling a fast reload of custom code. This cluster type can optionally mount databases including cache and savedown storage. For this cluster type, the node count is fixed at 1. It does not support autoscaling and supports only `SINGLE` AZ mode.
	   - Tickerplant – A tickerplant cluster allows you to subscribe to feed handlers based on IAM permissions. It can publish to RDBs, other Tickerplants, and real-time subscribers (RTS). Tickerplants can persist messages to log, which is readable by any RDB environment. It supports only single-node that is only one kdb process.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
