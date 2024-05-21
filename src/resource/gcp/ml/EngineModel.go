package ml

import types "DesignSphere_Server/src/resource/gcp/types"

type EngineModel struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The list of regions where the model is going to be deployed.
	   Currently only one region per model is supported
	*/
	Regions string `json:"regions,omitempty" yaml:"regions,omitempty"`

	/*
	   The default version of the model. This version will be used to handle
	   prediction requests that do not specify a version.
	   Structure is documented below.
	*/
	DefaultVersion types.Ml_EngineModelDefaultVersion `json:"defaultVersion,omitempty" yaml:"defaultVersion,omitempty"`

	// The description specified for the model when it was created.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   One or more labels that you can add, to organize your models.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The name specified for the model.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// If true, online prediction nodes send stderr and stdout streams to Stackdriver Logging
	OnlinePredictionConsoleLogging bool `json:"onlinePredictionConsoleLogging,omitempty" yaml:"onlinePredictionConsoleLogging,omitempty"`

	// If true, online prediction access logs are sent to StackDriver Logging.
	OnlinePredictionLogging bool `json:"onlinePredictionLogging,omitempty" yaml:"onlinePredictionLogging,omitempty"`
}
