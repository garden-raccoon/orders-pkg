package models

import (
	proto "github.com/garden-raccoon/orders-pkg/protocols/orders-pkg"
	"github.com/gofrs/uuid"
	"time"
)

func ParamToProto(p *Param) *proto.Param {
	return &proto.Param{
		MealUuid: p.MealUuid.Bytes(),
		Date:     p.Date.Unix(),
		Quantity: int64(p.Quantity),
	}
}

func ParamsToProto(p []*Param) *proto.Params {
	protoParams := &proto.Params{}
	for i := range p {
		protoParams.Params = append(protoParams.Params, ParamToProto(p[i]))
	}
	return protoParams
}

func ParamFromProto(pb *proto.Param) *Param {
	return &Param{
		MealUuid: uuid.FromBytesOrNil(pb.MealUuid),
		Quantity: int(pb.Quantity),
		Date:     time.Unix(pb.Date, 0),
	}
}

func ParamsFromProto(pb *proto.Params) []*Param {
	var params []*Param
	for i := range pb.Params {
		params = append(params, ParamFromProto(pb.Params[i]))
	}
	return params
}
