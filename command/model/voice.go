package model

import (
	"errors"
	"io"

	"github.com/bwmarrin/discordgo"
	"github.com/makitune/discob/errr"
	"github.com/rylio/ytdl"
	"github.com/yyewolf/dca-disgord"
)

type Voice struct {
	Connection *discordgo.VoiceConnection
	session    *dca.EncodeSession
	stopChan   chan struct{}
	youtube    *Youtube
}

func (v *Voice) Playing() bool {
	return v.stopChan != nil && v.youtube != nil
}

func (v *Voice) Play(y *Youtube) error {
	v.youtube = y
	vi, err := ytdl.GetVideoInfo(y.UrlString())
	if err != nil {
		v.Stop()
		return err
	}

	format := vi.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := vi.GetDownloadURL(format)
	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = dca.AudioApplicationLowDelay
	es, err := dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		v.Stop()
		return err
	}

	v.session = es
	v.stopChan = make(chan struct{})
	go func() {
		for {
			select {
			case <-v.stopChan:
				return
			default:
				data, err := es.OpusFrame()
				if err != nil {
					if err != io.EOF {
						errr.Printf("%s\n", err)
						return
					}

					if err = v.Stop(); err != nil {
						errr.Printf("%s\n", err)
						return
					}
				}

				v.Connection.OpusSend <- data
			}
		}
	}()
	return nil
}

func (v *Voice) Stop() error {
	if v.stopChan == nil {
		return errors.New("not playing")
	}

	v.stopChan <- struct{}{}
	close(v.stopChan)
	v.stopChan = nil
	v.youtube = nil

	err := v.session.Stop()
	v.session.Cleanup()
	v.session = nil
	return err
}
