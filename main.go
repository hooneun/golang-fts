package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type documents struct {
	Documents []document `xml:"doc"`
}

type document struct {
	Title string `xml:"title"`
	URL   string `xml:"url"`
	Text  string `xml:"abstract"`
	ID    int
}

func loadDocuments(path string) ([]document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	dec := xml.NewDecoder(f)
	var dump documents
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}

	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}

func search(docs documents, term string) documents {
	var r documents
	for _, doc := range docs.Documents {
		if strings.Contains(doc.Text, term) {
			r.Documents = append(r.Documents, doc)
		}
	}
	return r
}

func main() {
	doc, err := loadDocuments("./enwiki-latest-abstract1.xml")
	if err != nil {
		panic(err)
	}

	fmt.Println(doc[0])

}
