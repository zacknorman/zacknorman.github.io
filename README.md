# Zack Norman — New Eleventy + Go contact scaffold

This scaffolds a lean Eleventy site with a small Go contact handler.

Local dev

```bash
npm ci
npm run dev
# run the contact server separately if you want to test it locally
go run tools/contact/main.go
```

Build (CI)

GitHub Actions builds the static site (`_site`) and the Go binary (`bin/contact`). Artifacts are uploaded by the workflow. To deploy to your VPS, enable the commented `rsync` step and add the required secrets.

Contact server

The Go server listens on `/contact` and accepts JSON POSTs with `{name,email,message}`. If you set `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASS`, and `TO_EMAIL` in the environment, messages will be sent via SMTP; otherwise they are appended to `contact_messages.log` in the working directory.
# zacknorman.github.io
here's a bon jovi pun

__It's my site__! Now or never! `I dont wanna dev forever!`
