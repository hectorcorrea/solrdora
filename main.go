package main

import (
	"fmt"
	"gosiah/web"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		displayHelp()
	}
	settingsFile := os.Args[1]
	web.StartWebServer(settingsFile)
}

func displayHelp() {
	sample := `{
	  "serverUrl": "localhost:9001",
	  "solrCoreUrl": "http://localhost:8983/solr/bibdata",
	  "solrOptions" : {
	    "defType": "edismax",
	    "qf": "authorsAll title^100"
	  },
	  "solrFacets": {
	    "subjects_str": "Subjects",
	    "publisher_str": "Publisher"
	  }
	}`
	msg := fmt.Sprintf("Must indicate a settings.json file with an structure like this:\r\n%s", sample)
	log.Fatal(msg)
}
