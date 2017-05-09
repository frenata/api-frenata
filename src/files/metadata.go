package files

import (
	"encoding/json"
	"log"
	"mime/multipart"
	"os"
)

type FileMetadataResponse struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// pretty print the values in JSON format
func (fr FileMetadataResponse) String() string {
	json, _ := json.MarshalIndent(fr, "", "    ")
	return string(json)
}

func getFileMetadata(header *multipart.FileHeader) string {
	filename := header.Filename
	file, err := header.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var size int64
	switch t := file.(type) {
	case *os.File:
		stats, err := t.Stat()
		if err != nil {
			log.Println(err)
		}
		size = stats.Size()
	default:
		bytes, err := file.Seek(0, 2)
		if err != nil {
			log.Println(err)
		}
		size = bytes
	}

	return FileMetadataResponse{filename, size}.String()
}
