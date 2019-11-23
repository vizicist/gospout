package gospout

import (
	"github.com/vizicist/gospout/internal/cspout"
)

// Config returns a config string
func Config() string {
	return "gospout config"
}

// Sender is a spout sender
type Sender struct {
	sender cspout.Sender
}

// CreateSender returns a Sender
func CreateSender(name string, width int, height int) Sender {
	var s Sender
	s.sender = cspout.CreateSender(name, width, height)
	return s
}

// SendTexture sends a texture
func SendTexture(s Sender, texture uint32, width int, height int) bool {
	return cspout.SendTexture(s.sender, texture, width, height)
}
