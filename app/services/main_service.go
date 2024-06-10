package services

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type FileBuffer struct {
	Name    string
	Content *bytes.Buffer
}

type Directory struct {
	Name    string
	Files   []FileBuffer
	SubDirs map[string]*Directory
}

func MainService(ctx *fiber.Ctx, files []*multipart.FileHeader) error {
	for _, file := range files {
		directory, err := saveFile(ctx, file)
		if err != nil {
			return err
		}

		printDirectory(directory, 0)
	}

	return nil
}

func saveFile(ctx *fiber.Ctx, file *multipart.FileHeader) (*Directory, error) {
	fileContent, err := file.Open()
	if err != nil {
		return nil, err
	}

	defer fileContent.Close()

	buffer := new(bytes.Buffer)
	if _, err := io.Copy(buffer, fileContent); err != nil {
		return nil, err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(buffer.Bytes()), int64(buffer.Len()))
	if err != nil {
		return nil, err
	}

	rootDir := &Directory{
		Name:    "",
		Files:   []FileBuffer{},
		SubDirs: make(map[string]*Directory),
	}

	for _, zipFile := range zipReader.File {
		if err := processFile(zipFile, rootDir); err != nil {
			return nil, err
		}
	}

	return rootDir, nil
}

func processFile(zipFile *zip.File, currentDir *Directory) error {
	if zipFile.FileInfo().IsDir() {
		subDirName := strings.TrimSuffix(zipFile.Name, "/")
		currentDir.SubDirs[subDirName] = &Directory{
			Name:    subDirName,
			Files:   []FileBuffer{},
			SubDirs: make(map[string]*Directory),
		}
		return nil
	}

	rc, err := zipFile.Open()
	if err != nil {
		return err
	}

	defer rc.Close()

	var fileBuffer bytes.Buffer

	if _, err := io.Copy(&fileBuffer, rc); err != nil {
		return err
	}

	pathParts := strings.Split(zipFile.Name, "/")
	dir := currentDir
	for _, part := range pathParts[:len(pathParts)-1] {
		if subDir, exists := dir.SubDirs[part]; exists {
			dir = subDir
		} else {
			newDir := &Directory{
				Name:    part,
				Files:   []FileBuffer{},
				SubDirs: make(map[string]*Directory),
			}

			dir.SubDirs[part] = newDir
			dir = newDir
		}
	}

	dir.Files = append(dir.Files, FileBuffer{
		Name:    pathParts[len(pathParts)-1],
		Content: &fileBuffer,
	})

	return nil
}

func printDirectory(dir *Directory, indent int) {
	prefix := strings.Repeat(" ", indent)
	fmt.Printf("%s%s\n", prefix, dir.Name)
	for _, file := range dir.Files {
		fmt.Printf("%s %s (size: %d bytes)\n", prefix, file.Name, file.Content.Len())
	}

	for _, subDir := range dir.SubDirs {
		printDirectory(subDir, indent+1)
	}
}
