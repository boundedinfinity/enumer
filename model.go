package enumer

type EnumData struct {
	Type        string        `json:"type" yaml:"type"`
	Struct      string        `json:"struct" yaml:"struct"`
	Package     string        `json:"package" yaml:"package"`
	InputPath   string        `json:"input-path" yaml:"input-path"`
	OutputPath  string        `json:"output-path" yaml:"output-path"`
	Desc        string        `json:"desc" yaml:"desc"`
	Header      string        `json:"header" yaml:"header"`
	HeaderFrom  string        `json:"header-from" yaml:"header-from"`
	HeaderLines []string      `json:"header-lines" yaml:"header-lines"`
	SkipFormat  bool          `json:"skip-format" yaml:"skip-format"`
	Debug       bool          `json:"debug" yaml:"debug"`
	Overwrite   bool          `json:"overwrite" yaml:"overwrite"`
	Serialize   EnumSerialize `json:"serialize" yaml:"serialize"`
	Values      []EnumValue   `json:"values" yaml:"values"`
}

type EnumSerialize struct {
	Type  string `json:"type" yaml:"type"`
	Value string `json:"value" yaml:"value"`
}

type EnumValue struct {
	Name       string `json:"name" yaml:"name"`
	Serialized string `json:"serialized" yaml:"serialized"`
}
