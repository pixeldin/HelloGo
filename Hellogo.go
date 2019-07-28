package main

import (
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strconv"
	"os"
	"image/jpeg"
	"github.com/corona10/goimagehash"
)

const scrapURL = "https://indienova.com/channel/news/page/"

func ExampleScrape() {

	//fmt.Println(doc.Html())
	sum := 0
	for pn := 1; pn <= 15; pn++ {
		nextScrapURL := scrapURL + strconv.Itoa(pn)
		// Request the HTML page.
		fmt.Println("Climb Page", pn, "...")
		res, err := http.Get(nextScrapURL)

		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find("div .article-panel").Each(func(i int, contentSelection *goquery.Selection) {
			title := contentSelection.Find("h4 a").Text()
			like := contentSelection.Find(".views-badge .number-first").Text()
			comment := contentSelection.Find(".number-last").Text()
			fmt.Print("第", pn, "页, 第", i+1, "个帖子,标题：\n")
			//img url
			img_URL, _ := contentSelection.Find("img").Attr("src")
			fmt.Println(title, "\t点赞:", like, "\t评论:", comment, ",imgURL:", img_URL)
			sum++
		})

	}
	fmt.Println("Climb all of :", sum)

}

//return two images diff level with averHash,differenceHash
func imgDiff(loca1, loca2 string) (averDistance, difDistance int) {
	file1, err := os.Open(loca1)
	defer file1.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	file2, err := os.Open(loca2)
	defer file2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	img1, _ := jpeg.Decode(file1)
	img2, _ := jpeg.Decode(file2)
	hash1, _ := goimagehash.AverageHash(img1)
	hash2, _ := goimagehash.AverageHash(img2)
	averDistance, _ = hash1.Distance(hash2)
	fmt.Printf("Distance between images, AverageHash: %v\n", averDistance)

	hash1, _ = goimagehash.DifferenceHash(img1)
	hash2, _ = goimagehash.DifferenceHash(img2)
	difDistance, _ = hash1.Distance(hash2)
	fmt.Printf("Distance between images, DifferenceHash: %v\n", difDistance)
	return averDistance, difDistance
}

func main() {
	//ExampleScrape()
	path, _ := os.Getwd()
	img1 := path + "\\Imgs\\2.jpg"
	img2 := path + "\\Imgs\\1.jpg"
	imgDiff(img1, img2)
}
