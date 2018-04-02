package catalog

import (
	"github.com/hectorcorrea/solr"
)

type BibRecord struct {
	Bib     string
	Title   string
	Version float64
	Authors []string
}

func NewBibRecord(doc solr.Document) BibRecord {
	id := doc.Value("id")
	title := doc.Value("title")
	version := doc.ValueFloat("_version_")
	authors := doc.Values("authorsAll")
	return BibRecord{Bib: id, Title: title, Version: version, Authors: authors}
}
