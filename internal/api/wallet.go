package api

import (
	"E-Wallet-wallet/constants"
	"E-Wallet-wallet/helpers"
	"E-Wallet-wallet/internal/interfaces"
	"E-Wallet-wallet/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletAPI struct {
	WalletService interfaces.IWalletService
}

func (api *WalletAPI) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Wallet
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if req.UserID == 0 {
		log.Error("user_id is empty")
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	err := api.WalletService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("failed to create wallet: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, req)
}

func (api *WalletAPI) CreditBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.TransactionRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.CreditBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to credit balance: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *WalletAPI) DebitBalance(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.TransactionRequest
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("failed to parse request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.DebitBalanceBalance(c.Request.Context(), int(tokenData.UserID), req)
	if err != nil {
		log.Error("failed to debit balance: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *WalletAPI) GetBalance(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.GetBalance(c.Request.Context(), int(tokenData.UserID))
	if err != nil {
		log.Error("failed to get wallet: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}

func (api *WalletAPI) GetWalletHistory(c *gin.Context) {
	var (
		log   = helpers.Logger
		param models.WalletHistoryParam
	)

	if err := c.ShouldBindQuery(&param); err != nil {
		log.Error("failed to parse query param: ", err)
		helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
		return
	}

	if param.WalletTransactionType != "" {
		if param.WalletTransactionType != "CREDIT" && param.WalletTransactionType != "DEBIT" {
			log.Error("invalid wallet_transaction_type")
			helpers.SendResponse(c, http.StatusBadRequest, constants.ErrFailedBadRequest, nil)
			return
		}
	}

	token, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenData, ok := token.(models.TokenData)
	if !ok {
		log.Error("failed to parse token data")
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	resp, err := api.WalletService.GetWalletHistory(c.Request.Context(), int(tokenData.UserID), param)
	if err != nil {
		log.Error("failed to get wallet history: ", err)
		helpers.SendResponse(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponse(c, http.StatusOK, constants.SuccessMessage, resp)
}
