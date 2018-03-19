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

	title := ""
	titles, ok := doc["title_str"].([]interface{})
	if ok && len(titles) > 0 {
		title = fmt.Sprintf("%s", titles[0])
	}
	record := BibRecord{Bib: id, Title: title}
	return record, nil
}

// func x() {
// 	solr := solr.New("http://localhost:8983/solr/bibdata")
//
// 	r, err := solr.Search("*", "id,author,publisher")
// 	if err != nil {
// 		log.Fatal("ERROR: ", err)
// 	}
//
// 	log.Printf("num found:%d", r.Response.NumFound)
// 	log.Printf("params:")
// 	for key, value := range r.ResponseHeader.Params {
// 		log.Printf("\t%s = %s", key, value)
// 	}
//
// 	for i, doc := range r.Response.Documents {
// 		log.Printf("%d", i)
// 		for key, values := range doc {
// 			log.Printf("\t %s = %v", key, values)
// 		}
// 	}
//
// 	log.Printf("Get 00000018")
// 	d, err := solr.Get("00000018", "")
// 	log.Printf("%v", d)
// }
