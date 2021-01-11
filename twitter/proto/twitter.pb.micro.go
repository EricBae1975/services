// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/twitter.proto

package twitter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Api service

func NewApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Api service

type ApiService interface {
	Tweet(ctx context.Context, in *TweetRequest, opts ...client.CallOption) (*TweetResponse, error)
}

type apiService struct {
	c    client.Client
	name string
}

func NewApiService(name string, c client.Client) ApiService {
	return &apiService{
		c:    c,
		name: name,
	}
}

func (c *apiService) Tweet(ctx context.Context, in *TweetRequest, opts ...client.CallOption) (*TweetResponse, error) {
	req := c.c.NewRequest(c.name, "Api.Tweet", in)
	out := new(TweetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiHandler interface {
	Tweet(context.Context, *TweetRequest, *TweetResponse) error
}

func RegisterApiHandler(s server.Server, hdlr ApiHandler, opts ...server.HandlerOption) error {
	type api interface {
		Tweet(ctx context.Context, in *TweetRequest, out *TweetResponse) error
	}
	type Api struct {
		api
	}
	h := &apiHandler{hdlr}
	return s.Handle(s.NewHandler(&Api{h}, opts...))
}

type apiHandler struct {
	ApiHandler
}

func (h *apiHandler) Tweet(ctx context.Context, in *TweetRequest, out *TweetResponse) error {
	return h.ApiHandler.Tweet(ctx, in, out)
}