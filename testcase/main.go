package main

import (
	"fmt"

	"os"
	"time"

	"limgo"
	"limgo/futupb/GetGlobalState"
	"limgo/futupb/InitConnect"
	"limgo/futupb/KeepAlive"
	"limgo/futupb/Qot_Common"
	"limgo/futupb/Qot_GetSubInfo"
	"limgo/futupb/Qot_GetTicker"
	"limgo/futupb/Qot_RegQotPush"
	"limgo/futupb/Qot_Sub"

	"github.com/golang/protobuf/proto"
)

func main() {

	block := make(chan bool)

	req := &limgo.Request{Host: "127.0.0.1", Port: "11111"}
	err := req.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// pack
	pack := &limgo.FutuPack{}

	go func() {

		// for {
		// recv
		req.Recv()
		// }

	}()

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
	time.Sleep(1 * time.Second) // for test
	pack.SetProtoID(1002)
	userID := req.FutuID
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
	req.Send(pack)

	// 3001
	pack.SetProtoID(3001)
	var securityList []*Qot_Common.Security
	market := int32(11)
	code := "AAPL"
	security := &Qot_Common.Security{
		Market: &market,
		Code:   &code,
	}
	securityList = append(securityList, security)

	// fmt.Println(securityList)
	// os.Exit(1)

	var subTypeList []int32
	subType := int32(4)
	subTypeList = append(subTypeList, subType)
	isSubOrUnSub := true
	// isRegOrUnRegPush := true
	var regPushRehabTypeList []int32
	regPushRehabType := int32(1)
	regPushRehabTypeList = append(regPushRehabTypeList, regPushRehabType)

	reqData3001 := &Qot_Sub.Request{
		C2S: &Qot_Sub.C2S{
			SecurityList:         securityList,
			SubTypeList:          subTypeList,
			IsSubOrUnSub:         &isSubOrUnSub,
			RegPushRehabTypeList: regPushRehabTypeList,
		},
	}
	// fmt.Println(reqData3001)
	// os.Exit(1)
	pbData, err = proto.Marshal(reqData3001)
	if err != nil {
		fmt.Println(3001, err)
	}
	pack.SetBody(pbData)
	req.Send(pack)

	// 3002
	pack.SetProtoID(3002)
	isRegOrUnReg := true
	reqData3002 := &Qot_RegQotPush.Request{
		C2S: &Qot_RegQotPush.C2S{
			SecurityList:  securityList,
			SubTypeList:   subTypeList,
			RehabTypeList: regPushRehabTypeList,
			IsRegOrUnReg:  &isRegOrUnReg,
		},
	}
	pbData, _ = proto.Marshal(reqData3002)
	pack.SetBody(pbData)
	req.Send(pack)

	// 3003
	pack.SetProtoID(3003)
	isReqAllConn := true
	reqData3003 := &Qot_GetSubInfo.Request{
		C2S: &Qot_GetSubInfo.C2S{
			IsReqAllConn: &isReqAllConn,
		},
	}
	pbData, _ = proto.Marshal(reqData3003)
	pack.SetBody(pbData)
	req.Send(pack)

	// 3010
	pack.SetProtoID(3010)
	maxRetNum := int32(10)
	reqData3010 := &Qot_GetTicker.Request{
		C2S: &Qot_GetTicker.C2S{
			Security:  security,
			MaxRetNum: &maxRetNum,
		},
	}
	pbData, _ = proto.Marshal(reqData3010)
	pack.SetBody(pbData)
	// send
	req.Send(pack)

	req.KeepAlive(true)

	<-block

}
