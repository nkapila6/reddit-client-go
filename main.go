package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Post struct {
	Subreddit    string `json:"subreddit"`
	Title        string `json:"title"`
	Author       string `json:"author"`
	CommentCount string `json:"comments"`
	Permalink    string `json:"permalink,omitempty"`
	URL          string `json:"url,omitempty"`
	Image        string `json:"image,omitempty"`
	Domain       string `json:"domain,omitempty"`
	Time         string `json:"time,omitempty"`
}

func main() {
	sub := flag.String("sub", "golang", "subreddit name (no /r/ prefix)")
	limit := flag.Int("limit", 25, "max posts to output")
	pages := flag.Int("pages", 1, "max pages to follow on old.reddit.com (each page ~25 posts)")
	sort := flag.String("sort", "hot", "hot|new|top|rising|controversial")
	timeFilter := flag.String("t", "all", "top time filter: hour|day|week|month|year|all (only used for sort=top)")
	imagesOnly := flag.Bool("images", false, "only images or all post types")
	flag.Parse()

	startURL := buildOldRedditURL(*sub, *sort, *timeFilter)

	c := colly.NewCollector(
		colly.AllowedDomains("old.reddit.com"),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:125.0) Gecko/20100101 Firefox/125.0"),
	)
	c.SetRequestTimeout(20 * time.Second)

	var posts []Post
	seen := make(map[string]struct{})
	nextURL := ""

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.9")
		r.Headers.Set("Referer", "https://old.reddit.com/")
	})

	c.OnHTML("div.thing", func(h *colly.HTMLElement) {
		if len(posts) >= *limit {
			return
		}

		title := strings.TrimSpace(h.ChildText("a.title"))
		author := h.Attr("data-author")
		permalink := h.Request.AbsoluteURL(h.Attr("data-permalink"))
		num_comments := h.Attr("data-comments-count")

		ms, err := strconv.ParseInt(h.Attr("data-timestamp"), 10, 64)
		t := ""
		if err == nil {
			t = time.Unix(0, ms*int64(time.Millisecond)).In(time.Local).Format("Monday, January 2, 2006 at 3:04:05 PM")
		}

		if permalink == "" {
			permalink = h.Request.AbsoluteURL(h.ChildAttr("a.comments", "href"))
		}
		if permalink == "" {
			return
		}

		postURL := h.Request.AbsoluteURL(h.Attr("data-url"))
		if postURL == "" || postURL == permalink {
			postURL = permalink
		}

		domain := strings.TrimSpace(h.Attr("data-domain"))

		if _, ok := seen[permalink]; ok {
			return
		}
		seen[permalink] = struct{}{}

		img := ""
		if isImageURL(postURL) {
			img = postURL
			postURL = ""
		} else {
			thumbHref := h.Request.AbsoluteURL(h.ChildAttr("a.thumbnail", "href"))
			if isImageURL(thumbHref) {
				img = thumbHref
			} else {
				thumbSrc := h.Request.AbsoluteURL(h.ChildAttr("a.thumbnail img", "src"))
				if isImageURL(thumbSrc) {
					img = thumbSrc
				}
			}
		}

		if *imagesOnly && img == "" {
			return
		}

		posts = append(posts, Post{
			Subreddit:    *sub,
			Title:        title,
			CommentCount: num_comments,
			Author:       author,
			Permalink:    permalink,
			// URL:       postURL,
			Time:   t,
			Image:  img,
			Domain: domain,
		})
	})

	c.OnHTML("span.next-button a", func(h *colly.HTMLElement) {
		nextURL = h.Attr("href")
	})

	c.OnError(func(r *colly.Response, err error) {
		if r != nil {
			body := string(r.Body)
			if len(body) > 500 {
				body = body[:500] + "...[truncated]"
			}
			log.Fatalf("request failed: %s -> %d: %v\n%s", r.Request.URL, r.StatusCode, err, body)
		}
		log.Fatalf("request failed: %v", err)
	})

	// loop thru pages
	cur := startURL
	for page := 0; page < *pages && cur != "" && len(posts) < *limit; page++ {
		nextURL = ""
		if err := c.Visit(cur); err != nil {
			log.Fatal(err)
		}
		cur = nextURL
	}

	// json encoder to stdout
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(posts)
}

func buildOldRedditURL(sub, sort, t string) string {
	escaped := url.PathEscape(sub)

	u := fmt.Sprintf("https://old.reddit.com/r/%s/%s/", escaped, sort)

	q := url.Values{}
	q.Set("limit", "100") // old reddit caps anyway; we paginate with "next"
	if sort == "top" {
		q.Set("t", t)
	}

	return u + "?" + q.Encode()
}

func isImageURL(u string) bool {
	pu, err := url.Parse(u)
	if err != nil {
		return false
	}
	host := strings.ToLower(pu.Host)
	if host == "i.redd.it" || host == "preview.redd.it" || host == "i.imgur.com" {
		return true
	}
	p := strings.ToLower(pu.Path)
	return strings.HasSuffix(p, ".jpg") ||
		strings.HasSuffix(p, ".jpeg") ||
		strings.HasSuffix(p, ".png") ||
		strings.HasSuffix(p, ".gif") ||
		strings.HasSuffix(p, ".webp")
}
