package dataflow

import types "DesignSphere_Server/src/resource/gcp/types"

type Pipeline struct {
	// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
	   "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
	   "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
	   "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A reference to the region
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
	SchedulerServiceAccountEmail string `json:"schedulerServiceAccountEmail,omitempty" yaml:"schedulerServiceAccountEmail,omitempty"`

	/*
	   The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
	   Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	/*
	   Workload information for creating new jobs.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
	   Structure is documented below.
	*/
	Workload types.Dataflow_PipelineWorkload `json:"workload,omitempty" yaml:"workload,omitempty"`

	/*
	   The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	*/
	PipelineSources map[string]string `json:"pipelineSources,omitempty" yaml:"pipelineSources,omitempty"`

	/*
	   Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
	   Structure is documented below.
	*/
	ScheduleInfo types.Dataflow_PipelineScheduleInfo `json:"scheduleInfo,omitempty" yaml:"scheduleInfo,omitempty"`

	/*
	   The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
	   https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
	   Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.


	   - - -
	*/
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
