package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

type PageInfo struct {
	PageNo     int         `json:"pageNo"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Records    interface{} `json:"records"`
}

//分页封装
func Paginate(page *PageInfo, c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		pageNo, _ := strconv.Atoi(c.Query("pageNo"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		if pageNo == 0 {
			pageNo = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize

		page.PageNo = pageNo
		page.PageSize = pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
