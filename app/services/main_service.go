package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"markdown-to-pages-converter/app/types"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func MainService(ctx *fiber.Ctx, files []*multipart.FileHeader) ([]byte, error) {

	directory, err := WriteFileIntoBuffer(files[0])
	if err != nil {
		return []byte(nil), err
	}

	buffer := new(bytes.Buffer)

	zipWriter := zip.NewWriter(buffer)
	var output []string
	printDirectory(&output, directory, 0)

	if err := TemplateService(output, directory); err != nil {
		return []byte(nil), err
	}

	if err := AddDirectoryToZip(zipWriter, directory, ""); err != nil {
		return []byte(nil), err
	}

	if err := zipWriter.Close(); err != nil {
		return []byte(nil), err
	}

	return buffer.Bytes(), nil
}

func printDirectory(output *[]string, dir *types.Directory, indent int) {
	prefix := strings.Repeat(" ", indent)
	fmt.Printf("%s%s\n", prefix, dir.Name)
	clearedName := strings.Split(dir.Name, "/")
	selectedName := clearedName[1:]
	if len(selectedName) > 0 && selectedName[0] != "" {
		*output = append(*output, strings.Join(selectedName, "/"))
	}
	for _, file := range dir.Files {
		fmt.Printf("%s %s (size: %d bytes)\n", prefix, file.Name, file.Content.Len())
	}

	for _, subDir := range dir.SubDirs {
		printDirectory(output, subDir, indent+1)
	}
}
