package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/AiLiaa/tag-service/pkg/bapi"
	"github.com/AiLiaa/tag-service/pkg/errcode"
	pb "github.com/AiLiaa/tag-service/proto"
)

type TagServer struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		err := errcode.TogRPCError(errcode.ErrorGetTagListFail)
		sts := errcode.FromError(err)
		details := sts.Details()
		fmt.Println(err)
		fmt.Println(details) //告诉内部客户端 业务错误类型
		//告诉服务端
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
		return nil, err
	}

	return &tagList, nil
}
