package controller

import (
	"proto-workshop/client"
	"time"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UmurControllerInterface interface {
	HitungUmur(tanggal_lahir string) (string, error)
}

type UmurControllerStruct struct {
	logger interfaces.Logger
	uc     client.UmurClientInterface
}

func NewUmurController(logger interfaces.Logger, uc client.UmurClientInterface) UmurControllerInterface {
	return &UmurControllerStruct{
		logger: logger,
		uc:     uc,
	}
}

func (c *UmurControllerStruct) HitungUmur(tanggal_lahir string) (string, error) {
	date, error := time.Parse("2006-01-02", tanggal_lahir)
	if error != nil {
		return "", status.Error(codes.FailedPrecondition, "Format tanggal tidak valid")
	}
	umur := c.uc.HitungUmur(date)
	return umur, nil
}