package client

import (
	"fmt"
	"time"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type UmurClientInterface interface {
	HitungUmur(tanggal_lahir time.Time) string
}

type UmurClientStruct struct {
	logger interfaces.Logger
}

func NewUmurClientStruct(logger interfaces.Logger) UmurClientInterface {
	return &UmurClientStruct{
		logger: logger,
	}
}

func (u UmurClientStruct) HitungUmur(tanggal_lahir time.Time) string {
	u.logger.Info("TANGGAL LAHIR", tanggal_lahir)
	now := time.Now()
	umur := now.Sub(tanggal_lahir)
	tahun := int64(umur.Hours() / 24 / 365)
	umur_sekarang := fmt.Sprintf("Umur sekarang adalah %d tahun", tahun)
	return umur_sekarang
}