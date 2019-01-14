package rpc

import (
	. "github.com/mythay/znp/codec"
)

type Sys struct {
	*Client
}

func SYS(c *Client) *Sys {
	return &Sys{c}
}

func (s *Sys) Ping() (*PingSrspFormat, error) {

	resp, err := s.Call(NewCmd(MT_RPC_SYS_SYS, MT_SYS_PING))
	if err != nil {
		return nil, err
	}
	return resp.(*PingSrspFormat), err
}

func (s *Sys) SetExtAddr(req *SetExtAddrFormat) error {
	_, err := s.Call(req)

	return err
}
func (s *Sys) GetExtAddr() (*GetExtAddrSrspFormat, error) {
	resp, err := s.Call(NewCmd(MT_RPC_SYS_SYS, MT_SYS_GET_EXTADDR))
	if err != nil {
		return nil, err
	}
	return resp.(*GetExtAddrSrspFormat), err
}

func (s *Sys) RamRead(req *RamReadFormat) (*RamReadSrspFormat, error) {
	resp, err := s.Call(req)
	if err != nil {
		return nil, err
	}
	return resp.(*RamReadSrspFormat), err
}
func (s *Sys) RamWrite(req *RamWriteFormat) error {
	_, err := s.Call(req)

	return err
}
func (s *Sys) ResetReq(req *ResetReqFormat) error {
	_, err := s.Call(req)

	return err
}

func (s *Sys) Version() (*VersionSrspFormat, error) {
	resp, err := s.Call(NewCmd(MT_RPC_SYS_SYS, MT_SYS_VERSION))
	if err != nil {
		return nil, err
	}
	return resp.(*VersionSrspFormat), err
}

func (s *Sys) OsalNvRead(req *OsalNvReadFormat) (*OsalNvReadSrspFormat, error) {
	resp, err := s.Call(req)
	if err != nil {
		return nil, err
	}
	return resp.(*OsalNvReadSrspFormat), err
}

func (s *Sys) OsalNvWrite(req *OsalNvWriteFormat) error {
	_, err := s.Call(req)

	return err
}

func (s *Sys) SetTxPower(req *SetTxPowerFormat) (*SetTxPowerSrspFormat, error) {
	resp, err := s.Call(req)
	if err != nil {
		return nil, err
	}
	return resp.(*SetTxPowerSrspFormat), err
}

func (s *Sys) GetAntennaMode() (*GetAntennaModeSrspFormat, error) {
	resp, err := s.Call(NewCmd(MT_RPC_SYS_SYS, MT_SYS_GET_ANTENNA_MODE))
	if err != nil {
		return nil, err
	}
	return resp.(*GetAntennaModeSrspFormat), err
}
