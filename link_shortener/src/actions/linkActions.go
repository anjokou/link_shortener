package actions

import (
	"errors"
	"link_shortener/analytics"
	"link_shortener/links"
	"net/url"
)

var linkStorage links.LinkDA
var analyticsEngine analytics.AnalyticsEngine

func InitLinkctions(injectLinkStorage links.LinkDA, injectAnalytcs analytics.AnalyticsEngine) {
	linkStorage = injectLinkStorage
	analyticsEngine = injectAnalytcs
}

func Create(link links.Link) (*links.Link, error) {
	returnLink, err := linkStorage.Save(link)

	if returnLink != nil && err == nil {
		analyticsEngine.Created(*returnLink)
	}

	if errors.Is(err, links.ErrUrlDuplicate{}) {
		returnLink, err = linkStorage.GetByUrl(link.ExternalURL)
	}

	return returnLink, err
}

func Get(linkId string) (*links.Link, error) {
	link, err := linkStorage.Get(linkId)

	if err == nil {
		analyticsEngine.Accessed(*link, 1)
	}

	return link, err
}

func Validate(link *links.Link) error {
	_, err := url.ParseRequestURI(link.ExternalURL)
	return err
}
