package services

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// サンプルデータ全件取得
func GetAllSamples(ctx *gin.Context) {
	result, err := models.SelectAllSamples()
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

// サンプルデータのID検索
func GetSampleByID(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := models.SelectSampleByID(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

// 　サンプルデータの登録
func PostSample(ctx *gin.Context) {
	var json models.Sample
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := models.InsertSample(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"id": result})
}

// 　サンプルデータの更新
func PutSample(ctx *gin.Context) {
	var json models.Sample
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateSample(&json)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}

// 　サンプルデータの削除
func DeleteSample(ctx *gin.Context) {
	id := ctx.Param("id")
	err := models.DeleteSample(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}
