# Golang Web Scraper

## Overview
This is a multithreaded web scraper built using **Go** and the **Colly** library. It scrapes web pages for titles and links, stores the data in an SQLite database, and supports exporting the scraped data to JSON.

## Features
- **Multithreading**: Uses Goroutines to scrape multiple URLs concurrently.
- **Colly Integration**: Efficient HTML parsing and rate-limiting.
- **SQLite Database**: Stores scraped data.
- **JSON Export**: Saves the extracted data in a structured JSON format.
- **Configurable Parameters**: Allows control over parallelism and request delays using a `config.json` file.

## Installation
### Prerequisites
- Install [Go](https://go.dev/doc/install)
- Install required dependencies:
  ```sh
  go mod tidy
  ```

## Usage
### 1. Configure Parameters
Modify `config.json` to set parallelism and random delay:
```json
{
    "parallelism": 2,
    "random_delay": 3
}
```

### 2. Initialize Database and Run Scraper
Run the scraper to initialize the database and start scraping:
```sh
 go run main.go
```
This will start scraping the predefined URLs and store the results in `scraper.db`.

### 3. Export Data to JSON
The scraper automatically exports data to `results.json` after scraping.

## File Structure
```
.
├── scraper
│   ├── scraper.go     # Main scraper logic
│   ├── config.json    # Configuration file for parallelism and delays
│   ├── scraper.db     # SQLite database storing scraped data
│   ├── output/        # Folder for exported JSON data
│   ├── go.mod         # Go module dependencies
│   ├── go.sum         # Checksums
├── main.go            # Entry point
```

## Notes
- To **retain previous scraped data**, comment out the `db.Migrator().DropTable(&ScrapedData{})` line in `InitDB()`.
- To add more URLs, modify the `urls` list in `scraper.StartScraper()`.
- The `ExportToJSON("results.json")` function is automatically called in `main.go` to save the scraped data.

## License
This project is open-source and available under the MIT License.

