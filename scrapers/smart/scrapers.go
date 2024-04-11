package smart

import (
	"log"

	scraper "github.com/ferretcode-freelancing/sportsbook-scraper/scrapers"
	"gorm.io/gorm"
)

type Scrapers struct {
	BetOnline BetOnline
}

func GetScrapers(db *gorm.DB) Scrapers {
	betonline, err := NewBetOnline(db)

	if err != nil {
		log.Fatal(err)
	}

	return Scrapers{
		BetOnline: betonline,
	}
}

func (s *Scrapers) StartScrapers(
	newProps chan scraper.Props,
	errChan chan error,
	fatalError chan scraper.FatalError,
) {
	go s.BetOnline.Scraper.GetProps(newProps, errChan, fatalError)
}

func HandleError(err error, errChan chan error) {
	if err != nil {
		errChan <- err
	}
}
