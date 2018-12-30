package model

import (
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type Voice struct {
	Connection *discordgo.VoiceConnection
	stopChan   chan bool
	Youtube    *Youtube
}

func (v *Voice) Playing() bool {
	return v.stopChan != nil && v.Youtube != nil
}

func (v *Voice) Play() {
	v.stopChan = make(chan bool)
	dgvoice.PlayAudioFile(v.Connection, v.Youtube.FilePath, v.stopChan)
}

func (v *Voice) Stop() error {
	close(v.stopChan)
	v.stopChan = nil

	if _, err := os.Stat(v.Youtube.FilePath); err != nil {
		return err
	}

	if err := os.Remove(v.Youtube.FilePath); err != nil {
		return err
	}

	v.Youtube = nil
	return nil
}
