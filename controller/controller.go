package controller

import (
	"proto-workshop/client"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControllerInterface interface {
	HitungLuasPersegi(panjang int32, lebar int32) (int32, error)
	HitungKelilingPersegi(panjang int32, lebar int32) (int32, error)
	HitungLuasLingkaran(jarijari int32) (float64, error)
	HitungKelilingLingkaran(jarijari float64) (float64, error)
	HitungLuasSegitiga(alas int32, tinggi int32) (float64, error)
	HitungKelilingSegitiga(sisi1 int32, sisi2 int32, sisi3 int32) (int32, error)
	HitungVolumeKubus(sisi int32) (int32, error)
}

type ControllerStruct struct {
	logger interfaces.Logger
	cl     client.ClientInterFace
}

func NewController(logger interfaces.Logger, cl client.ClientInterFace) ControllerInterface {
	return &ControllerStruct{
		logger: logger,
		cl:     cl,
	}
}

func (c *ControllerStruct) HitungLuasPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := c.cl.HitungLuasPersegi(panjang, lebar)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungKelilingPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	keliling := c.cl.HitungKelilingPersegi(panjang, lebar)
	c.logger.Info("LUAS", keliling)
	return keliling, nil
}

func (c *ControllerStruct) HitungLuasLingkaran(jarijari int32) (float64, error) {
	if jarijari < 1 {
		c.logger.Error("Error", "Jari-jari Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Jari-jari harus Lebih Dari 0")
	}
	c.logger.Info("JARI-JARI", jarijari)
	luas := c.cl.HitungLuasLingkaran(jarijari)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungKelilingLingkaran(jarijari float64) (float64, error) {
	if jarijari < 1 {
		c.logger.Error("Error", "Jari-jari Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Jari-jari harus Lebih Dari 0")
	}
	c.logger.Info("JARI-JARI", jarijari)
	keliling := c.cl.HitungKelilingLingkaran(jarijari)
	c.logger.Info("LUAS", keliling)
	return keliling, nil
}

func (c *ControllerStruct) HitungLuasSegitiga(alas int32, tinggi int32) (float64, error) {
	if alas < 1 {
		c.logger.Error("Error", "Alas Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Alas harus Lebih Dari 0")
	}
	if tinggi < 1 {
		c.logger.Error("Error", "Tinggi Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Tinggi harus Lebih Dari 0")
	}
	c.logger.Info("ALAS", alas)
	c.logger.Info("TINGGI", tinggi)
	luas := c.cl.HitungLuasSegitiga(alas, tinggi)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungKelilingSegitiga(sisi1 int32, sisi2 int32, sisi3 int32) (int32, error) {
	if sisi1 < 1 {
		c.logger.Error("Error", "Sisi 1 Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi 1 harus Lebih Dari 0")
	}
	if sisi2 < 1 {
		c.logger.Error("Error", "Sisi 2 Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi 2 harus Lebih Dari 0")
	}
	if sisi3 < 1 {
		c.logger.Error("Error", "Sisi 3 Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi 3 harus Lebih Dari 0")
	}
	c.logger.Info("SISI 1", sisi1)
	c.logger.Info("SISI 2", sisi2)
	c.logger.Info("SISI 3", sisi3)
	keliling := c.cl.HitungKelilingSegitiga(sisi1, sisi2, sisi3)
	c.logger.Info("KELILING", keliling)
	return keliling, nil
}

func (c *ControllerStruct) HitungVolumeKubus(sisi int32) (int32, error) {
	if sisi < 1 {
		c.logger.Error("Error", "Sisi Harus Lebih Dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi harus Lebih Dari 0")
	}
	c.logger.Info("SISI", sisi)
	volume := c.cl.HitungVolumeKubus(sisi)
	c.logger.Info("VOLUME", volume)
	return volume, nil
}