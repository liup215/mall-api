package client

import (
	"context"
	"mall/app/api/web/member/api"
	"time"

	"google.golang.org/grpc"
)

func (c *Client) MemberCreate(uniacid int, mobile, pwd string) (*api.MemberCreateReply, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	defer conn.Close()
	c := api.NewMemberServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return c.MemberCreateMobile(ctx, &api.MemberInfoByMobileRequest{
		Uniacid: uniacid,
		Mobile:  mobile,
	})

}
