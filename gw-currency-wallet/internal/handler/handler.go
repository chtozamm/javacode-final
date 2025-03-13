package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	pb "github.com/chtozamm/javacode-final/proto-exchange/exchange"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Register(r *gin.Engine) {
	{
		v1 := r.Group("/api/v1")

		// v1.POST("/register", nil)
		// v1.POST("/login", nil)

		// v1.GET("/balance", nil)
		// v1.POST("/wallet/deposit", nil)
		// v1.POST("/wallet/withdraw", nil)

		v1.GET("/exchange/rates", handleGetExchangeRates)
		v1.POST("/exchange", handleGetExchangeRateForCurrency)
	}
}

func handleGetExchangeRates(c *gin.Context) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// TODO: handle error
		return
	}
	defer conn.Close()

	client := pb.NewExchangeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	rates, err := client.GetExchangeRates(ctx, &pb.Empty{})
	if err != nil {
		// TODO: handle error
		return
	}
	fmt.Println(rates.Rates)

	c.JSON(http.StatusOK, gin.H{
		"rates": rates.Rates,
	})
}

func handleGetExchangeRateForCurrency(c *gin.Context) {
	type RequestBody struct {
		FromCurrency string  `json:"from_currency"`
		ToCurrency   string  `json:"to_currency"`
		Amount       float32 `json:"amount"`
	}

	var reqBody RequestBody
	err := json.NewDecoder(c.Request.Body).Decode(&reqBody)
	if err != nil {
		// TODO: handle error
		return
	}
	if reqBody.Amount <= 0 {
		// TODO: handle error
		return
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// TODO: handle error
		return
	}
	defer conn.Close()

	client := pb.NewExchangeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	rate, err := client.GetExchangeRateForCurrency(ctx, &pb.CurrencyRequest{
		FromCurrency: reqBody.FromCurrency,
		ToCurrency:   reqBody.ToCurrency,
	})
	if err != nil {
		// TODO: handle error
		return
	}

	calculated := rate.Rate * reqBody.Amount

	c.JSON(http.StatusOK, gin.H{
		"message":          "Exchange successful",
		"exchanged_amount": calculated,
		// TODO: return new_balance
		// "new_balance":      gin.H{},
	})
}
