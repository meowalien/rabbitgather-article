package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/meowalien/rabbitgather-article/sec/logger"
)

// 文章基本的增刪查改
type Basic struct {
}

//type Article struct {
//	Id int `gorm:"primaryKey"`
//	Title string //`gorm:"uniqueIndex"`
//	Content string //`gorm:"uniqueIndex"`
//	CreateTime time.Time
//	UpdateTime time.Time
//}



func (receiver Basic) Get(ctx *gin.Context) {
	type request struct {
		ID int `form:"id" binding:"required"`
	}
	var req request
	err := ctx.ShouldBindQuery(&req)

	if !ErrorCheck(err,ctx,WrongFormat) {
		return
	}
	logger.Logger.Debug("req: ", req)

	//mariadb.GlobalConn.Select()

}
func (receiver Basic) POST(ctx *gin.Context) {

}
func (receiver Basic) DELETE(ctx *gin.Context) {

}
func (receiver Basic) PATCH(ctx *gin.Context) {

}
