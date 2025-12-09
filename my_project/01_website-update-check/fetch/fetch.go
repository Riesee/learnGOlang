package fetch

import (
	"context"
	"io"
	"net/http"
	"time"
	"website-checker/logger"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func Get(url string, selector string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Failed to fetch the URL","error", err)
		return nil, err
	}
	decoder := charmap.ISO8859_9.NewDecoder()
	utf8Reader := transform.NewReader(resp.Body, decoder)

	defer resp.Body.Close()
	if selector != "" {
		doc, err := goquery.NewDocumentFromReader(utf8Reader)
		if err != nil {
			logger.Error("Failed to parse the HTML document","error", err)
			return nil, err
		}

		selectorFind, err := doc.Find(selector).Html()
		if err != nil {
			logger.Error("Failed to parse the HTML document","error", err)
			return nil, err
		}

		return []byte(selectorFind), nil
	} else {
		body, err := io.ReadAll(resp.Body)
		logger.Info(string(body),"body")
		if err != nil {
			return nil, err
		}
		return body, nil

	}
}


func GetWithBrowser(url string, selector string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var result string

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(selector),
		chromedp.Text(selector, &result, chromedp.ByQuery),
	)
	

	logger.Info("result", "New Price", result)


	if err != nil {
		logger.Error("Failed to fetch the URL","error", err)
		return nil, err
	}

	return []byte(result), nil
}
