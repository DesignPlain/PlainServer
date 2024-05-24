package types

type Glue_ClassifierJsonClassifier struct {
	// A `JsonPath` string defining the JSON data for the classifier to classify. AWS Glue supports a subset of `JsonPath`, as described in [Writing JsonPath Custom Classifiers](https://docs.aws.amazon.com/glue/latest/dg/custom-classifier.html#custom-classifier-json).
	JsonPath string `json:"jsonPath,omitempty" yaml:"jsonPath,omitempty"`
}
