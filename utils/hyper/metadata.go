package hyper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"main/utils/task"
)

func SaveTrackInfoJSON(track *task.Track) error {
	// Extract year from ReleaseDate
	year := ""
	if len(track.Resp.Attributes.ReleaseDate) >= 4 {
		year = track.Resp.Attributes.ReleaseDate[:4]
	}

	rackInfo := map[string]string{
		"title":           track.Resp.Attributes.Name,
		"artist":          track.Resp.Attributes.ArtistName,
		"album":           track.Resp.Attributes.AlbumName,
		"year":            year,
		"trackNumber":     strconv.Itoa(track.Resp.Attributes.TrackNumber),
		"discNumber":      strconv.Itoa(track.Resp.Attributes.DiscNumber),
		"downloadUrl":     track.OriginalUrl,
		"isrc":            track.Resp.Attributes.Isrc,
		"providerTrackId": track.ID,
		"releaseDate":     track.Resp.Attributes.ReleaseDate,
		"contentRating":   track.Resp.Attributes.ContentRating,
		"artworkUrl":      track.Resp.Attributes.Artwork.URL,
	}

	// Get the directory of the track file
	dir := track.SaveDir
	infoPath := filepath.Join(dir, "info.json")

	// Marshal to JSON
	data, err := json.MarshalIndent(rackInfo, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	// Write to file
	err = os.WriteFile(infoPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write info.json: %v", err)
	}

	return nil
}
