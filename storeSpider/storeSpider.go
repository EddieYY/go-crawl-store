package storeSpider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
	"strings"
	"sync"
)

type UnitPrice struct {
	Name         string
	Price        string
	SpecialPrice string
}

type StorePrice struct {
	Store string
	Body  []UnitPrice
}

func CarrefourSpider(Search string) StorePrice {
	// Request the HTML page.
	res, err := http.Get("https://online.carrefour.com.tw/search?key=" + Search)
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
	sel := doc.Find("script")

	fmt.Printf("Start Search %s in 家樂福 \n", Search)

	mainnode := make(chan string)

	for i := range sel.Nodes {
		//fmt.Print(i)
		go func(i int) {
			single := sel.Eq(i)
			script := single.Text()
			if strings.Index(script, "searchProductListModel") > -1 {
				mainnode <- script
				close(mainnode)
			}
		}(i)
	}
	//fmt.Println("%s", mainnode)
	v := <-mainnode
	//fmt.Println(v)
	vm := otto.New()
	vm.Run(v)
	// taker js data
	var DT []map[string]interface{}
	if value, err := vm.Get("searchProductListModel"); err == nil {
		goData, _ := value.Export()
		if goData != nil && strings.Index(goData.(string), "Price") > -1 {
			vm.Run(`var JSS = JSON.parse(searchProductListModel).ProductListModel`)
			v, _ := vm.Get("JSS")
			ve, _ := v.Export()
			DT = ve.([]map[string]interface{})
		}
	}

	size := make(chan UnitPrice, len(DT))
	var wg sync.WaitGroup
	for i := range DT {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("家樂福 物品:%s, 價格:NT$ %s, 折扣價:NT$ %s  \n", DT[i]["Name"], DT[i]["Price"], DT[i]["SpecialPrice"])
			size <- UnitPrice{DT[i]["Name"].(string), "NTD $" + DT[i]["Price"].(string), DT[i]["SpecialPrice"].(string)}
		}(i)
	}
	go func() {
		wg.Wait()
		close(size)
	}()
	var Priceout []UnitPrice
	for t := range size {
		Priceout = append(Priceout, t)
	}

	return StorePrice{"家樂福", Priceout}
}

func RtmartSpider(Search string) StorePrice {
	// Request the HTML page.
	res, err := http.Get("http://www.rt-mart.com.tw/direct/index.php?action=product_search&prod_keyword=" + Search)
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
	sel := doc.Find("script")
	//out := StorePrice{"大潤發"}
	//var Priceout []UnitPrice
	fmt.Printf("Start Search %s in 大潤發 \n", Search)

	mainnode := make(chan string)

	for i := range sel.Nodes {
		//fmt.Print(i)
		go func(i int) {
			single := sel.Eq(i)
			script := single.Text()
			if strings.Index(script, "dataLayer.push") > -1 {
				mainnode <- script
				close(mainnode)
			}
		}(i)
	}
	v := <-mainnode
	//fmt.Println(v)
	vm := otto.New()
	vm.Run(`var dataLayer = []`)
	vm.Run(v)
	vm.Run(`var js = dataLayer[0].ecommerce.impressions`)
	var DT []map[string]interface{}
	if value, err := vm.Get("js"); err == nil {
		goData, _ := value.Export()
		DT = goData.([]map[string]interface{})
	}

	size := make(chan UnitPrice, len(DT))
	var wg sync.WaitGroup
	for i := range DT {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("大潤發 物品:%s, 價格:NT$ %s, 折扣價:NT$ %s  \n", DT[i]["name"], DT[i]["price"], "")
			size <- UnitPrice{DT[i]["name"].(string), "NTD $" + DT[i]["price"].(string), ""}
		}(i)
	}

	go func() {
		wg.Wait()
		close(size)
	}()

	var Priceout []UnitPrice
	for t := range size {
		Priceout = append(Priceout, t)
	}

	return StorePrice{"大潤發", Priceout}
}
