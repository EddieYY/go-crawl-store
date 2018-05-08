package main

import (
	"fmt"
	"github.com/EddieYY/go-crawl-store/storeSpider"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	//"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ComparePrice(c *gin.Context) {
	Search := c.Param("id")

	outt := make(chan storeSpider.StorePrice)

	go func() { outt <- storeSpider.RtmartSpider(Search) }()
	go func() { outt <- storeSpider.CarrefourSpider(Search) }()

	var out []storeSpider.StorePrice

	for i := 0; i < 2; i++ {
		select {
		case unit := <-outt:
			out = append(out, unit)
		case <-time.After(6 * time.Second):
			fmt.Println("timeout")
			continue
		}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": out})
}

func Carrefourcontroller(c *gin.Context) {
	Search := c.Param("id")
	out := storeSpider.CarrefourSpider(Search)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": out})

}

func Rtmartcontroller(c *gin.Context) {
	Search := c.Param("id")
	out := storeSpider.RtmartSpider(Search)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": out})

}

func main() {
	router := gin.Default()
	store := persistence.NewInMemoryStore(time.Minute * 5)

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ALL/:id", cache.CachePage(store, time.Minute*5, ComparePrice))
		v1.GET("/家樂福/:id", cache.CachePage(store, time.Minute*5, Carrefourcontroller))
		v1.GET("/大潤發/:id", cache.CachePage(store, time.Minute*5, Rtmartcontroller))
	}

	router.Run()
}
