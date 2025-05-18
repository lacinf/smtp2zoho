# API Payload Specification

`smtp2zoho` converts incoming emails into structured JSON and sends them via HTTP POST to the `ZOHO_API_URL` endpoint, using Zoho's OAuth2 token authentication.

This document describes the format and structure of that request.

---

## 🔗 Request

- **Method:** `POST`
- **URL:** value from the environment variable `ZOHO_API_URL`  
  (e.g. `https://mail.zoho.com/api/accounts/123456789/messages`)
- **Headers:**
  - `Authorization: Zoho-oauthtoken <access_token>`
  - `Content-Type: application/json`

---

## 📨 Payload Format

The JSON payload sent in the request body includes:

- The entire payload is sent as the raw body of the request.

| Field         | Type   | Description                                          |
|---------------|--------|------------------------------------------------------|
| `fromAddress` | string | Always set to the value of `ZOHO_FROM_ADDRESS`       |
| `toAddress`   | string | Parsed from the SMTP `To:` field                     |
| `subject`     | string | Parsed from the email `Subject:` header              |
| `content`     | string | Parsed plain-text body of the email                  |

> 🧼 All other email fields (Cc, Bcc, custom headers, HTML body, attachments) are ignored.

---

## 📦 Example Payload

```json
{
  "fromAddress": "example@domain.com",
  "toAddress": "test@destination.com",
  "subject": "Test from swaks",
  "content": "This is a simple plain-text body"
}
```

---

## ✅ Expected Response

- **Status Code:** `200 OK`
- No content is expected in the response body.

> A `200 OK` response from Zoho indicates that the message was successfully accepted for delivery.

If the request fails (e.g., expired token, invalid config), the error will be logged and the email will **not** be retried by default.

> 📋 Logs will include any failure with the Zoho response status for easier troubleshooting.
