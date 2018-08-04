package limgo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"limgo/futupb/GetGlobalState"
	"limgo/futupb/InitConnect"
	"limgo/futupb/KeepAlive"
	"limgo/futupb/Qot_GetSubInfo"
	"limgo/futupb/Qot_GetTicker"
	"limgo/futupb/Qot_RegQotPush"
	"limgo/futupb/Qot_Sub"
	"limgo/futupb/Qot_UpdateTicker"

	"github.com/golang/protobuf/proto"
)

// Request is a service for connect futuOpenD
type Request struct {
	Host   string
	Port   string
	FutuID uint64
	ConnID uint64
	conn   net.Conn
}

// Connect futuOpenD
func (r *Request) Connect() error {
	var err error

	r.conn, err = net.Dial("tcp", r.Host+":"+r.Port)
	if err != nil {
		return errors.New("cannot connect server")
	}

	fmt.Println("connected server: " + r.Host + ":" + r.Port)

	return nil
}

// Send data
func (r *Request) Send(pack *FutuPack) error {
	var packData []byte
	var err error

	// pack
	packData, err = pack.Pack()
	if err != nil {
		return err
	}

	// send
	r.conn.Write(packData)

	return nil
}

// Recv data
func (r *Request) Recv() {

	// scanner
	scanner := bufio.NewScanner(r.conn)
	scanner.Buffer([]byte{}, bufio.MaxScanTokenSize*10)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 < len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})

	fmt.Println("test")

	for scanner.Scan() {
		pack := new(FutuPack)
		err := pack.Unpack(scanner.Bytes())
		if err != nil {
			fmt.Println("unpack error", err)
			return
		}

		if pack.nProtoID == uint32(1001) {
			fut := &InitConnect.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}

			r.FutuID = *fut.S2C.LoginUserID
			r.ConnID = *fut.S2C.ConnID

			fmt.Println(fut.String())
		} else if pack.nProtoID == uint32(1002) {
			fut := &GetGlobalState.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(fut.String())
		} else if pack.nProtoID == uint32(1003) {

		} else if pack.nProtoID == uint32(1004) {
			fut := &KeepAlive.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			// fmt.Println(fut.String())
		} else if pack.nProtoID == uint32(3001) {
			fut := &Qot_Sub.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(pack.nProtoID, fut.GetErrCode())
		} else if pack.nProtoID == uint32(3002) {
			fut := &Qot_RegQotPush.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(pack.nProtoID, fut.String())
		} else if pack.nProtoID == uint32(3003) {
			fut := &Qot_GetSubInfo.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(pack.nProtoID, fut.String())
		} else if pack.nProtoID == uint32(3010) {
			fut := &Qot_GetTicker.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(pack.nProtoID, fut.String())

		} else if pack.nProtoID == uint32(3011) {
			fut := &Qot_UpdateTicker.Response{}
			err = proto.Unmarshal(pack.arrBody, fut)
			if err != nil {
				log.Fatal("unmarshaling error:", err)
			}
			fmt.Println(pack.nProtoID, fut.String())

		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("invalid package", err)
	}
}

// KeepAlive keep alive
func (r *Request) KeepAlive(output bool) {
	go func() {
		// pack
		pack := &FutuPack{}

		// 1004
		pack.SetProtoID(1004)
		now := time.Now().Unix()
		reqData1004 := &KeepAlive.Request{
			C2S: &KeepAlive.C2S{
				Time: &now,
			},
		}
		pbData, _ := proto.Marshal(reqData1004)
		pack.SetBody(pbData)

		tick := time.NewTicker(3 * time.Second)

		for {
			select {
			case <-tick.C:
				if output {
					// fmt.Println("tick")
				}
				// send
				r.Send(pack)
			}
		}
	}()

}

// Close connect
func (r *Request) Close() {
	r.Close()
}
