package datapipeline

import types "DesignSphere_Server/src/resource/aws/types"

type PipelineDefinition struct {
	// Configuration block for the parameter objects used in the pipeline definition. See below
	ParameterObjects []types.Datapipeline_PipelineDefinitionParameterObject `json:"parameterObjects,omitempty" yaml:"parameterObjects,omitempty"`

	// Configuration block for the parameter values used in the pipeline definition. See below
	ParameterValues []types.Datapipeline_PipelineDefinitionParameterValue `json:"parameterValues,omitempty" yaml:"parameterValues,omitempty"`

	// ID of the pipeline.
	PipelineId string `json:"pipelineId,omitempty" yaml:"pipelineId,omitempty"`

	/*
	   Configuration block for the objects that define the pipeline. See below

	   The following arguments are optional:
	*/
	PipelineObjects []types.Datapipeline_PipelineDefinitionPipelineObject `json:"pipelineObjects,omitempty" yaml:"pipelineObjects,omitempty"`
}
