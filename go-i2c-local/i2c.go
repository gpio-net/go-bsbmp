package i2c

import (
	"encoding/hex"
	"fmt"
	"os"
)

// I2C represents a connection to I2C-device.
type I2C struct {
	addr uint8
	bus  int
	rc   *os.File
}

func NewI2C(addr uint8, bus int) (*I2C, error) {
	v := &I2C{rc: nil, bus: bus, addr: addr}
	return v, nil
}

/* add functions to I2C struct to make it a class */

func (v *I2C) WriteBytes(buf []byte) (int, error) {
	fmt.Printf("Write %d hex bytes: [%+v]", len(buf), hex.EncodeToString(buf))
	//v.write(buf)
	return 0, nil
}

func (v *I2C) Close() error {
	// v.rc.Close()
	return nil
}

func (v *I2C) ReadBytes(buf []byte) (int, error) {
	return 0, nil
}

func (v *I2C) ReadRegBytes(reg byte, n int) ([]byte, int, error) {
	buf := make([]byte, n)
	return buf, 0, nil
}

func (v *I2C) ReadRegU8(reg byte) (byte, error) {
	return 0, nil
}

func (v *I2C) WriteRegU8(reg byte, value byte) error {
	return nil
}

func (v *I2C) ReadRegU16BE(reg byte) (uint16, error) {
	return 0, nil
}