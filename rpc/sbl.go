package rpc

import (
	"errors"
	"time"

	. "github.com/mythay/bee/codec"
)

type Sbl struct {
	*Client
}

func SBL(c *Client) *Sbl {
	return &Sbl{c}
}
func (s *Sbl) VersionApp() (ind *SblAppVersionIndFormat, err error) {
	ind = &SblAppVersionIndFormat{}
	_, err = s.Call(NewCmd(MT_RPC_SYS_SBL, MT_SBL_VERSION_APP_REQ))
	if err != nil {
		return
	}
	_, err = s.WaitAsync(ind, time.Second*3)
	return
}

func (s *Sbl) StartApp() (ind *ResetIndFormat, err error) {
	ind = &ResetIndFormat{}
	_, err = s.Call(NewCmd(MT_RPC_SYS_SBL, MT_SBL_START_APP_REQ))
	if err != nil {
		return
	}
	status := &SblStartAppIndFormat{}

	_, err = s.WaitAsync(status, time.Second*3)
	if err != nil {
		return
	}
	if status.SblStatus != 0 {
		err = errors.New("start fail")
		return
	}
	_, err = s.WaitAsync(ind, time.Second*3)
	return
}
