package client

import (
	"math"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ClientInterFace interface {
	HitungLuasPersegi(panjang int32, lebar int32) int32
	HitungKelilingPersegi(panjang int32, lebar int32) int32
	HitungLuasLingkaran(jarijari int32) float64
	HitungKelilingLingkaran(jarijari float64) float64
	HitungLuasSegitiga(alas int32, tinggi int32) float64
	HitungKelilingSegitiga(sisi1 int32, sisi2 int32, sisi3 int32) int32
	HitungVolumeKubus(sisi int32) int32
}

type ClientStruct struct {
	logger interfaces.Logger
}

func NewClientStruct(logger interfaces.Logger) ClientInterFace {
	return &ClientStruct{
		logger: logger,
	}
}

//Luas Persegi Panjang
func (c *ClientStruct) HitungLuasPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := panjang * lebar
	c.logger.Info("LUAS", luas)
	return luas
}

//Keliling Persegi Panjang
func (c *ClientStruct) HitungKelilingPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	keliling := 2 * (panjang + lebar)
	c.logger.Info("LUAS", keliling)
	return keliling
}

func (c *ClientStruct) HitungLuasLingkaran(jarijari int32) float64 {
	c.logger.Info("JARI-JARI", jarijari)
	luas := math.Phi * math.Pow(float64(jarijari), 2)
	c.logger.Info("LUAS", luas)
	return luas
}

func (c *ClientStruct) HitungKelilingLingkaran(jarijari float64) float64 {
	c.logger.Info("JARI-JARI", jarijari)
	keliling := 2 * math.Phi * jarijari
	c.logger.Info("KELILING", keliling)
	return keliling
}

func (c *ClientStruct) HitungLuasSegitiga(alas int32, tinggi int32) float64 {
	c.logger.Info("ALAS", alas)
	c.logger.Info("TINGGI", tinggi)
	luas := (alas * tinggi) / 2
	c.logger.Info("LUAS", luas)
	return float64(luas)
}

func (c *ClientStruct) HitungKelilingSegitiga(sisi1 int32, sisi2 int32, sisi3 int32) int32 {
	c.logger.Info("SISI 1", sisi1)
	c.logger.Info("SISI 2", sisi2)
	c.logger.Info("SISI 3", sisi3)
	keliling := sisi1 + sisi2 + sisi3
	c.logger.Info("KELILING", keliling)
	return keliling
}

func (c *ClientStruct) HitungVolumeKubus(sisi int32) int32 {
	c.logger.Info("SISI", sisi)
	volume := sisi * sisi * sisi
	c.logger.Info("VOLUME", volume)
	return volume
}