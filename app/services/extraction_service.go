package services

import (
	"archive/zip"
	"bytes"
	"io"
	"markdown-to-pages-converter/app/types"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func WriteFileIntoBuffer(file *multipart.FileHeader) (*types.Directory, error) {
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

	rootDir := &types.Directory{
		Name:    "",
		Files:   []types.FileBuffer{},
		SubDirs: make(map[string]*types.Directory),
	}

	for _, zipFile := range zipReader.File {
		if err := processFile(zipFile, rootDir); err != nil {
			return nil, err
		}
	}

	return rootDir, nil

}

func processFile(zipFile *zip.File, currentDir *types.Directory) error {
	if zipFile.FileInfo().IsDir() {
		subDirName := strings.TrimSuffix(zipFile.Name, "/")
		currentDir.SubDirs[subDirName] = &types.Directory{
			Name:    subDirName,
			Files:   []types.FileBuffer{},
			SubDirs: make(map[string]*types.Directory),
		}
		return nil
	}

	rc, err := zipFile.Open()
	if err != nil {
		return err
	}

	fileBuffer, err := ConversionService(&rc)
	if err != nil {
		return err
	}

	defer rc.Close()

	pathParts := strings.Split(zipFile.Name, "/")

	dir := currentDir
	for _, part := range pathParts[:len(pathParts)-1] {
		if subDir, exists := dir.SubDirs[part]; exists {
			dir = subDir
		} else {
			newDir := &types.Directory{
				Name:    part,
				Files:   []types.FileBuffer{},
				SubDirs: make(map[string]*types.Directory),
			}

			dir.SubDirs[part] = newDir
			dir = newDir
		}
	}

	dir.Files = append(dir.Files, types.FileBuffer{
		Name:    "+page.svelte",
		Content: &fileBuffer,
	})

	return nil
}

func AddDirectoryToZip(zipWriter *zip.Writer, dir *types.Directory, basePath string) error {
	currentPath := filepath.Join(basePath, dir.Name)

	for _, file := range dir.Files {
		header := &zip.FileHeader{
			Name:   filepath.Join(currentPath, file.Name),
			Method: zip.Deflate,
		}

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if _, err := writer.Write(file.Content.Bytes()); err != nil {
			return err
		}
	}

	for _, subDir := range dir.SubDirs {
		if err := AddDirectoryToZip(zipWriter, subDir, currentPath); err != nil {
			return err
		}
	}

	return nil
}
