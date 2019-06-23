package codec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getCmdFromRequest(t *testing.T) {
	assert := assert.New(t)
	t.Run("the request is an address", func(t *testing.T) {

		c, _ := GetCmdFromRequest(&ResetReqFormat{})

		assert.Equal(c, Cmd{MT_RPC_SYS_SYS | MT_RPC_CMD_AREQ, MT_SYS_RESET_REQ})
		// Expect(ResetReqFormat).To(Equal(reflect.TypeOf(ResetReqFormat{})))

	})

	t.Run("the request is a struct", func(t *testing.T) {
		c, _ := GetCmdFromRequest(ResetReqFormat{})
		assert.Equal(c, Cmd{MT_RPC_SYS_SYS | MT_RPC_CMD_AREQ, MT_SYS_RESET_REQ})
	})

}
