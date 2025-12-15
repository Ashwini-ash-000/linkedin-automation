package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

type Session struct {
	Browser *rod.Browser
}

func NewSession(cfg interface{}, _ interface{}) *Session {
	l := launcher.New().
		NoSandbox(true).
		Leakless(false).
		Headless(false)

	u := l.MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect()

	return &Session{
		Browser: browser,
	}
}

func (s *Session) Close() {
	if s.Browser != nil {
		_ = s.Browser.Close()
	}
}
