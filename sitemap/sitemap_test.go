package sitemap

import "testing"

func TestFetchXML(t *testing.T) {
	url := "http://www.google.com/sitemaps_webmasters.xml"
	_, err := fetchXML(url)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFetchSitemapIndex(t *testing.T) {
	url := "http://www.google.com/sitemaps_webmasters.xml"
	_, err := FetchSitemapIndex(url)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFetchSitemap(t *testing.T) {
	url := "http://www.google.com/ventures/sitemap_ventures.xml"
	_, err := FetchSitemap(url)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
