package excel

type Header struct {
	headers []string
	excel   *Excel
}

type Excel struct {
	filePath      string
	configPath    string
	primaryKey    string
	isHeaderExist bool
	headersMap    map[string]int
	header        *Header
}