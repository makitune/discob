package model

import (
	"io"

	"github.com/andersfylling/disgord"
	"github.com/makitune/discob/errr"
	"github.com/rylio/ytdl"
	"github.com/yyewolf/dca-disgord"
)

type Voice struct {
	Connection disgord.VoiceConnection
	Youtube    *Youtube

	encodingSession  *dca.EncodeSession
	streamingSession *dca.StreamingSession
	doneChan         chan error
}

func New(c disgord.VoiceConnection) *Voice {
	return &Voice{
		Connection: c,
	}
}

func (v *Voice) Play(y *Youtube) error {
	v.Youtube = y
	vi, err := ytdl.GetVideoInfo(y.UrlString())
	if err != nil {
		v.Stop()
		return err
	}

	format := vi.Formats.Extremes(ytdl.FormatAudioBitrateKey, true)[0]
	downloadURL, err := vi.GetDownloadURL(format)
	if err != nil {
		v.Stop()
		return err
	}

	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = dca.AudioApplicationLowDelay
	v.encodingSession, err = dca.EncodeFile(downloadURL.String(), options)
	if err != nil {
		v.Stop()
		return err
	}

	if err = v.Connection.StartSpeaking(); err != nil {
		return err
	}

	v.doneChan = make(chan error)
	v.streamingSession = dca.NewStream(v.encodingSession, v.Connection, v.doneChan)
	go func() {
		err := <-v.doneChan
		if err != nil && err != io.EOF {
			errr.Printf("%s\n", err)
		}

		v.Stop()
	}()

	return nil
}

func (v *Voice) Stop() {
	if v.streamingSession != nil {
		v.streamingSession.Stop()
	}
	if v.encodingSession != nil {
		v.encodingSession.Cleanup()
	}

	v.Connection.StopSpeaking()
	v.encodingSession = nil
	v.streamingSession = nil
	v.doneChan = nil
	v.Youtube = nil
}
