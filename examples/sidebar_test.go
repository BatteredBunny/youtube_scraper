package examples

import (
	"git.catnip.ee/miisu/youtube_scraper"
	"testing"
)

func TestSidebarVideos(t *testing.T) {
	v, err := scraper.NewVideoScraper("nfCUTZWwlvo")
	if err != nil {
		t.Fatal(err)
	}

	HandleSidebarEntries(t, v.InitialSidebarEntries)

	var sidebarEntries []scraper.SidebarEntry
	for i := 0; i <= 3; i++ {
		sidebarEntries, err = v.NextSidebarVideosPage()
		if err != nil {
			t.Fatal(err)
		} else if len(sidebarEntries) == 0 {
			break
		}

		HandleSidebarEntries(t, sidebarEntries)
	}
}

func HandleSidebarEntries(t *testing.T, sidebarEntries []scraper.SidebarEntry) {
	for _, sidebarEntry := range sidebarEntries {
		switch sidebarEntry.Type {
		case scraper.SidebarEntryVideo:
			sidebarVideo := sidebarEntry.Entry.(scraper.SidebarVideo)
			t.Log("video:", sidebarVideo)
		case scraper.SidebarEntryPlaylist:
			sidebarPlaylist := sidebarEntry.Entry.(scraper.SidebarPlaylist)
			t.Log("playlist:", sidebarPlaylist)
		case scraper.SidebarEntryRadio:
			sidebarRadio := sidebarEntry.Entry.(scraper.SidebarRadio)
			t.Log("radio:", sidebarRadio)
		}
	}
}
