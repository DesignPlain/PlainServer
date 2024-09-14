package types

type Autoscalingplans_ScalingPlanScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification struct {
	// Metric type. Valid values: `ALBRequestCountPerTarget`, `ASGAverageCPUUtilization`, `ASGAverageNetworkIn`, `ASGAverageNetworkOut`, `DynamoDBReadCapacityUtilization`, `DynamoDBWriteCapacityUtilization`, `ECSServiceAverageCPUUtilization`, `ECSServiceAverageMemoryUtilization`, `EC2SpotFleetRequestAverageCPUUtilization`, `EC2SpotFleetRequestAverageNetworkIn`, `EC2SpotFleetRequestAverageNetworkOut`, `RDSReaderAverageCPUUtilization`, `RDSReaderAverageDatabaseConnections`.
	PredefinedScalingMetricType string `json:"predefinedScalingMetricType,omitempty" yaml:"predefinedScalingMetricType,omitempty"`

	// Identifies the resource associated with the metric type.
	ResourceLabel string `json:"resourceLabel,omitempty" yaml:"resourceLabel,omitempty"`
}
