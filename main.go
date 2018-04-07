package main

import (
	"github.com/EddieYY/go-crawl-store/storeSpider"
	"github.com/gin-gonic/gin"
	//"golang.org/x/sync/errgroup"
	//"log"
	//"github.com/gin-contrib/cache"
	//"github.com/gin-contrib/cache/persistence"
	"net/http"
	"runtime"
	//"time"
)

/*func ComparePrice() http.Handler {
	//Search := c.Param("id")
	//out := storeSpider.QureyStore(Search)
	//c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "大潤發": out["大潤發"], "家樂福": out["家樂福"]})

	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/api/v1/ALL/:id", func(c *gin.Context) {
		Search := c.Param("id")
		out := storeSpider.QureyStore(Search)
		c.JSON(
			http.StatusOK,
			gin.H{"status": http.StatusOK, "大潤發": out["大潤發"], "家樂福": out["家樂福"]},
		)
	})

	return e

}
*/

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func ComparePrice(c *gin.Context) {
	Search := c.Param("id")
	out1 := make(chan []storeSpider.UnitPrice)
	out2 := make(chan []storeSpider.UnitPrice)
	go func() { out1 <- storeSpider.RtmartSpider(Search) }()
	go func() { out2 <- storeSpider.CarrefourSpider(Search) }()
	//storeSpider.QureyStore(Search, out)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "大潤發": <-out1, "家樂福": <-out2})
}

func Carrefourcontroller(c *gin.Context) {
	Search := c.Param("id")
	out := storeSpider.CarrefourSpider(Search)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "家樂福": out})

}

func Rtmartcontroller(c *gin.Context) {
	Search := c.Param("id")
	out := storeSpider.RtmartSpider(Search)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "大潤發": out})

}

func main() {
	router := gin.Default()
	//store := persistence.NewInMemoryStore(60 * time.Second)
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ALL/:id", ComparePrice)
		//cache.CachePage(store, time.Minute
		v1.GET("/家樂福/:id", Carrefourcontroller)
		v1.GET("/大潤發/:id", Rtmartcontroller)
	}

	/*router.GET("/api/ALL/:id", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		//cCp := c.Copy()
		go func() {
			Search := c.Param("id")
			out := storeSpider.QureyStore(Search)
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "大潤發": out["大潤發"], "家樂福": out["家樂福"]})
		}()
	})*/
	router.Run()
}

/*func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      ComparePrice(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	g.Go(func() error {
		return server01.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}
*/
