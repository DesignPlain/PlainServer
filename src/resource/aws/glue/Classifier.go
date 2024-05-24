package glue

import types "DesignSphere_Server/src/resource/aws/types"

type Classifier struct {
	// A classifier for Csv content. Defined below.
	CsvClassifier types.Glue_ClassifierCsvClassifier `json:"csvClassifier,omitempty" yaml:"csvClassifier,omitempty"`

	// A classifier that uses grok patterns. Defined below.
	GrokClassifier types.Glue_ClassifierGrokClassifier `json:"grokClassifier,omitempty" yaml:"grokClassifier,omitempty"`

	// A classifier for JSON content. Defined below.
	JsonClassifier types.Glue_ClassifierJsonClassifier `json:"jsonClassifier,omitempty" yaml:"jsonClassifier,omitempty"`

	// The name of the classifier.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A classifier for XML content. Defined below.
	XmlClassifier types.Glue_ClassifierXmlClassifier `json:"xmlClassifier,omitempty" yaml:"xmlClassifier,omitempty"`
}
