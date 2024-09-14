package types

type Dataloss_PreventionStoredInfoTypeLargeCustomDictionary struct {
	/*
	   Field in a BigQuery table where each cell represents a dictionary phrase.
	   Structure is documented below.
	*/
	BigQueryField Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryBigQueryField `json:"bigQueryField,omitempty" yaml:"bigQueryField,omitempty"`

	/*
	   Set of files containing newline-delimited lists of dictionary phrases.
	   Structure is documented below.
	*/
	CloudStorageFileSet Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryCloudStorageFileSet `json:"cloudStorageFileSet,omitempty" yaml:"cloudStorageFileSet,omitempty"`

	/*
	   Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API.
	   If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.
	   Structure is documented below.
	*/
	OutputPath Dataloss_PreventionStoredInfoTypeLargeCustomDictionaryOutputPath `json:"outputPath,omitempty" yaml:"outputPath,omitempty"`
}
