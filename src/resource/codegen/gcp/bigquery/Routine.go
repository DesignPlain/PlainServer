package bigquery

import types "libds/gcp/types"

type Routine struct {
	/*
	   The body of the routine. For functions, this is the expression in the AS clause.
	   If language=SQL, it is the substring inside (but excluding) the parentheses.


	   - - -
	*/
	DefinitionBody string `json:"definitionBody,omitempty" yaml:"definitionBody,omitempty"`

	// The description of the routine if defined.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Optional. If language = "JAVASCRIPT", this field stores the path of the
	   imported JAVASCRIPT libraries.
	*/
	ImportedLibraries []string `json:"importedLibraries,omitempty" yaml:"importedLibraries,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".
	   If absent, the return table type is inferred from definitionBody at query time in each query
	   that references this routine. If present, then the columns in the evaluated table result will
	   be cast to match the column types specificed in return table type, at query time.
	*/
	ReturnTableType string `json:"returnTableType,omitempty" yaml:"returnTableType,omitempty"`

	/*
	   A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
	   If absent, the return type is inferred from definitionBody at query time in each query
	   that references this routine. If present, then the evaluated result will be cast to
	   the specified returned type at query time. ~>--NOTE--: Because this field expects a JSON
	   string, any changes to the string will create a diff, even if the JSON itself hasn't
	   changed. If the API returns a different value for the same schema, e.g. it switche
	   d the order of values or replaced STRUCT field type with RECORD field type, we currently
	   cannot suppress the recurring diff this causes. As a workaround, we recommend using
	   the schema as returned by the API.
	*/
	ReturnType string `json:"returnType,omitempty" yaml:"returnType,omitempty"`

	/*
	   The determinism level of the JavaScript UDF if defined.
	   Possible values are: `DETERMINISM_LEVEL_UNSPECIFIED`, `DETERMINISTIC`, `NOT_DETERMINISTIC`.
	*/
	DeterminismLevel string `json:"determinismLevel,omitempty" yaml:"determinismLevel,omitempty"`

	/*
	   The type of routine.
	   Possible values are: `SCALAR_FUNCTION`, `PROCEDURE`, `TABLE_VALUED_FUNCTION`.
	*/
	RoutineType string `json:"routineType,omitempty" yaml:"routineType,omitempty"`

	// The ID of the dataset containing this routine
	DatasetId string `json:"datasetId,omitempty" yaml:"datasetId,omitempty"`

	/*
	   The language of the routine.
	   Possible values are: `SQL`, `JAVASCRIPT`, `PYTHON`, `JAVA`, `SCALA`.
	*/
	Language string `json:"language,omitempty" yaml:"language,omitempty"`

	/*
	   Input/output argument of a function or a stored procedure.
	   Structure is documented below.
	*/
	Arguments []types.Bigquery_RoutineArgument `json:"arguments,omitempty" yaml:"arguments,omitempty"`

	/*
	   Remote function specific options.
	   Structure is documented below.
	*/
	RemoteFunctionOptions types.Bigquery_RoutineRemoteFunctionOptions `json:"remoteFunctionOptions,omitempty" yaml:"remoteFunctionOptions,omitempty"`

	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.
	RoutineId string `json:"routineId,omitempty" yaml:"routineId,omitempty"`

	/*
	   Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	   Structure is documented below.
	*/
	SparkOptions types.Bigquery_RoutineSparkOptions `json:"sparkOptions,omitempty" yaml:"sparkOptions,omitempty"`
}
