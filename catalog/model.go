package catalog

import (
	"fmt"
	"gosiah/solr"
)

type Catalog struct {
	coreUrl string
}

type BibRecord struct {
	Bib   string
	Title string
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
	id := fieldValue(doc, "id")
	title := fieldValues(doc, "title_str")
	return BibRecord{Bib: id, Title: title}
}

func fieldValue(doc solr.Document, field string) string {
	value, ok := doc[field].(string)
	if ok {
		return fmt.Sprintf("%s", value)
	}
	return ""
}

func fieldValues(doc solr.Document, field string) string {
	values, ok := doc[field].([]interface{})
	if ok && len(values) > 0 {
		// get the first now for now
		return fmt.Sprintf("%s", values[0])
	}
	return ""
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
		// log.Printf("%v", doc)
		record := DocToRecord(doc)
		records = append(records, record)
	}
	return records, nil
}
