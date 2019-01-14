package znp

import (
	"errors"
	"time"

	"fmt"

	"github.com/mythay/znp/codec"
	. "github.com/mythay/znp/rpc"
	"github.com/tarm/serial"
)

var ErrNoFirmware = errors.New("No firmware, only bootloader")

type Device struct {
	c       *Client
	Type    string
	Version string
}

func NewDevice(c *Client) *Device {

	return &Device{c: c}
}

func NewAcmClient(name string, baud int) *Client {
	c := &serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		return nil
	}
	return NewClient(s)
}

// StartStack try to start the zigbee stack in ZNP device. error ErrNoFirmware means there is no App firmware, only bootloader exist
func (d *Device) StartStack() (err error) {
	ver, err := SYS(d.c).Version()
	if err != nil {
		return err
	}
	if ver.Product == codec.ProductId_Schneider_Sbl { //it is in bootloader mode
		appver, err := SBL(d.c).VersionApp()
		if appver.SblStatus != 0 {
			return ErrNoFirmware
		}
		_, err = SBL(d.c).StartApp()
		if err != nil {
			return err
		}
	} else { //
		_, err := d.Reset()
		if err != nil {
			return err
		}
	}
	if ver, err = SYS(d.c).Version(); err == nil {
		switch ver.Product {
		case codec.ProductId_Schneider_Pro_GP_Znp:
			d.Type = "ZNP Pro GreenPower"
		case codec.ProductId_Schneider_Sbl:
			d.Type = "ZNP Bootloader"
		default:
			d.Type = "Unknown"
		}
		d.Version = fmt.Sprintf("%d.%d.%d", ver.MajorRel, ver.MinorRel, ver.MaintRel)
	}

	return nil
}

//Reset to reset the chip and wait for it to startup
func (d *Device) Reset() (ind *codec.ResetIndFormat, err error) {
	ind = &codec.ResetIndFormat{}
	err = SYS(d.c).ResetReq(&codec.ResetReqFormat{1})
	if err != nil {
		return
	}
	_, err = d.c.WaitAsync(ind, time.Second*3)
	return
}

type CordinatorDevice struct {
	Device
	cfg CordinatorConfig
}
type CordinatorConfig struct {
	PortName string
	PortBaud int
	Channel  int
}

func NewCordinatorDevice(cfg *CordinatorConfig) *CordinatorDevice {
	c := NewAcmClient(cfg.PortName, cfg.PortBaud)
	if c != nil {
		return &CordinatorDevice{Device: Device{c: c}, cfg: *cfg}
	}
	return nil
}
