package dataproc

import types "libds/gcp/types"

type Job struct {
	// The config of Hadoop job
	HadoopConfig types.Dataproc_JobHadoopConfig `json:"hadoopConfig,omitempty" yaml:"hadoopConfig,omitempty"`

	// The config of hive job
	HiveConfig types.Dataproc_JobHiveConfig `json:"hiveConfig,omitempty" yaml:"hiveConfig,omitempty"`

	/*
	   The list of labels (key/value pairs) to add to the job.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field 'effective_labels' for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The Cloud Dataproc region. This essentially determines which clusters are available
	   for this job to be submitted to. If not specified, defaults to `global`.
	*/
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Optional. Job scheduling configuration.
	Scheduling types.Dataproc_JobScheduling `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`

	/*
	   By default, you can only delete inactive jobs within
	   Dataproc. Setting this to true, and calling destroy, will ensure that the
	   job is first cancelled before issuing the delete.
	*/
	ForceDelete bool `json:"forceDelete,omitempty" yaml:"forceDelete,omitempty"`

	// The config of job placement.
	Placement types.Dataproc_JobPlacement `json:"placement,omitempty" yaml:"placement,omitempty"`

	// The config of the Spark job.
	SparkConfig types.Dataproc_JobSparkConfig `json:"sparkConfig,omitempty" yaml:"sparkConfig,omitempty"`

	// The config of pag job.
	PigConfig types.Dataproc_JobPigConfig `json:"pigConfig,omitempty" yaml:"pigConfig,omitempty"`

	// The config of presto job
	PrestoConfig types.Dataproc_JobPrestoConfig `json:"prestoConfig,omitempty" yaml:"prestoConfig,omitempty"`

	// The config of pySpark job.
	PysparkConfig types.Dataproc_JobPysparkConfig `json:"pysparkConfig,omitempty" yaml:"pysparkConfig,omitempty"`

	// The reference of the job
	Reference types.Dataproc_JobReference `json:"reference,omitempty" yaml:"reference,omitempty"`

	/*
	   The project in which the `cluster` can be found and jobs
	   subsequently run against. If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The config of SparkSql job
	SparksqlConfig types.Dataproc_JobSparksqlConfig `json:"sparksqlConfig,omitempty" yaml:"sparksqlConfig,omitempty"`
}
