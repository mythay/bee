package znp

import (
	"testing"

	. "github.com/mythay/znp"
	. "github.com/onsi/gomega"
)

func TestNewAcmDevice(t *testing.T) {
	RegisterTestingT(t)
	t.Run("USB ACM connection", func(t *testing.T) {
		device := NewCordinatorDevice(&CordinatorConfig{PortName: "COM11", PortBaud: 115200, Channel: 11})
		device.StartStack()
	})
}
