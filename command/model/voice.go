package model

import (
	"errors"
	"os"

	"github.com/bwmarrin/dgvoice"
	"github.com/bwmarrin/discordgo"
)

type Voice struct {
	Connection *discordgo.VoiceConnection
	stopChan   chan bool
	music      *Music
}

func (v *Voice) Playing() bool {
	return v.stopChan != nil && v.music != nil
}

func (v *Voice) Play(m *Music) error {
	if m.FilePath == nil {
		return errors.New("audioFile not found")
	}

	v.music = m
	v.stopChan = make(chan bool)
	dgvoice.PlayAudioFile(v.Connection, *v.music.FilePath, v.stopChan)
	return nil
}

func (v *Voice) Stop() error {
	close(v.stopChan)
	v.stopChan = nil

	if v.music.FilePath != nil {
		if _, err := os.Stat(*v.music.FilePath); !os.IsNotExist(err) {
			if err := os.Remove(*v.music.FilePath); err != nil {
				return err
			}
		}
	}

	v.music = nil
	return nil
}
