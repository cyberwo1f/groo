package protocol_test

import (
	"github.com/cyberfox1002/groo/protocol"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProtocol(t *testing.T) {
	var prt protocol.Protocol

	// return protocol https
	prt = protocol.GetProtocol("https://hogehoge.com")
	assert.Equal(t, prt, protocol.Https)

	// return protocol ssh
	prt = protocol.GetProtocol("git@hoge.com:aaaa")
	assert.Equal(t, prt, protocol.SSH)

	// return protocol unknown
	prt = protocol.GetProtocol("hogefuga")
	assert.Equal(t, prt, protocol.Unknown)
}
