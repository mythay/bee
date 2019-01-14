package codec

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_getCmdFromRequest(t *testing.T) {
	RegisterTestingT(t)
	t.Run("the request is an address", func(t *testing.T) {
		c, _ := GetCmdFromRequest(&ResetReqFormat{})

		Expect(c).To(Equal(Cmd{MT_RPC_SYS_SYS | MT_RPC_CMD_AREQ, MT_SYS_RESET_REQ}))

	})

	t.Run("gthe request is a struct", func(t *testing.T) {
		c, _ := GetCmdFromRequest(ResetReqFormat{})
		Expect(c).To(Equal(Cmd{MT_RPC_SYS_SYS | MT_RPC_CMD_AREQ, MT_SYS_RESET_REQ}))

	})

}
