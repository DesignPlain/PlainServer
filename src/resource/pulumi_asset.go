package asset

import "encoding/json"

type FileAsset struct {
	FileAssetName string `yaml:"fn::FileAsset,omitempty"`
}

/* Custom Marshaller for FileAsset */
func (s FileAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.FileAssetName)
}

/* Custom Unmarshaller for FileAsset */
func (s *FileAsset) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	s.FileAssetName = str
	return nil
}
