// Code generated by goctl. DO NOT EDIT!
// Source: mag.proto

package magclient

import (
	"context"

	"mag_service/rpc/mag"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Abstract    = mag.Abstract
	CommonResp  = mag.CommonResp
	NlpTags     = mag.NlpTags
	ReqAbsId    = mag.ReqAbsId
	ReqKeyWord  = mag.ReqKeyWord
	Request     = mag.Request
	RespAbsMore = mag.RespAbsMore
	Response    = mag.Response

	Mag interface {
		CreateDocument(ctx context.Context, in *Abstract, opts ...grpc.CallOption) (*CommonResp, error)
		UpdateDocument(ctx context.Context, in *Abstract, opts ...grpc.CallOption) (*CommonResp, error)
		GetDocumentById(ctx context.Context, in *ReqAbsId, opts ...grpc.CallOption) (*Abstract, error)
		SearchDocumentByKey(ctx context.Context, in *ReqKeyWord, opts ...grpc.CallOption) (*RespAbsMore, error)
		GetNlpById(ctx context.Context, in *ReqAbsId, opts ...grpc.CallOption) (*NlpTags, error)
	}

	defaultMag struct {
		cli zrpc.Client
	}
)

func NewMag(cli zrpc.Client) Mag {
	return &defaultMag{
		cli: cli,
	}
}

func (m *defaultMag) CreateDocument(ctx context.Context, in *Abstract, opts ...grpc.CallOption) (*CommonResp, error) {
	client := mag.NewMagClient(m.cli.Conn())
	return client.CreateDocument(ctx, in, opts...)
}

func (m *defaultMag) UpdateDocument(ctx context.Context, in *Abstract, opts ...grpc.CallOption) (*CommonResp, error) {
	client := mag.NewMagClient(m.cli.Conn())
	return client.UpdateDocument(ctx, in, opts...)
}

func (m *defaultMag) GetDocumentById(ctx context.Context, in *ReqAbsId, opts ...grpc.CallOption) (*Abstract, error) {
	client := mag.NewMagClient(m.cli.Conn())
	return client.GetDocumentById(ctx, in, opts...)
}

func (m *defaultMag) SearchDocumentByKey(ctx context.Context, in *ReqKeyWord, opts ...grpc.CallOption) (*RespAbsMore, error) {
	client := mag.NewMagClient(m.cli.Conn())
	return client.SearchDocumentByKey(ctx, in, opts...)
}

func (m *defaultMag) GetNlpById(ctx context.Context, in *ReqAbsId, opts ...grpc.CallOption) (*NlpTags, error) {
	client := mag.NewMagClient(m.cli.Conn())
	return client.GetNlpById(ctx, in, opts...)
}
