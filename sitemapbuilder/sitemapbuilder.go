package sitemapbuilder

import (
	"net/url"
	"strings"

	"github.com/mohsalsaleem/sitemap/apiclient"
	"github.com/mohsalsaleem/sitemap/htmlparser"
	"github.com/mohsalsaleem/sitemap/link"
	"golang.org/x/net/html"
)

type urlset struct {
	xmlns string
}

type sitemap struct {
	urlset urlset
}

func build(links []link.Link) {

}

// Process - Process URL and build sitemap
func Process(rawurl string) error {
	_url, err := url.Parse(rawurl)
	if err != nil {
		return err
	}

	htmlString, err := apiclient.MakeHTTPGetRequest(_url)
	if err != nil {
		return err
	}

	doc, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return err
	}

	links, err := htmlparser.CollectLinks(doc, _url.String())
	if err != nil {
		return err
	}
	build(links)

	return nil
}
