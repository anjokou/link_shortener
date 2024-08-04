package analytics

import "link_shortener/links"

type AnalyticsEngine interface {
	Created(link links.Link)
	Accessed(link links.Link, timesAccessed int)
}
