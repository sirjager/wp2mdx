package service

type Category struct {
	ID          int           `json:"id"`
	Count       int           `json:"count"`
	Description string        `json:"description"`
	Link        string        `json:"link"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Taxonomy    string        `json:"taxonomy"`
	Parent      int           `json:"parent"`
	Meta        []interface{} `json:"meta"`
	Links       Links         `json:"_links"`
}

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Slug        string `json:"slug"`
	AvatarURLs  struct {
		Size24 string `json:"24"`
		Size48 string `json:"48"`
		Size96 string `json:"96"`
	} `json:"avatar_urls"`
	Links Links `json:"_links"`
}

type FeaturedMedia struct {
	ID    int    `json:"id"`
	Date  string `json:"date"`
	Slug  string `json:"slug"`
	Type  string `json:"type"`
	Link  string `json:"link"`
	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Author  int `json:"author"`
	Caption struct {
		Rendered string `json:"rendered"`
	} `json:"caption"`
	AltText      string `json:"alt_text"`
	MediaType    string `json:"media_type"`
	MimeType     string `json:"mime_type"`
	MediaDetails struct {
		Width    int    `json:"width"`
		Height   int    `json:"height"`
		File     string `json:"file"`
		FileSize int    `json:"filesize"`
		Sizes    struct {
			Medium    ImageSize `json:"medium"`
			Large     ImageSize `json:"large"`
			Thumbnail ImageSize `json:"thumbnail"`
			Full      ImageSize `json:"full"`
		}
	}
}

type ImageSize struct {
	File      string `json:"file"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	FileSize  int    `json:"filesize"`
	MimeType  string `json:"mime_type"`
	SourceURL string `json:"source_url"`
}

type Post struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	DateGMT string `json:"date_gmt"`
	GUID    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGMT string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Categories  []int  `json:"categories"`
	Tags        []int  `json:"tags"`

	Title struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`

	Excerpt struct {
		Rendered string `json:"rendered"`
	} `json:"excerpt"`

	Author        int `json:"author"`
	FeaturedMedia int `json:"featured_media"`

	Embed struct {
		Author        []User          `json:"author"`
		FeaturedMedia []FeaturedMedia `json:"wp:featuredmedia"`
	} `json:"_embedded"`

	Links Links `json:"_links"`
}

type Links struct {
	Self []struct {
		Href string `json:"href"`
	} `json:"self"`
	Collection []struct {
		Href string `json:"href"`
	} `json:"collection"`
	PostType []struct {
		Href string `json:"href"`
	} `json:"wp:post_type"`
}

type Tag struct {
	ID          int    `json:"id"`
	Count       int    `json:"count"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Taxonomy    string `json:"taxonomy"`
	Links       Links  `json:"_links"`
}

type Media struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	DateGMT string `json:"date_gmt"`
	GUID    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGMT string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Author      int `json:"author"`
	Description struct {
		Rendered string `json:"rendered"`
	} `json:"description"`
	Caption struct {
		Rendered string `json:"rendered"`
	} `json:"caption"`
	AltText   string `json:"alt_text"`
	MediaType string `json:"media_type"`
	MimeType  string `json:"mime_type"`
	Post      int    `json:"post"`
	SourceURL string `json:"source_url"`
	Embedded  struct {
		Author []User `json:"author"`
	} `json:"_embedded"`

	MediaDetails MediaDetails `json:"media_details"`

	Links Links `json:"_links"`
}

type MediaDetails struct {
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	File     string `json:"file"`
	FileSize int    `json:"filesize"`
	Sizes    struct {
		Medium                   ImageSize `json:"medium"`
		Thumbnail                ImageSize `json:"thumbnail"`
		MediumLarge              ImageSize `json:"medium_large"`
		OceanThumbM              ImageSize `json:"ocean-thumb-m"`
		OceanThumbML             ImageSize `json:"ocean-thumb-ml"`
		WebStoriesPosterPortrait ImageSize `json:"web-stories-poster-portrait"`
		WebStoriesPublisherLogo  ImageSize `json:"web-stories-publisher-logo"`
		WebStoriesThumbnail      ImageSize `json:"web-stories-thumbnail"`
		Full                     ImageSize `json:"full"`
	} `json:"sizes"`
}

type Page struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	DateGMT string `json:"date_gmt"`
	GUID    struct {
		Rendered string `json:"rendered"`
	} `json:"guid"`
	Modified    string `json:"modified"`
	ModifiedGMT string `json:"modified_gmt"`
	Slug        string `json:"slug"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	Link        string `json:"link"`
	Title       struct {
		Rendered string `json:"rendered"`
	} `json:"title"`
	Content struct {
		Rendered string `json:"rendered"`
	} `json:"content"`
	Protected bool `json:"protected"`
	Excerpt   struct {
		Rendered string `json:"rendered"`
	} `json:"excerpt"`
	Author        int   `json:"author"`
	FeaturedMedia int   `json:"featured_media"`
	Links         Links `json:"_links"`
	Embedded      struct {
		Author []User `json:"author"`
	} `json:"_embedded"`
}
