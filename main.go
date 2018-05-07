package main

import (
	//"fmt"
	"github.com/EddieYY/go-crawl-store/storeSpider"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ComparePrice(c *gin.Context) {
	Search := c.Param("id")
	var wg sync.WaitGroup
	out := make([]storeSpider.StorePrice, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		out[0] = storeSpider.RtmartSpider(Search)
	}()
	go func() {
		defer wg.Done()
		out[1] = storeSpider.CarrefourSpider(Search)
	}()
	//storeSpider.QureyStore(Search, out)
	wg.Wait()
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
