package main

import (
	"fmt"
	"os"
)

const MermaidTemplate = "graph TD\n"

type Mermaid struct {
	content string
}

func (m *Mermaid) Create(content string) {
	m.content = content
}

func (m *Mermaid) RenderAndDisplay() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileName := pwd + "/mermaid_tree.md"
	// get tmp file path with random name
	tmpFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer tmpFile.Close()
	// write content to tmp file
	tmpFile.WriteString("```mermaid\n" + MermaidTemplate + m.content + "\n```")

	return fileName
}
