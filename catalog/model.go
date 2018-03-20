package catalog

import (
	"gosiah/solr"
	// "log"
)

type Catalog struct {
	coreUrl string
}

type BibRecord struct {
	Bib      string
	Title    string
	Version  float64
	Subjects []string
}

func New(coreUrl string) Catalog {
	return Catalog{coreUrl: coreUrl}
}

func (c Catalog) Get(id string) (BibRecord, error) {
	s := solr.New(c.coreUrl)
	doc, err := s.Get(id, "")
	if err != nil {
		return BibRecord{}, err
	}
	return DocToRecord(doc), nil
}

func DocToRecord(doc solr.Document) BibRecord {
	id := doc.Value("id")
	title := doc.FirstValue("title_str")
	version := doc.ValueFloat("_version_")
	subjects := doc.Values("subjects_str")
	return BibRecord{Bib: id, Title: title, Version: version, Subjects: subjects}
}

func (c Catalog) Search(q string) ([]BibRecord, error) {
	s := solr.New(c.coreUrl)
	r, err := s.Search("*", "")
	if err != nil {
		return []BibRecord{}, err
	}

	// log.Printf("num found:%d", r.Response.NumFound)
	// log.Printf("params:")
	// for key, value := range r.ResponseHeader.Params {
	// 	log.Printf("\t%s = %s", key, value)
	// }

	records := []BibRecord{}
	for _, doc := range r.Response.Documents {
		record := DocToRecord(doc)
		records = append(records, record)
	}
	return records, nil
}
