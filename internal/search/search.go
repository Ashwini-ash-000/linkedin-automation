package search

import "github.com/go-rod/rod"

type Profile struct {
	ID string
}

func CollectProfiles(page *rod.Page) ([]Profile, error) {
	cards := page.MustElements(".card")

	var profiles []Profile
	for _, card := range cards {
		id, err := card.Attribute("data-profile-id")
		if err != nil || id == nil {
			continue
		}
		profiles = append(profiles, Profile{ID: *id})
	}
	return profiles, nil
}

func HasNextPage(page *rod.Page) bool {
	return page.MustHas("#next")
}

func GoToNextPage(page *rod.Page) {
	page.MustElement("#next").MustClick()
}
