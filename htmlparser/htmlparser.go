package htmlparser

import (
	"github.com/mohsalsaleem/sitemap/link"
	"golang.org/x/net/html"
)

func getAttributeMap(node *html.Node) map[string]string {
	attributes := node.Attr
	attrMap := make(map[string]string)
	for _, attribute := range attributes {
		attrMap[attribute.Key] = attribute.Val
	}
	return attrMap
}

// CollectLinks - Collects all the html links in a HTML Node
func CollectLinks(htm *html.Node, domain string) ([]link.Link, error) {
	links := make([]link.Link, 0)
	var _err error = nil
	var crawler func(*html.Node)
	crawler = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			attrMap := getAttributeMap(node)
			link, err := link.New(attrMap["href"], domain)
			if err != nil {
				_err = err
				return
			}
			links = append(links, *link)
		}
	}
	crawler(htm)
	if _err != nil {
		return nil, _err
	}
	return links, nil
}
