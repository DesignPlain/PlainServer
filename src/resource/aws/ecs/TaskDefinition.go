package ecs

import types "DesignSphere_Server/src/resource/aws/types"

type TaskDefinition struct {
	// A list of valid [container definitions](http://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_ContainerDefinition.html) provided as a single valid JSON document. Please note that you should only provide values that are part of the container definition document. For a detailed description of what parameters are available, see the [Task Definition Parameters](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task_definition_parameters.html) section from the official [Developer Guide](https://docs.aws.amazon.com/AmazonECS/latest/developerguide).
	ContainerDefinitions string `json:"containerDefinitions,omitempty" yaml:"containerDefinitions,omitempty"`

	// Configuration block for volumes that containers in your task may use. Detailed below.
	Volumes []types.Ecs_TaskDefinitionVolume `json:"volumes,omitempty" yaml:"volumes,omitempty"`

	// Set of launch types required by the task. The valid values are `EC2` and `FARGATE`.
	RequiresCompatibilities []string `json:"requiresCompatibilities,omitempty" yaml:"requiresCompatibilities,omitempty"`

	// Whether to retain the old revision when the resource is destroyed or replacement is necessary. Default is `false`.
	SkipDestroy bool `json:"skipDestroy,omitempty" yaml:"skipDestroy,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// ARN of the task execution role that the Amazon ECS container agent and the Docker daemon can assume.
	ExecutionRoleArn string `json:"executionRoleArn,omitempty" yaml:"executionRoleArn,omitempty"`

	// Process namespace to use for the containers in the task. The valid values are `host` and `task`.
	PidMode string `json:"pidMode,omitempty" yaml:"pidMode,omitempty"`

	// Configuration block for rules that are taken into consideration during task placement. Maximum number of `placement_constraints` is `10`. Detailed below.
	PlacementConstraints []types.Ecs_TaskDefinitionPlacementConstraint `json:"placementConstraints,omitempty" yaml:"placementConstraints,omitempty"`

	// Configuration block for the App Mesh proxy. Detailed below.
	ProxyConfiguration types.Ecs_TaskDefinitionProxyConfiguration `json:"proxyConfiguration,omitempty" yaml:"proxyConfiguration,omitempty"`

	// Configuration block for runtime_platform that containers in your task may use.
	RuntimePlatform types.Ecs_TaskDefinitionRuntimePlatform `json:"runtimePlatform,omitempty" yaml:"runtimePlatform,omitempty"`

	// IPC resource namespace to be used for the containers in the task The valid values are `host`, `task`, and `none`.
	IpcMode string `json:"ipcMode,omitempty" yaml:"ipcMode,omitempty"`

	// Docker networking mode to use for the containers in the task. Valid values are `none`, `bridge`, `awsvpc`, and `host`.
	NetworkMode string `json:"networkMode,omitempty" yaml:"networkMode,omitempty"`

	/*
	   A unique name for your task definition.

	   The following arguments are optional:
	*/
	Family string `json:"family,omitempty" yaml:"family,omitempty"`

	// Configuration block(s) with Inference Accelerators settings. Detailed below.
	InferenceAccelerators []types.Ecs_TaskDefinitionInferenceAccelerator `json:"inferenceAccelerators,omitempty" yaml:"inferenceAccelerators,omitempty"`

	// Amount (in MiB) of memory used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
	Memory string `json:"memory,omitempty" yaml:"memory,omitempty"`

	// ARN of IAM role that allows your Amazon ECS container task to make calls to other AWS services.
	TaskRoleArn string `json:"taskRoleArn,omitempty" yaml:"taskRoleArn,omitempty"`

	// Whether should track latest task definition or the one created with the resource. Default is `false`.
	TrackLatest bool `json:"trackLatest,omitempty" yaml:"trackLatest,omitempty"`

	// Number of cpu units used by the task. If the `requires_compatibilities` is `FARGATE` this field is required.
	Cpu string `json:"cpu,omitempty" yaml:"cpu,omitempty"`

	// The amount of ephemeral storage to allocate for the task. This parameter is used to expand the total amount of ephemeral storage available, beyond the default amount, for tasks hosted on AWS Fargate. See Ephemeral Storage.
	EphemeralStorage types.Ecs_TaskDefinitionEphemeralStorage `json:"ephemeralStorage,omitempty" yaml:"ephemeralStorage,omitempty"`
}
