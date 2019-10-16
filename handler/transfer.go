package handler

import (
	"context"
	"net/http"

	tronMsg "github.com/TRON-US/soter-demo/tron/pb/core"
	"github.com/TRON-US/soter-demo/utils"

	eth "github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/viper"
)

type Response struct {
	Code    int
	Message string
}

func Transfer(c *gin.Context) {
	address := c.PostForm("address")

	ownerAddress, err := utils.Decode58Check(viper.GetString("public_key"))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	toAddress, err := utils.Decode58Check(address)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	transactionExtension, err := utils.TronClient().TransferAsset2(context.Background(),
		&tronMsg.TransferAssetContract{
			AssetName:    []byte(viper.GetString("token_id")),
			OwnerAddress: ownerAddress,
			ToAddress:    toAddress,
			Amount:       1000000000,
		})

	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    2,
			Message: err.Error(),
		})
		return
	}

	if !transactionExtension.GetResult().GetResult() {
		c.JSON(http.StatusOK, Response{
			Code:    2,
			Message: string(transactionExtension.GetResult().GetMessage()),
		})
		return
	}

	// Marshal raw data.
	rawData, err := proto.Marshal(transactionExtension.GetTransaction().GetRawData())
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	ecdsaPK, err := eth.HexToECDSA(viper.GetString("private_key"))
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	// Sign transaction.
	sign, err := utils.Sign(rawData, ecdsaPK)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    1,
			Message: err.Error(),
		})
		return
	}

	transactionExtension.Transaction.Signature = append(transactionExtension.Transaction.Signature, sign)

	result, err := utils.TronClient().BroadcastTransaction(context.Background(), transactionExtension.Transaction)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Code:    2,
			Message: err.Error(),
		})
		return
	}

	if !result.GetResult() {
		c.JSON(http.StatusOK, Response{
			Code:    2,
			Message: string(result.GetMessage()),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "OK",
	})

}
