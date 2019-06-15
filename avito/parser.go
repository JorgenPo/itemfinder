package avito

import (
	"findthing/types"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type Parser struct {

}

func parseUser(page *goquery.Document, item *types.Item) {
	name := page.Find(".seller-info-name a").First(); if len(name.Text()) > 0 {
		item.User.Name = strings.Trim(name.Text(), "\n ")
	}

	userType := page.Find(".seller-info-col").Children().Eq(1); if len(userType.Text()) > 0 {
		if userType.Text() == "Частное лицо" {
			item.User.IsCompany = false
		} else {
			item.User.IsCompany = true
		}
	}

	additionalInfo := page.Find(".seller-info-col").Children().Eq(2); if len(additionalInfo.Text()) > 0 {
		item.User.RegisterDate = strings.Trim(additionalInfo.Text(), "\n ")
	}
}

func (*Parser) Parse(itemUrl string) (*types.Item, error) {
	resp, err := http.Get(itemUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to '%v': %v", itemUrl, err)
	}

	defer resp.Body.Close()
	page, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse a page '%v': %v", itemUrl, err)
	}

	var item types.Item

	title := page.Find(".title-info-title-text"); if len(title.Text()) > 0 {
		item.Title = title.Text()
	}

	phone := page.Find(".item-phone-button-sub-text"); if len(phone.Text()) > 0 {
		item.PhoneNumber = phone.Text()
	}

	desc := page.Find(".item-description-text > p"); if len(desc.Text()) > 0 {
		item.Description = desc.Text()
	}

	address := page.Find(".item-map-address"); if len(address.Text()) > 0 {
		item.Address = strings.Trim(address.Text(), "\n ")
	}

	price := page.Find(".js-item-price").First(); if len(price.Text()) > 0 {
		priceText := strings.Replace(price.Text(), " ", "", -1)
		item.Price, _ = strconv.Atoi(priceText)
	}


	parseUser(page, &item)

	return &item, nil
}


