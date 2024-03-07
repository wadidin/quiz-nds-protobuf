package service

import (
	"context"
	"proto-workshop/controller"
	"proto-workshop/generated/proto"

	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
)

type ServiceStruct struct {
	logger interfaces.Logger
	ctr    controller.ControllerInterface
}

func NewService(logger interfaces.Logger, ctr controller.ControllerInterface) *ServiceStruct {
	return &ServiceStruct{
		logger: logger,
		ctr:    ctr,
	}
}

func (s *ServiceStruct) HitungLuasPersegi(ctx context.Context, req *proto.RequestHitungLuasPersegi) (*proto.ResponseLuasPersegi, error) {
	data, err := s.ctr.HitungLuasPersegi(req.GetPanjang(), req.GetLebar())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseLuasPersegi{
		Luas: data,
	}, nil
}

func (s *ServiceStruct) HitungKelilingPersegi(ctx context.Context, req *proto.RequestHitungKelilingPersegi) (*proto.ResponseKelilingPersegi, error) {
	data, err := s.ctr.HitungKelilingPersegi(req.GetPanjang(), req.GetLebar())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseKelilingPersegi{
		Keliling: data,
	}, nil
}

func (s *ServiceStruct) HitungLuasLingkaran(ctx context.Context, req *proto.RequestHitungLuasLingkaran) (*proto.ResponseLuasLingkaran, error) {
	data, err := s.ctr.HitungLuasLingkaran(req.GetJarijari())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseLuasLingkaran{
		Luas: data,
	}, nil
}

func (s *ServiceStruct) HitungKelilingLingkaran(ctx context.Context, req *proto.RequestHitungKelilingLingkaran) (*proto.ResponseKelilingLingkaran, error) {
	data, err := s.ctr.HitungKelilingLingkaran(req.GetJarijari())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseKelilingLingkaran{
		Keliling: data,
	}, nil
}

func (s *ServiceStruct) HitungLuasSegitiga(ctx context.Context, req *proto.RequestHitungLuasSegitiga) (*proto.ResponseLuasSegitiga, error) {
	data, err := s.ctr.HitungLuasSegitiga(req.GetAlas(), req.GetTinggi())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseLuasSegitiga{
		Luas: data,
	}, nil
}

func (s *ServiceStruct) HitungKelilingSegitiga(ctx context.Context, req *proto.RequestHitungKelilingSegitiga) (*proto.ResponseKelilingSegitiga, error) {
	data, err := s.ctr.HitungKelilingSegitiga(req.GetSisi1(), req.GetSisi2(), req.GetSisi3())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseKelilingSegitiga{
		Keliling: data,
	}, nil
}

func (s *ServiceStruct) HitungVolumeKubus(ctx context.Context, req *proto.RequestKubus) (*proto.ResponseVolumeKubus, error) {
	data, err := s.ctr.HitungVolumeKubus(req.GetSisi())
	if err != nil {
		return nil, err
	}
	return &proto.ResponseVolumeKubus{
		Volume: data,
	}, nil
}