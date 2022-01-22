package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"

	e "github.com/VxVxN/market_analyzer/pkg/error"
)

func UnmarshalRequest(c *gin.Context, reqStruct interface{}) *e.ErrObject {
	if err := c.BindJSON(&reqStruct); err != nil {
		return e.NewError("Bad Request", http.StatusBadRequest, err)
	}

	return nil
}
