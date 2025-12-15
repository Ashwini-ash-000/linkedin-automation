# LinkedIn Automation – Go + Rod (Proof of Concept)

**Built by:** Ashwini C

---

## Project Overview

This project is a **Go-based LinkedIn automation proof-of-concept** built using the **Rod browser automation library**.  
It demonstrates advanced browser automation, human-like interaction patterns, and anti-bot stealth strategies using clean and modular Go architecture.

For ethical and legal reasons, the automation is demonstrated against a **mock LinkedIn-style UI** that closely mirrors real-world DOM structure and interaction flows.

The focus of this project is **automation engineering quality**, not bypassing platform protections.

---

## Features Implemented

### Core Capabilities (PoC)

- Browser lifecycle management with Rod
- Human-like mouse movement using Bézier curves
- Realistic typing simulation with randomized delays
- Search & profile parsing (mocked)
- Connection request handling with daily limits
- Messaging accepted connections
- Duplicate detection using persistent state
- Pagination handling (mocked)
- Structured logging and modular architecture

---

## Project Structure

linkedin-automation/
├── cmd/
│ └── bot/
│ └── main.go
├── internal/
│ ├── auth/
│ ├── browser/
│ ├── config/
│ ├── search/
│ ├── connect/
│ ├── messaging/
│ ├── stealth/
│ ├── store/
│ ├── scheduler/
│ └── logger/
├── mock/
│ └── linkedin.html
├── go.mod
├── go.sum
└── README.md

---

## How to Run

### Prerequisites
- Go 1.21 or later
- Internet connection (Rod downloads Chromium on first run)

### Steps

```bash
go build ./...
go run ./cmd/bot 
```

---

## Expected Behavior

When the bot is executed:

- Chromium browser launches automatically
- Mock LinkedIn page loads
- Mouse moves along natural curved paths
- Typing occurs with realistic human-like delays
- Connection requests are sent where applicable
- Messages are sent only to accepted connections
- Duplicate profiles are skipped
- Pagination is handled
- Logs explain every decision taken
- Browser exits cleanly after one execution cycle

---

## Why a Mock UI Is Used

Automating LinkedIn directly violates LinkedIn’s Terms of Service and can lead to account bans.

This project intentionally uses a **mock LinkedIn-style UI** to:

- Avoid legal and ethical risks
- Prevent real account misuse
- Enable safe, repeatable demonstrations
- Focus evaluation on automation architecture and stealth techniques

The mock UI mirrors real-world behavior including:

- Profile cards
- Connect and Message buttons
- Accepted vs pending connections
- Pagination flow
- DOM structure similar to LinkedIn

---

## Stealth & Anti-Bot Techniques

The project demonstrates **8+ anti-detection techniques**, including all mandatory requirements.

### Mandatory Techniques

- Human-like mouse movement using Bézier curves
- Randomized timing patterns and think delays
- Browser fingerprint masking (automation flags, viewport control)

### Additional Techniques

- Realistic typing simulation with variable delays
- Cursor movement before interactions (no teleport clicks)
- Rate limiting and daily action caps
- Activity scheduling (single-cycle execution)
- Persistent state to avoid repeated actions

---

## State Persistence

Local state is persisted using **SQLite**.

The system tracks:

- Sent connection requests
- Daily request counts
- Messaged profiles

This ensures:

- Idempotent behavior across runs
- Safe interruption and restart
- No duplicate or excessive actions

---

## Limitations (Intentional)

- No real LinkedIn login
- No captcha or 2FA bypass logic
- Search and pagination are mocked
- Not intended for production use

These limitations are deliberate and align with ethical automation practices.

---

## Disclaimer

This project is strictly for **educational and evaluation purposes only**.  
It must not be used against real LinkedIn accounts or production systems.
