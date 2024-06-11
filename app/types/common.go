package types

import "bytes"

type FileBuffer struct {
	Name    string
	Content *bytes.Buffer
}

type Directory struct {
	Name    string
	Files   []FileBuffer
	SubDirs map[string]*Directory
}
