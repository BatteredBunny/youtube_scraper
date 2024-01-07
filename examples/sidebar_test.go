package examples

import (
	"testing"

	scraper "github.com/BatteredBunny/youtube_scraper/video"
)

func TestSidebarVideos(t *testing.T) {
	v, err := scraper.NewVideoScraper("n8Rp_5lvhFI")
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
		case scraper.SidebarEntryTypeVideo:
			sidebarVideo := sidebarEntry.Entry.(scraper.SidebarVideo)
			t.Log("video:", sidebarVideo)
		case scraper.SidebarEntryTypePlaylist:
			sidebarPlaylist := sidebarEntry.Entry.(scraper.SidebarPlaylist)
			t.Log("playlist:", sidebarPlaylist)
		case scraper.SidebarEntryTypeRadio:
			sidebarRadio := sidebarEntry.Entry.(scraper.SidebarRadio)
			t.Log("radio:", sidebarRadio)
		}
	}
	t.Log("-------------")
}
