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
	// SEARCH INPUT (DEMO VISUAL)
	// ===============================
	searchBox, err := page.Element("#search")
	if err == nil {
		stealth.MoveMouseBezier(
			page,
			stealth.Point{X: 200, Y: 200},
			stealth.Point{X: 420, Y: 260},
		)
		stealth.ShortPause()

		stealth.TypeLikeHuman(searchBox, "Software Engineer")
		stealth.LongPause()

		if btn, err := page.Element("#searchBtn"); err == nil {
			btn.MustClick()
			stealth.LongPause()
		}
	}

	// ===============================
	// SEARCH RESULTS (MOCKED)
	// ===============================
	profiles, err := search.CollectProfiles(page)
	if err != nil {
		return err
	}

	for _, p := range profiles {

		// ---- find profile card safely ----
		card, err := page.Element(`[data-profile-id="` + p.ID + `"]`)
		if err != nil {
			log.(interface{ Info(string) }).Info(
				"Profile card not found for " + p.ID + ", skipping",
			)
			continue
		}

		// ===============================
		// CONNECT FLOW
		// ===============================
		connectBtn, err := card.Element(".connect")
		if err == nil {

			if sent, _ := st.HasSent(p.ID); sent {
				log.(interface{ Info(string) }).Info(
					"Already sent request to " + p.ID,
				)
				continue
			}

			box := connectBtn.MustShape().Box()

			from := stealth.Point{X: box.X - 50, Y: box.Y}
			to := stealth.Point{
				X: box.X + box.Width/2,
				Y: box.Y + box.Height/2,
			}

			stealth.MoveMouseBezier(page, from, to)
			time.Sleep(300 * time.Millisecond)
			connectBtn.MustClick()

			log.(interface{ Info(string) }).Info("Sent request to " + p.ID)
			_ = st.MarkSent(p.ID)

			stealth.LongPause()
			continue
		}

		// ===============================
		// MESSAGE FLOW
		// ===============================
		msgBtn, err := card.Element(".message")
		if err == nil {

			if done, _ := st.HasMessaged(p.ID); done {
				log.(interface{ Info(string) }).Info(
					"Already messaged " + p.ID,
				)
				continue
			}

			msgBox, err := card.Element(".msgbox")
			if err != nil {
				continue
			}

			msgBtn.MustClick()
			stealth.ShortPause()

			message := messaging.Render(
				"Hi {{name}}, thanks for connecting!",
				map[string]string{
					"name": p.ID,
				},
			)

			stealth.TypeLikeHuman(msgBox, message)

			log.(interface{ Info(string) }).Info("Sent message to " + p.ID)
			_ = st.MarkMessaged(p.ID)

			stealth.LongPause()
			continue
		}

		// ===============================
		// NOTHING TO DO
		// ===============================
		log.(interface{ Info(string) }).Info(
			"No actionable buttons for " + p.ID + ", skipping",
		)
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
