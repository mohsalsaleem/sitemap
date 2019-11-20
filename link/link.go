package link

import (
	"net/url"
	"strings"
)

// Link - the link datastructure
type Link struct {
	href string
	url  *url.URL
}

// New - Create a new link
func New(href string, domain string) (*Link, error) {
	link := Link{}
	link.href = href
	if strings.Contains(href, domain) {
		url, err := url.Parse(href)
		if err != nil {
			return nil, err
		}
		link.url = url
	} else {
		url, err := url.Parse(domain + href)
		if err != nil {
			return nil, err
		}
		link.url = url
	}

	return &link, nil
}
