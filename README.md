# webbundle
Tool for creating web bundles from a URL in Go.

```
go install github.com/WICG/webpackage/go/bundle/cmd/...@latest
```

```
cd cmd/crawler
go build
./crawler -url=https://nytimes.com
```

That will generate a urls.txt file that you can then use to create the webbundle with:

```
gen-bundle -primaryURL https://nytimes.com -URLList=urls.txt
```