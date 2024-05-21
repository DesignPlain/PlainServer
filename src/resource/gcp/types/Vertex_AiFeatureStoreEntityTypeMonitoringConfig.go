package types

type Vertex_AiFeatureStoreEntityTypeMonitoringConfig struct {
	/*
	   Threshold for categorical features of anomaly detection. This is shared by all types of Featurestore Monitoring for categorical features (i.e. Features with type (Feature.ValueType) BOOL or STRING).
	   Structure is documented below.
	*/
	CategoricalThresholdConfig Vertex_AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfig `json:"categoricalThresholdConfig,omitempty" yaml:"categoricalThresholdConfig,omitempty"`

	/*
	   The config for ImportFeatures Analysis Based Feature Monitoring.
	   Structure is documented below.
	*/
	ImportFeaturesAnalysis Vertex_AiFeatureStoreEntityTypeMonitoringConfigImportFeaturesAnalysis `json:"importFeaturesAnalysis,omitempty" yaml:"importFeaturesAnalysis,omitempty"`

	/*
	   Threshold for numerical features of anomaly detection. This is shared by all objectives of Featurestore Monitoring for numerical features (i.e. Features with type (Feature.ValueType) DOUBLE or INT64).
	   Structure is documented below.
	*/
	NumericalThresholdConfig Vertex_AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfig `json:"numericalThresholdConfig,omitempty" yaml:"numericalThresholdConfig,omitempty"`

	/*
	   The config for Snapshot Analysis Based Feature Monitoring.
	   Structure is documented below.
	*/
	SnapshotAnalysis Vertex_AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysis `json:"snapshotAnalysis,omitempty" yaml:"snapshotAnalysis,omitempty"`
}
