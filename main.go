package main

import (
	"fmt"
	"os"
	"solrdora/web"
)

func main() {
	if len(os.Args) < 2 {
		displayHelp()
		return
	}
	settingsFile := os.Args[1]
	web.StartWebServer(settingsFile)
}

func displayHelp() {
	sample := `
Must indicate a settings.json file with an structure like this:

	{
	  "serverAddress": "localhost:9001",
	  "solrCoreUrl": "http://localhost:8983/solr/bibdata",
	  "solrOptions" : {
	    "defType": "edismax",
	    "qf": "authorsAll title^100"
	  },
	  "solrFacets": {
	    "subjects_str": "Subjects",
	    "publisher_str": "Publisher"
	  },
		"searchFl": ["id", "title", "subjects", "author"],
	  "viewOneFl": ["id", "title", "authorsAll", "_version_"]
	}`
	fmt.Printf("%s\r\n\r\n", sample)
}
