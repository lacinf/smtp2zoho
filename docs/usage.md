# Usage Guide

This guide explains how to test and use `smtp2zoho` locally. It covers sending emails with and without SMTP authentication, including tests against a running Docker container.

---

## 🚀 Basic Flow

`smtp2zoho` acts as a lightweight SMTP server. When it receives an email, it extracts:

- Subject
- Plain-text body

Then it builds a JSON object and sends it via HTTP POST to the URL defined in `ZOHO_API_URL`.

---

## ✉️ Sending without SMTP authentication

If `SMTP_AUTH_REQUIRED=false` (default), you can send mail directly to the server.

Recommended tool: [`swaks`](https://www.jetmore.org/john/code/swaks/)

> Install on Ubuntu: `sudo apt install swaks`

If running the app **locally**:

```bash
swaks --to test@yourdomain.com \
      --server 127.0.0.1 \
      --data "Subject: Hello from swaks\n\nThis is a test email."
```

If running inside a **Docker container** and exposed on port 2525:

```bash
docker run --rm --env-file .env -p 2525:25 lacinf/smtp2zoho:latest

# Then in another terminal:
swaks --to test@yourdomain.com --server 127.0.0.1 --port 2525 \
      --data "Subject: Hello from Docker\n\nTesting container run."
```

> 💡 You should see success logs from `smtp2zoho` and a confirmation that the JSON was sent to the API.

---

## 🔐 Sending with SMTP AUTH

If `SMTP_AUTH_REQUIRED=true`, you must provide credentials via `SMTP_USER` and `SMTP_PASSWORD`.

```bash
swaks --to test@yourdomain.com \
      --server 127.0.0.1 \
      --auth \
      --auth-user "user" \
      --auth-password "password" \
      --data "Subject: Authenticated email\n\nThis is a test with login."
```

---

## 📜 Logs and what to expect

Once a message is received and parsed:

- A log entry confirms reception
- Another log shows if the HTTP POST was successful
- If `LOG_LEVEL=debug`, you’ll see detailed parsing and payload output

Example log output (simplified):

```text
[INFO] Received email for test@yourdomain.com
[INFO] Parsed subject: Hello from swaks
[INFO] Sending JSON payload to https://mail.zoho.com/api/accounts/...
[INFO] API responded with 200 OK
```

> ❗ If there's an issue with Zoho credentials, you’ll see error messages related to token retrieval or authorization failure.

---

## 🧪 Other tools

Any SMTP client can be used (e.g. `telnet`, Python libs, email clients), as long as it connects to the configured port and obeys AUTH settings if enabled.

We recommend `swaks` for easy scripting and validation.

---

## 📌 Important Notes

- 📤 The `From` field in the original email is **always overridden** by the value of the `ZOHO_FROM_ADDRESS` environment variable.  
  The user-supplied `From:` is ignored.

- ✉️ Only the following fields are parsed and included in the HTTP request:
  - `To`
  - `Subject`
  - Plain-text `Body`

- 🗑️ Any other email headers (e.g. `Cc`, `Bcc`, `Reply-To`, `Message-ID`, custom headers) are **discarded** and will not appear in the forwarded JSON.

> This behavior ensures simplicity and consistency for downstream processing.
