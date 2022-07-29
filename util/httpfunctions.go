package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": ""})
		return
	}
}

func BadStatusReturn(c *gin.Context, err error) {
	if err != nil {
		c.IndentedJSON(http.StatusNotAcceptable, err)
		return
	}
}

func CreateOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusCreated, gin.H{"New object registred": obj})

}
func FoundOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusFound, gin.H{"Registred object Found": obj})
}

func ModifiedOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusOK, gin.H{"Registred object modified": obj})
}

func DeleteOrNotStatusReturn(err error, c *gin.Context, obj interface{}) {
	BadStatusReturn(c, err)
	c.JSON(http.StatusOK, gin.H{"Registred object Deleted": obj})
}
