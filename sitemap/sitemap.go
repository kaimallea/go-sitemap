package sitemap

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

// Represents a <sitemapindex> element
type SitemapIndex struct {
	XMLName          xml.Name          `xml:"sitemapindex"`
	SitemapLocations []SitemapLocation `xml:"sitemap"`
}

// Represents a <sitemap> element within a <sitemapindex> element
type SitemapLocation struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

// Represents a <urlset> element
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Urls    []Url    `xml:"url"`
}

// Represents a <url> element within a <urlset>
type Url struct {
	Loc        string `xml:"loc"`
	Lastmod    string `xml:"lastmod"`
	Changefreq string `xml:"changefreq"`
	Priority   string `xml:"priority"`
}

// Return number of <sitemap> elements within a <sitemapindex>
func (s SitemapIndex) Length() int {
	return len(s.SitemapLocations)
}

// Return number of <url> elements within a <urlset> element
func (s Sitemap) Length() int {
	return len(s.Urls)
}

func (s SitemapLocation) String() string {
	return s.Loc
}

func (u Url) String() string {
	return u.Loc
}

// Fetch a remote XML file and return its data
func fetchXML(url string) (data []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return
}

// Fetch a remote sitemap index
func FetchSitemapIndex(url string) (s SitemapIndex, err error) {
	xmlData, err := fetchXML(url)
	if err != nil {
		return
	}

	s = SitemapIndex{}
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		return
	}

	return
}

// Fetch a sitemap
func FetchSitemap(url string) (s Sitemap, err error) {
	xmlData, err := fetchXML(url)
	if err != nil {
		return
	}

	s = Sitemap{}
	err = xml.Unmarshal(xmlData, &s)
	if err != nil {
		return
	}

	return
}
