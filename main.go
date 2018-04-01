package main

import (
	"github.com/EddieYY/go-crawl-store/storeSpider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ComparePrice(c *gin.Context) {
	Search := c.Param("id")
	out := storeSpider.QureyStore(Search)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "大潤發": out["大潤發"], "家樂福": out["家樂福"]})
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
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ALL/:id", ComparePrice)
		v1.GET("/家樂福/:id", Carrefourcontroller)
		v1.GET("/大潤發/:id", Rtmartcontroller)
	}
	router.Run()
}
