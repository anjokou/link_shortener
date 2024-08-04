package routers

import (
	"errors"
	"link_shortener/actions"
	"link_shortener/links"
	"net/http"

	"github.com/gin-gonic/gin"
)

var hashFunction links.Hasher

type createRequestBody struct {
	Link string `json:"link"`
}

const createBodyKey = "body"
const linkKey = "link"

func InitLinks(hasher links.Hasher) {
	hashFunction = hasher
}

func ApplyLinkRoutes(engine *gin.Engine) {
	engine.POST("/links", parseLink, validateLink, addLink)
	engine.GET("/:id", redirect)
}

func parseLink(c *gin.Context) {
	body := new(createRequestBody)
	err := c.BindJSON(body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	c.Set(createBodyKey, body)

	c.Next()
}

func validateLink(c *gin.Context) {
	body := c.MustGet(createBodyKey).(*createRequestBody)

	link := &links.Link{
		Id:          string(hashFunction.Hash(body.Link)),
		ExternalURL: body.Link,
	}

	err := actions.Validate(link)

	if err == nil {
		c.Set(linkKey, link)
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err.Error()))
	}
}

func addLink(c *gin.Context) {
	link := c.MustGet(linkKey).(*links.Link)

	returnLink, err := actions.Create(*link)

	if err == nil {
		c.JSON(http.StatusOK, successRespose(returnLink.Id))
	} else {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

func redirect(c *gin.Context) {
	link, err := actions.Get(c.Param("id"))

	if err == nil {
		c.Redirect(http.StatusPermanentRedirect, link.ExternalURL)
	} else if errors.Is(err, links.ErrNotFound{}) {
		c.JSON(http.StatusNotFound, errorResponse("Not found"))
	} else {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}
}

func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}

func successRespose(linkId string) gin.H {
	return gin.H{"link": linkId}
}
