# SolrDora
This is a small Go program to explore the data in a Solr core.

The core of the Solr functionality is provided by
[this other Go package](https://github.com/hectorcorrea/solr). What SolrDora
does is put a web user interface on top of it.

See a [screenshot of what it looks like](https://github.com/hectorcorrea/solrdora/blob/master/misc/screenshot.png) once it's running.


## Source Code
The main components of the code are:

* `main.go`: The launcher.
* `web/server.go`: The web server.
* `web/controllers.go`: The controllers are very skinny, they are a pass-through from the HTTP request to the model.
* `models/search.go`: This is the core search functionality, it calls the Go library that submits the request to Solr and uses `models/searchResult.go` to format the results in a way that can be used in the views.
* `views/`: The HTML views.


## Compiling and running
To get started

```
git clone https://github.com/hectorcorrea/solrdora.git
cd solrdora
go get
go build
./solrdora settings.json
```

## Settings.json
You'll need to provide a file with the settings for SolrDora to know what
Solr to connect to and some basics about it. The file `settings.json` in this
repo is a good place to start.

```
{
  "serverAddress": "localhost:9001",
  "solrCoreUrl": "http://localhost:8983/solr/bibdata",
  "solrOptions" : {
    "defType": "edismax",
    "qf": "author_txt_en title_txt_en^100",
    "wt": "json",
    "facet.limit": "20",
    "facet.mincount": "1",
    "hl": "on"
  },
  "solrFacets": {
    "subjects_ss": "1|Subject",
    "publisher_name_s": "2|Publisher"
  },
  "searchFl": ["id", "title_txt_en", "author_txt_en"],
  "viewOneFl": ["id", "title_txt_en", "author_txt_en", "authors_other_txts_en", "subjects_ss", "publisher_name_s", "publisher_place_s"]
}
```

The format for `solrFacets` is `"field_name" : "display title"`, notice that you can force the order of the facets by prefixing the display value with a number (e.g. "1|Subjects").

* `serverAddress` is the address at which SolrDora will run on your local box. Do not prepend "http" to this setting.
* `solrCoreUrl` is the URL of the Solr core that you want to connect to.
* `solrOptions` is a set of options that will be passed to Solr on every request. In the previous example you can see how we set the default parser (defType) to "edismax" and defining the query field (qf) to use the author and title fields. Any valid Solr parameter can be configured here.
* `solrFacets` is the list of fields that you want to use for facets. In the previous example we defined two fields "subjects_str" and "publisher_str". The syntax is `"field_name" : "Display Value"`
* `searchFl` defines the names of the fields that will be fetched from Solr (fl) during the search.
* `viewOneFl` defines the names of the fields that will be fetched from Solr (fl) when fetching the details of a document.

This program expects Solr's responses to be in JSON. If you Solr core does not return JSON by default (i.e. if you are using an old version of Solr), make sure to add `"wt":"json"` to `solrOptions` to request JSON as the output format.


## Executable
If you don't care about the source code,
[download the executable](https://github.com/hectorcorrea/solrdora/releases)
for your operating system, and follow the instructions on Release tab.
