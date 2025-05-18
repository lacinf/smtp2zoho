# Troubleshooting Guide

This document helps you diagnose and resolve common issues when running `smtp2zoho`.

---

## ✅ Verify required environment variables

Ensure that all mandatory variables are set, either via `.env` or the environment:

- `ZOHO_API_URL`
- `ZOHO_FROM_ADDRESS`
- `ZOHO_CLIENT_ID`
- `ZOHO_CLIENT_SECRET`
- `ZOHO_REFRESH_TOKEN`

> You can run the app with `LOG_LEVEL=debug` to see detailed output, including which variables are missing or misconfigured.

---

## 📭 SMTP not receiving emails

Check the following:

- Is the process running and listening on the correct port?  
  Use:

```bash
netstat -tuln | grep :25
```

> 🧪 If `netstat` is not available, try: `ss -tuln | grep :25`


- If using Docker, ensure the port is exposed (`-p 2525:25`) and you're connecting to the right mapped port.
- Try sending a test email using `swaks`:

```bash
swaks --to test@yourdomain.com --server 127.0.0.1
```

---

## 🔐 SMTP authentication failing

If `SMTP_AUTH_REQUIRED=true`, confirm:

- `SMTP_USER` and `SMTP_PASSWORD` are correctly set
- The sending client uses AUTH PLAIN or LOGIN (as `swaks` does)

Test with:

```bash
swaks --to test@yourdomain.com \
      --server 127.0.0.1 \
      --auth \
      --auth-user "user" \
      --auth-password "password"
```

---

## 🔑 Failed to obtain Zoho access token

Typical causes:

- Invalid `ZOHO_CLIENT_ID` or `ZOHO_REFRESH_TOKEN`
- Client not authorized in the Zoho console
- Token expired or revoked

To test manually:

```bash
curl -X POST "https://accounts.zoho.com/oauth/v2/token" \
  -d "refresh_token=YOUR_REFRESH_TOKEN" \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "grant_type=refresh_token"
```

You should receive a JSON response with an `access_token`. If not, review your Zoho API configuration.

> 📘 If you haven’t obtained your `refresh_token` yet, follow the guide at [zoho-auth.md](zoho-auth.md).

---

## 📤 Failed to send JSON to API

Check the following:

- Is `ZOHO_API_URL` correctly set and reachable?
- Is the access token valid (200 OK from Zoho)?
- Are you behind a firewall or proxy blocking outbound HTTP?

Logs will show the response status, such as:

```text
[ERROR] send failed (status: 401 Unauthorized)
```

---

## 🪵 Enabling useful logs

Use the environment variable:

```bash
LOG_LEVEL=debug
```

This will show:

- All environment variable loading steps
- Email parsing steps
- Token exchange results
- Final POST payload and response

> Logs are printed to stdout by default.
