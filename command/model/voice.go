package model

import (
	"errors"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type Voice struct {
	Connection *discordgo.VoiceConnection
	stopChan   chan struct{}
	Youtube    *Youtube
}

func (v *Voice) Playing() bool {
	return v.stopChan != nil && v.Youtube != nil
}

func (v *Voice) Play() error {
	if v.Youtube.FilePath == nil {
		return errors.New("AudioFile Not Found")
	}
	v.stopChan = make(chan struct{})
	dgvoice.PlayAudioFile(v.Connection, *v.Youtube.FilePath, v.stopChan)
	return nil
}

func (v *Voice) Stop() error {
	close(v.stopChan)
	v.stopChan = nil

	if v.Youtube.FilePath != nil {
		if _, err := os.Stat(*v.Youtube.FilePath); !os.IsNotExist(err) {
			if err := os.Remove(*v.Youtube.FilePath); err != nil {
				return err
			}
		}
	}

	v.Youtube = nil
	return nil
}
