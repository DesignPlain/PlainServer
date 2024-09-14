package types

type Glue_ClassifierXmlClassifier struct {
	// An identifier of the data format that the classifier matches.
	Classification string `json:"classification,omitempty" yaml:"classification,omitempty"`

	// The XML tag designating the element that contains each record in an XML document being parsed. Note that this cannot identify a self-closing element (closed by `/>`). An empty row element that contains only attributes can be parsed as long as it ends with a closing tag (for example, `<row item_a="A" item_b="B"></row>` is okay, but `<row item_a="A" item_b="B" />` is not).
	RowTag string `json:"rowTag,omitempty" yaml:"rowTag,omitempty"`
}
