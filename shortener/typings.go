package shortener

type ShortenLinkRequestItem struct {
	Domain      string `json:"domain"`
	OriginalUrl string `json:"originalURL"`
}

type ShortenLinkResponseItem struct {
	OriginalUrl    string `json:"originalURL"`
	ShortURL       string `json:"shortURL"`
	SecureShortUrl string `json:"secureShortURL"`
	DomainId       int    `json:"DomainId"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Cloaking       bool   `json:"cloaking"`
	Archived       bool   `json:"archived"`
	Duplicate      bool   `json:"duplicate"`
}
