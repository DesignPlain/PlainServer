package types

type Glue_MLTransformParameters struct {
	// The parameters for the find matches algorithm. see Find Matches Parameters.
	FindMatchesParameters Glue_MLTransformParametersFindMatchesParameters `json:"findMatchesParameters,omitempty" yaml:"findMatchesParameters,omitempty"`

	// The type of machine learning transform. For information about the types of machine learning transforms, see [Creating Machine Learning Transforms](http://docs.aws.amazon.com/glue/latest/dg/add-job-machine-learning-transform.html).
	TransformType string `json:"transformType,omitempty" yaml:"transformType,omitempty"`
}
