package shortener

func (c *Client) mapUrlToRequest(link string) ShortenLinkRequestItem {
	return ShortenLinkRequestItem{
		OriginalUrl: link,
		Domain:      c.config.Domain,
	}
}

func (c *Client) mapUrlsToBulkRequest(links []string) []ShortenLinkRequestItem {
	result := []ShortenLinkRequestItem{}
	for _, i := range links {
		result = append(result, c.mapUrlToRequest(i))
	}
	return result
}
