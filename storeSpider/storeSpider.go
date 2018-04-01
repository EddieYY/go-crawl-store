package storeSpider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"log"
	"net/http"
	"strings"
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

func CarrefourSpider(Search string) []UnitPrice {
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
	//out := StorePrice{"大潤發"}
	var Priceout []UnitPrice
	fmt.Printf("Start Search %s in 家樂福 \n", Search)
	for i := range sel.Nodes {
		single := sel.Eq(i)
		script := single.Text()
		if strings.Index(script, "searchProductListModel") > -1 {
			vm := otto.New()
			//vm.Run(`var dataLayer = []`)
			vm.Run(script)
			if value, err := vm.Get("searchProductListModel"); err == nil {
				goData, _ := value.Export()
				//mapData := goData.([]map[string]interface{})
				if goData != nil && strings.Index(goData.(string), "Price") > -1 {
					vm.Run(`var JSS = JSON.parse(searchProductListModel).ProductListModel`)
					v, _ := vm.Get("JSS")
					ve, _ := v.Export()
					DT := ve.([]map[string]interface{})
					for i := range DT {
						fmt.Printf("家樂福 物品:%s, 數量:%s, 價格:NT$ %s, 折扣價:NT%s  \n", DT[i]["Name"],
							DT[i]["ItemQtyPerPackFormat"], DT[i]["Price"], DT[i]["SpecialPrice"])
						Priceout = append(Priceout, UnitPrice{DT[i]["Name"].(string) + " " + DT[i]["ItemQtyPerPackFormat"].(string),
							"NTD $" + DT[i]["Price"].(string), DT[i]["SpecialPrice"].(string)})
					}

				}

			}
		}
	}
	return Priceout
}

func RtmartSpider(Search string) []UnitPrice {
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
	var Priceout []UnitPrice
	fmt.Printf("Start Search %s in 大潤發 \n", Search)
	for i := range sel.Nodes {
		single := sel.Eq(i)
		script := single.Text()
		if strings.Index(script, "dataLayer") > -1 {
			vm := otto.New()
			vm.Run(`var dataLayer = []`)
			vm.Run(script)
			vm.Run(`var js = dataLayer[0].ecommerce.impressions`)
			if value, err := vm.Get("js"); err == nil {
				goData, _ := value.Export()
				if goData != nil {
					DT := goData.([]map[string]interface{})
					for i := range DT {
						fmt.Printf("大潤發 物品:%s, 價格:NT$ %s, 折扣價:NT%s  \n", DT[i]["name"], DT[i]["price"], "")
						Priceout = append(Priceout, UnitPrice{DT[i]["name"].(string), "NTD $" + DT[i]["price"].(string), ""})
					}
				}
			}
		}
	}
	return Priceout
}

func QureyStore(Search string) map[string][]UnitPrice {
	out := make(map[string][]UnitPrice, 2)
	out["大潤發"] = RtmartSpider(Search)
	out["家樂福"] = CarrefourSpider(Search)
	return out
}

/*
func main() {
	Search := "可口可樂"
	out := QureyStore(Search)
	fmt.Print(out)
}
*/
