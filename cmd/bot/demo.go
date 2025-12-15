package main

import (
	"path/filepath"
	"time"

	"github.com/yourusername/linkedin-automation/internal/browser"
	"github.com/yourusername/linkedin-automation/internal/messaging"
	"github.com/yourusername/linkedin-automation/internal/search"
	"github.com/yourusername/linkedin-automation/internal/stealth"
	"github.com/yourusername/linkedin-automation/internal/store"
)

func runBot(
	session *browser.Session,
	st *store.Store,
	log interface{},
) error {

	// ---- Create page ----
	page := session.Browser.MustPage()
	page.MustNavigate("about:blank")
	stealth.Apply(page)

	// ---- Navigate to mock LinkedIn ----
	path, _ := filepath.Abs("mock/linkedin.html")
	page.MustNavigate("file:///" + path)

	log.(interface{ Info(string) }).Info("Opened mock LinkedIn page")
	stealth.LongPause()

	// ===============================
	// SEARCH + CONNECTION REQUESTS
	// ===============================
	profiles, err := search.CollectProfiles(page)
	if err != nil {
		return err
	}

	for _, p := range profiles {

		// Skip duplicates
		if sent, _ := st.HasSent(p.ID); sent {
			log.(interface{ Info(string) }).Info("Skipping profile: " + p.ID)
			continue
		}

		// Daily limit
		count, _ := st.CountSentToday()
		if count >= 15 {
			log.(interface{ Info(string) }).Info("Daily limit reached")
			return nil
		}

		card := page.MustElement(`[data-profile-id="` + p.ID + `"]`)
		btn := card.MustElement(".connect")

		box := btn.MustShape().Box()
		from := stealth.Point{X: box.X - 50, Y: box.Y}
		to := stealth.Point{X: box.X + box.Width/2, Y: box.Y + box.Height/2}

		stealth.MoveMouseBezier(page, from, to)
		time.Sleep(300 * time.Millisecond)
		btn.MustClick()

		log.(interface{ Info(string) }).Info("Sent request to " + p.ID)
		_ = st.MarkSent(p.ID)

		stealth.LongPause()
	}

	// ===============================
	// MESSAGING SYSTEM
	// ===============================
	cards := page.MustElements(".card")

	for _, card := range cards {

		accepted, _ := card.Attribute("data-accepted")
		if accepted == nil || *accepted != "true" {
			continue
		}

		id, _ := card.Attribute("data-profile-id")
		name, _ := card.Attribute("data-name")
		if id == nil || name == nil {
			continue
		}

		if done, _ := st.HasMessaged(*id); done {
			log.(interface{ Info(string) }).Info("Already messaged " + *id)
			continue
		}

		msgBtn := card.MustElement(".message")
		msgBox := card.MustElement(".msgbox")

		template := "Hi {{name}}, thanks for connecting! Looking forward to staying in touch."
		message := messaging.Render(template, map[string]string{
			"name": *name,
		})

		msgBtn.MustClick()
		stealth.ShortPause()
		stealth.TypeLikeHuman(msgBox, message)

		log.(interface{ Info(string) }).Info("Sent message to " + *id)
		_ = st.MarkMessaged(*id)

		stealth.LongPause()
	}

	// ===============================
	// PAGINATION (MOCK)
	// ===============================
	if search.HasNextPage(page) {
		log.(interface{ Info(string) }).Info("Navigating to next page")
		search.GoToNextPage(page)
	}

	time.Sleep(30 * time.Second)
	return nil
}
