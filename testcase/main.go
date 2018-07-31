package main

import (
	"fmt"
	"limgo"
	"limgo/futupb/GetGlobalState"
	"limgo/futupb/InitConnect"
	"limgo/futupb/KeepAlive"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
)

func main() {

	req := &limgo.Request{Host: "127.0.0.1", Port: "11111"}
	err := req.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// pack
	pack := &limgo.FutuPack{}

	// 1001
	pack.SetProtoID(1001)
	ClientVer := int32(307)
	ClientID := "test123456"
	reqData1001 := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  &ClientID,
			ClientVer: &ClientVer,
		},
	}
	pbData, _ := proto.Marshal(reqData1001)
	pack.SetBody(pbData)
	// send
	req.Send(pack)

	// 1002
	pack.SetProtoID(1002)
	userID := uint64(req.FutuId)
	reqData1002 := &GetGlobalState.Request{
		C2S: &GetGlobalState.C2S{
			UserID: &userID,
		},
	}
	pbData, _ = proto.Marshal(reqData1002)
	pack.SetBody(pbData)
	// send
	req.Send(pack)

	// 1004
	pack.SetProtoID(1004)
	time := time.Now().Unix()
	reqData1004 := &KeepAlive.Request{
		C2S: &KeepAlive.C2S{
			Time: &time,
		},
	}
	pbData, _ = proto.Marshal(reqData1004)
	pack.SetBody(pbData)
	// send
	req.Send(pack)

	req.KeepAlive(true)

	// recv
	req.Recv()

}
