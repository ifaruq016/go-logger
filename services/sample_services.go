package services

import (
	// "log"
	// "fmt"
	"github.com/gin-gonic/gin"

	"logger/app/utils"
)

type sampleCodeRequest struct {
	QRContent           string `json:"qr_content"`
	ReferenceNumber     string `json:"reference_number"`
	TransactionAmount   string `json:"transaction_amount"`
	ServiceChargeAmount string `json:"service_charge_amount"`
}

func RequestServices(ctx *gin.Context, apiIdentifier string) (int, gin.H, error) { 
	// var requestJSON sampleCodeRequest

	// if err := ctx.ShouldBindJSON(&requestJSON); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(requestJSON.QRContent)
	go utils.LogInfo(ctx, apiIdentifier, utils.DataInfo{
		State	: "LOGGER",
	})

	return 200, gin.H{"api_call" : "API_CALL_200"}, nil
}