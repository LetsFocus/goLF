package excel

import "os"

type Header struct {
	headers []string
	excel   *Excel
}

type Excel struct {
	filePath 	  string
	configPath 	  string
	writeFile     *os.File
	readFile      *os.File
	writeConfig   *os.File
	readConfig    *os.File
	primaryKey    string
	doesHeaderExist bool
	headersMap    map[string]int
	header        *Header
}
