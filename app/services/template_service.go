package services

import (
	"bytes"
	"fmt"
	"io"
	"markdown-to-pages-converter/app/types"
	"os"
	"path/filepath"
	"strings"
)

func TemplateService(routes []string, directory *types.Directory) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	// TODO: adjust by framework
	fp := filepath.Join(cwd, "/templates/svelte/+layout.svelte")
	file, err := os.OpenFile(fp, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	content, err := io.ReadAll(file)

	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	marking := "<!-- navigations -->"

	var contents []string
	var titles []string
	for _, route := range routes {
		if title := strings.Split(route, "/")[0]; !contains(titles, title) {
			titles = append(titles, title)
			contents = append(contents, fmt.Sprintf(`<p class="title">%s</p>`, title))
		}
		contents = append(contents, fmt.Sprintf(`<a href="/%s" class="content-name">%s</a>`, route, route))
	}

	allContent := strings.Join(contents, "")

	contentString := string(content)
	modifiedString := strings.Replace(contentString, marking, allContent, -1)

	directory.Files = append(directory.Files, types.FileBuffer{
		Name:    "+layout.svelte",
		Content: bytes.NewBufferString(modifiedString),
	})

	file.Close()
	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
