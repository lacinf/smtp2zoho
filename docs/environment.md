# Environment Variables

This project is configured entirely through environment variables. Below is the list of required and optional variables, along with their descriptions and default values (if any).

> 💡 You can use a `.env` file in your project root to load these variables automatically when running locally or via Docker.

---

## 🔐 Zoho API Configuration

| Variable             | Required | Description                                                                                                            |
|----------------------|----------|------------------------------------------------------------------------------------------------------------------------|
| `ZOHO_API_URL`       | ✅       | Full Zoho Mail API URL, including your account ID. Example:<br>`https://mail.zoho.com/api/accounts/123456789/messages` |
| `ZOHO_FROM_ADDRESS`  | ✅       | Email address configured and authorized in your Zoho Mail account.                                                     |
| `ZOHO_CLIENT_ID`     | ✅       | OAuth2 client ID provided by Zoho when registering your app.                                                           |
| `ZOHO_CLIENT_SECRET` | ✅       | OAuth2 client secret from Zoho.                                                                                        |
| `ZOHO_REFRESH_TOKEN` | ✅       | Long-lived refresh token obtained from Zoho's token flow.                                                              |

---

## 📬 SMTP Server Configuration

| Variable              | Required | Default     | Description                                                                 |
|-----------------------|----------|-------------|-----------------------------------------------------------------------------|
| `SMTP_PORT`           | ❌       | `25`        | Port number to bind the SMTP server.                                       |
| `SMTP_AUTH_REQUIRED`  | ❌       | `false`     | Whether SMTP AUTH is required. Accepts `true`, `false`, `1`, or `0`.       |
| `SMTP_USER`           | ❌       | `"user"`    | Username required for SMTP AUTH (if enabled).                              |
| `SMTP_PASSWORD`       | ❌       | `"password"`| Password for SMTP AUTH (if enabled).                                       |

---

## 📄 Logging Configuration

| Variable     | Required | Default | Description                                                                 |
|--------------|----------|---------|-----------------------------------------------------------------------------|
| `LOG_LEVEL`  | ❌       | `info`  | Log verbosity: `error`, `info`, or `debug`.                                 |

---

## 🧪 Versioned `.env` loading

At startup, the application attempts to load an `.env` file named after the executable.

For example:
- If the binary is named `smtp2zoho`, it will attempt to load `smtp2zoho.env`.

This behavior allows multiple `.env` files for different environments or stages.

> This loading logic is implemented using `github.com/joho/godotenv`.

---

## 📂 Example file

See [`.env.example`](../.env.example) at the root of the repository for a ready-to-copy template.
