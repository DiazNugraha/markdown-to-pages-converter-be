package services

import (
	"bytes"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func ConversionService(rc *io.ReadCloser) (bytes.Buffer, error) {
	var fileBuffer bytes.Buffer
	if _, err := io.Copy(&fileBuffer, *rc); err != nil {
		return fileBuffer, err
	}

	byteSlice := fileBuffer.Bytes()
	byteConversion := convertMdToHtml(byteSlice)
	modifiedString := modifyString(string(byteConversion))
	byteBuffer := bytes.NewBuffer([]byte(modifiedString))

	return *byteBuffer, nil
}

func modifyString(text string) string {
	newString := strings.Replace(text, "{", "&#123;", -1)
	newString = strings.Replace(newString, "}", "&#125;", -1)
	return newString
}

func convertMdToHtml(md []byte) []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock

	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank

	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
