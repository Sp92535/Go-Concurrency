package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type ScrapedData struct {
	ID    uint `gorm:"primaryKey"`
	URL   string
	Title string
	Links string
}

type Config struct {
	Parallelism int `json:"parallelism"`
	RandomDelay int `json:"random_delay"`
}

func InitDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("scraper.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.Migrator().DropTable(&ScrapedData{})	// Comment this line to persist previous data 
	db.AutoMigrate(&ScrapedData{}) 
}

func StartScraper() {

	// Enter urls here
	var urls = []string{
		"https://timeless-treasures-psi.vercel.app/",
		"https://infinityplay.vercel.app",
	}

	fmt.Println("Scraping started")

	var wg sync.WaitGroup

	// Running a goroutine for each url
	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}

	wg.Wait()

	fmt.Println("Scraping completed")
}

func fetchURL(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	collector := colly.NewCollector()

	config := loadConfig()
	// Rate limiting by colly
	collector.Limit(&colly.LimitRule{
		DomainRegexp: ".*",
		Parallelism:  config.Parallelism,                              // Use a lower Parallelism for sites with strict anti-scraping policies.
		RandomDelay:  time.Duration(config.RandomDelay) * time.Second, // Increase RandomDelay if getting blocked.
	})

	var data ScrapedData
	data.URL = url

	collector.OnHTML("title", func(e *colly.HTMLElement) {
		data.Title = e.Text
		// fmt.Printf("Title of %s: %s\n", url, e.Text)
	})

	collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		if link != "" {
			data.Links += e.Request.AbsoluteURL(e.Attr("href")) + ","
		}

		// fmt.Printf("Found link on %s: %s\n", url, link)
	})

	collector.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error scraping %s: %s\n", url, err)

		if r.StatusCode >= 500 {
			fmt.Printf("Retrying: %s", url)
			r.Request.Retry()
		}

	})

	collector.OnScraped(func(r *colly.Response) {
		db.Create(&data) // Save to database
	})

	collector.Visit(url)
}

func ExportToJSON(filename string) {
	fmt.Println("Writing to json")
	
	pathName := "output"
	if err := os.MkdirAll(pathName, os.ModePerm); err != nil {
		fmt.Printf("Error writing JSON %v", err)
		return
	}
	
	file, _ := os.Create(pathName + "/" + filename)
	defer file.Close()
	
	var data []ScrapedData
    db.Find(&data)

	json.NewEncoder(file).Encode(data)

}

func loadConfig() Config {
	file, _ := os.Open("config.json")
	defer file.Close()

	var config Config
	json.NewDecoder(file).Decode(&config)
	return config
}
