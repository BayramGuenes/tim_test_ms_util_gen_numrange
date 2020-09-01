package main

import (
	"log"
	"net/http"

	boot "tim_presse/tim_test_ms_util_gen_numrange/bootstrap"

	"github.com/gin-gonic/gin"
)

var ()

func main() {
	routerGinGonic()
}

func routerGinGonic() {
	boot.LoadExternalSettings()
	_, err := boot.GetApplConf()
	if err != nil {
		return
	}

	router := gin.New()

	router.GET("/", amAliveHandler)

	// -----------------------------------------------------------------//
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./public")
	router.GET("/NumRangeServices", gotoMonitor)
	router.POST("/performOp", servicesPerformOp)
	// -----------------------------------------------------------------//

	log.Print("Starting Server on Port", ":"+boot.ThisPort) //setImageServiceRoutes(router)

	http.ListenAndServe(":"+boot.ThisPort, router)
}
func amAliveHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "ps_sysreg (micro-)service is alive")
}
