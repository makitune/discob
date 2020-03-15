package search

import "github.com/makitune/discob/command/model"

func newYoutube(resp *youtubeSearchResponse) *model.Youtube {
	item := &resp.Items[0]
	return &model.Youtube{
		Title:       item.Snippet.Title,
		Description: item.Snippet.Description,
		VideoID:     item.ID.VideoID,
		FilePath:    nil,
	}
}
