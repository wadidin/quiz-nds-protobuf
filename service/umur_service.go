package service

import (
	"context"
	"proto-workshop/controller"
	"proto-workshop/generated/proto"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type UmurServiceStruct struct {
	logger interfaces.Logger
	ctr    controller.UmurControllerInterface
}

func NewUmurService(logger interfaces.Logger, ctr controller.UmurControllerInterface) *UmurServiceStruct {
	return &UmurServiceStruct{
		logger: logger,
		ctr:    ctr,
	}
}

func (s *UmurServiceStruct) HitungUmur(ctx context.Context, req *proto.RequestHitungUmur) (*proto.ResponseHitungUmur, error) {
	data, err := s.ctr.HitungUmur(req.GetTanggalLahir())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseHitungUmur{
		Umur: data,
	}, nil
}