package services

import (
	"bytes"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func ConversionService(fileBuffer *bytes.Buffer, rc *io.ReadCloser) error {
	byteSlice := fileBuffer.Bytes()
	byteConversion := convertMdToHtml(byteSlice)
	byteBuffer := bytes.NewBuffer(byteConversion)
	*fileBuffer = *byteBuffer
	if _, err := io.Copy(fileBuffer, *rc); err != nil {
		return err
	}

	return nil
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
