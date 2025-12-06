package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func autoVersion() string {
	out, err := exec.Command("git", "describe", "--tags", "--always").Output()
	if err != nil {
		return "BLUE"
	}
	return strings.TrimSpace(string(out))
}

var version = autoVersion()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		html := `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <title>App Version 1</title>
  <meta name="viewport" content="width=device-width, initial-scale=1" />

  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; }

    body {
      font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
      min-height: 100vh;
      display: flex;
      align-items: center;
      justify-content: center;
      background:
        radial-gradient(circle at 10% 20%, #22c55e15 0, transparent 50%),
        radial-gradient(circle at 90% 80%, #22d3ee15 0, transparent 55%),
        radial-gradient(circle at 50% 10%, #3b82f615 0, transparent 45%),
        #020617;
      color: #e5e7eb;
    }

    .card {
      position: relative;
      padding: 3rem 4rem;
      border-radius: 1.75rem;
      background: radial-gradient(circle at top left, #0f172a, #020617 70%);
      border: 1px solid rgba(148, 163, 184, 0.5);
      box-shadow:
        0 0 0 1px rgba(15, 23, 42, 0.9),
        0 25px 70px rgba(15, 23, 42, 0.95);
      text-align: center;
      max-width: 520px;
      width: min(90%, 520px);
      overflow: hidden;
    }

    /* Animated border glow */
    .card::before {
      content: "";
      position: absolute;
      inset: -2px;
      border-radius: inherit;
      background:
        conic-gradient(
          from 0deg,
          #22c55e,
          #22d3ee,
          #3b82f6,
          #a855f7,
          #22c55e
        );
      opacity: 0.7;
      filter: blur(10px);
      z-index: -1;
      animation: border-glow 7s linear infinite;
    }

    @keyframes border-glow {
      from { transform: rotate(0deg); }
      to   { transform: rotate(360deg); }
    }

    .badge {
      display: inline-block;
      padding: 0.25rem 0.9rem;
      border-radius: 999px;
      font-size: 0.75rem;
      letter-spacing: 0.16em;
      text-transform: uppercase;
      margin-bottom: 1.2rem;
      background: rgba(56, 189, 248, 0.08);
      border: 1px solid rgba(56, 189, 248, 0.7);
      color: #e0f2fe;
    }

    .version {
      font-size: clamp(2.4rem, 6vw, 3.4rem);
      font-weight: 800;
      letter-spacing: 0.22em;
      text-transform: uppercase;
      background: linear-gradient(90deg, #2276c5ff, #22d3ee, #3b82f6);
      -webkit-background-clip: text;
      background-clip: text;
      color: transparent;
      margin-bottom: 1rem;
    }

    /* 🔥 Neon pulsing version number */
    .version-number {
      font-size: clamp(3.5rem, 9vw, 5.5rem);
      font-weight: 900;
      letter-spacing: 0.08em;
      color: #e0f2fe;
      text-shadow:
        0 0 8px #22d3ee,
        0 0 18px #22d3ee,
        0 0 30px #3b82f6,
        0 0 50px rgba(56, 189, 248, 0.9);
      animation: pulse-glow 1.4s ease-in-out infinite;
      margin-bottom: 1.5rem;
    }

    @keyframes pulse-glow {
      0% {
        transform: scale(1);
        opacity: 0.85;
      }
      50% {
        transform: scale(1.08);
        opacity: 1;
      }
      100% {
        transform: scale(1);
        opacity: 0.85;
      }
    }

    .desc {
      font-size: 0.98rem;
      color: #cbd5f5;
      opacity: 0.9;
      max-width: 26rem;
      margin: 0 auto;
    }

    .footer {
      margin-top: 2rem;
      font-size: 0.8rem;
      color: #6b7280;
    }
  </style>

</head>
<body>
  <div class="card">
    <div class="badge">Application Build</div>
    <div class="version">VERSION</div>

    <div class="version-number">` + version + `</div>

    <p class="desc">
      This is <strong>App Version </strong> running over HTTP.
    </p>

    <div class="footer">Listening on http://localhost:8080</div>
  </div>
</body>
</html>`

		fmt.Fprint(w, html)
	})

	log.Printf("Version1 running with version=%s", version)
	http.ListenAndServe(":8080", nil)
}
