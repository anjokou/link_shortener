package routers

import (
	"errors"
	"link_shortener_analytics/actions"
	"link_shortener_analytics/metrics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyMetricsRoutes(engine *gin.Engine) {
	engine.GET("links/:linkId", getLink)
}

func getLink(c *gin.Context) {
	id := c.Param("linkId")
	metric, err := actions.Get(id)

	if err == nil {
		c.JSON(http.StatusOK, metric)
	} else if errors.Is(err, metrics.ErrNotFound{}) {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}
