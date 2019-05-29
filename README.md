# get-href
Get hrefs

build:  
```
go get -u github.com/jidckii/get-href
make
```
Run:
```
  -area-filter string
        You filter for search href. (default "head")
  -find-filter string
        You filter for search href. (default "link[hreflang]")
  -json
        Response in json format.
  -url string
        URL for scraping (default "https://github.com/PuerkitoBio/goquery")
```
Example:  
```
build/linux/get-href -url https://you-domain/en
```
