// Created by GoLand.
// User: huang.peng@datatom.com
// Date: 2020-02-10
// Time: 20:34

package version

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	HashNumber   string
	CommitNumber string
)

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"result": gin.H{
			"HashNumber":   HashNumber,
			"CommitNumber": CommitNumber,
		},
	})
}
