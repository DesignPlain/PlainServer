package types

type Appflow_FlowTask struct {
	// Source fields to which a particular task is applied.
	SourceFields []string `json:"sourceFields,omitempty" yaml:"sourceFields,omitempty"`

	// Map used to store task-related information. The execution service looks for particular information based on the `TaskType`. Valid keys are `VALUE`, `VALUES`, `DATA_TYPE`, `UPPER_BOUND`, `LOWER_BOUND`, `SOURCE_DATA_TYPE`, `DESTINATION_DATA_TYPE`, `VALIDATION_ACTION`, `MASK_VALUE`, `MASK_LENGTH`, `TRUNCATE_LENGTH`, `MATH_OPERATION_FIELDS_ORDER`, `CONCAT_FORMAT`, `SUBFIELD_CATEGORY_MAP`, and `EXCLUDE_SOURCE_FIELDS_LIST`.
	TaskProperties map[string]string `json:"taskProperties,omitempty" yaml:"taskProperties,omitempty"`

	// Particular task implementation that Amazon AppFlow performs. Valid values are `Arithmetic`, `Filter`, `Map`, `Map_all`, `Mask`, `Merge`, `Passthrough`, `Truncate`, and `Validate`.
	TaskType string `json:"taskType,omitempty" yaml:"taskType,omitempty"`

	// Operation to be performed on the provided source fields. See Connector Operator for details.
	ConnectorOperators []Appflow_FlowTaskConnectorOperator `json:"connectorOperators,omitempty" yaml:"connectorOperators,omitempty"`

	// Field in a destination connector, or a field value against which Amazon AppFlow validates a source field.
	DestinationField string `json:"destinationField,omitempty" yaml:"destinationField,omitempty"`
}
