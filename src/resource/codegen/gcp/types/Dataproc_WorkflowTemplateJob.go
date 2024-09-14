package types

type Dataproc_WorkflowTemplateJob struct {
	// Job is a PySpark job.
	PysparkJob Dataproc_WorkflowTemplateJobPysparkJob `json:"pysparkJob,omitempty" yaml:"pysparkJob,omitempty"`

	// Job scheduling configuration.
	Scheduling Dataproc_WorkflowTemplateJobScheduling `json:"scheduling,omitempty" yaml:"scheduling,omitempty"`

	// Job is a Spark job.
	SparkJob Dataproc_WorkflowTemplateJobSparkJob `json:"sparkJob,omitempty" yaml:"sparkJob,omitempty"`

	// Job is a SparkSql job.
	SparkSqlJob Dataproc_WorkflowTemplateJobSparkSqlJob `json:"sparkSqlJob,omitempty" yaml:"sparkSqlJob,omitempty"`

	// Required. The step id. The id must be unique among all jobs within the template. The step id is used as prefix for job id, as job `goog-dataproc-workflow-step-id` label, and in field from other steps. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters.
	StepId string `json:"stepId,omitempty" yaml:"stepId,omitempty"`

	// Job is a Hadoop job.
	HadoopJob Dataproc_WorkflowTemplateJobHadoopJob `json:"hadoopJob,omitempty" yaml:"hadoopJob,omitempty"`

	// Job is a Hive job.
	HiveJob Dataproc_WorkflowTemplateJobHiveJob `json:"hiveJob,omitempty" yaml:"hiveJob,omitempty"`

	// The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: {0,63} No more than 32 labels can be associated with a given job.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Job is a Pig job.
	PigJob Dataproc_WorkflowTemplateJobPigJob `json:"pigJob,omitempty" yaml:"pigJob,omitempty"`

	// The optional list of prerequisite job step_ids. If not specified, the job will start at the beginning of workflow.
	PrerequisiteStepIds []string `json:"prerequisiteStepIds,omitempty" yaml:"prerequisiteStepIds,omitempty"`

	// Job is a Presto job.
	PrestoJob Dataproc_WorkflowTemplateJobPrestoJob `json:"prestoJob,omitempty" yaml:"prestoJob,omitempty"`

	// Job is a SparkR job.
	SparkRJob Dataproc_WorkflowTemplateJobSparkRJob `json:"sparkRJob,omitempty" yaml:"sparkRJob,omitempty"`
}
