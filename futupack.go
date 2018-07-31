package limgo

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"math/rand"
)

// FutuPack struct
type FutuPack struct {
	szHeaderFlag  [2]uint8  // u8_t szHeaderFlag[2];
	nProtoID      uint32    // u32_t nProtoID;
	nProtoFmtType uint8     // u8_t nProtoFmtType;
	nProtoVer     uint8     // u8_t nProtoVer;
	nSerialNo     uint32    // u32_t nSerialNo;
	nBodyLen      uint32    // u32_t nBodyLen;
	arrBodySHA1   [20]uint8 // u8_t arrBodySHA1[20];
	arrReserved   [8]uint8  // u8_t arrReserved[8];
	arrBody       []byte    // []byte dataBody;
}

// SetProtoID set nProtoID
func (p *FutuPack) SetProtoID(nProtoID int) {
	p.nProtoID = uint32(nProtoID)
}

// SetBody set body
func (p *FutuPack) SetBody(arrBody []byte) {
	p.arrBody = arrBody
	p.nBodyLen = uint32(len(arrBody))

	// fmt.Println(string(p.body))
	// fmt.Println(p.nBodyLen)

	sha := sha1.New()
	sha.Write(p.arrBody)
	arrBodySHA1 := sha.Sum(nil)
	copy(p.arrBodySHA1[:], arrBodySHA1)
}

// Pack pack
func (p *FutuPack) Pack() ([]byte, error) {
	var err error

	// szHeaderFlag
	p.szHeaderFlag[0] = uint8('F')
	p.szHeaderFlag[1] = uint8('T')

	// nProtoFmtType
	p.nProtoFmtType = uint8(0)

	// nProtoVer
	p.nProtoVer = 0

	// nSerialNo
	p.nSerialNo = uint32(rand.Int31())

	packBuf := new(bytes.Buffer)
	err = binary.Write(packBuf, binary.LittleEndian, &p.szHeaderFlag)
	err = binary.Write(packBuf, binary.LittleEndian, &p.nProtoID)
	err = binary.Write(packBuf, binary.LittleEndian, &p.nProtoFmtType)
	err = binary.Write(packBuf, binary.LittleEndian, &p.nProtoVer)
	err = binary.Write(packBuf, binary.LittleEndian, &p.nSerialNo)
	err = binary.Write(packBuf, binary.LittleEndian, &p.nBodyLen)
	err = binary.Write(packBuf, binary.LittleEndian, &p.arrBodySHA1)
	err = binary.Write(packBuf, binary.LittleEndian, &p.arrReserved)

	err = binary.Write(packBuf, binary.LittleEndian, &p.arrBody)

	return packBuf.Bytes(), err
}

// Unpack unpack
func (p *FutuPack) Unpack(arrPack []byte) error {
	var err error

	reader := bytes.NewReader(arrPack)
	err = binary.Read(reader, binary.LittleEndian, &p.szHeaderFlag)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoID)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoFmtType)
	err = binary.Read(reader, binary.LittleEndian, &p.nProtoVer)
	err = binary.Read(reader, binary.LittleEndian, &p.nSerialNo)
	err = binary.Read(reader, binary.LittleEndian, &p.nBodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.arrBodySHA1)
	err = binary.Read(reader, binary.LittleEndian, &p.arrReserved)

	p.arrBody = make([]byte, p.nBodyLen)
	err = binary.Read(reader, binary.LittleEndian, &p.arrBody)

	return err
}

// to string
func (p *FutuPack) String() string {
	// return fmt.Sprintf("nBodyLen: %d body: %s",
	// 	p.nBodyLen,
	// 	p.arrBody,
	// )
	return ""
}
