# smtp2zoho

**Lightweight SMTP server that captures emails and forwards parsed JSON payloads via HTTP — built to integrate with Zoho Mail.**

---

## 🚀 Overview

`smtp2zoho` is a minimal SMTP server designed to run locally or in containers.  
It listens for authenticated email submissions and converts each message into a clean JSON payload, which is forwarded to a configured Zoho Mail API endpoint.

- ✅ SMTP with optional AUTH
- ✅ JSON transformation
- ✅ HTTP POST with OAuth2 token
- ✅ Docker-ready
- ✅ Zero database or queue dependency

---

## 📦 Installation

See: [docs/installation.md](docs/installation.md)

```bash
git clone https://github.com/lacinf/smtp2zoho.git
cd smtp2zoho
make build
```

---

## ⚙️ Configuration

All settings are provided via environment variables.

See: [docs/environment.md](docs/environment.md)  
Template: [`.env.example`](.env.example)

> 🔑 Need a Zoho refresh token and account ID?  
> See: [docs/zoho-auth.md](docs/zoho-auth.md)

---

## ✉️ Usage

To test sending emails to your running server:

```bash
swaks --to test@yourdomain.com --server 127.0.0.1
```

See: [docs/usage.md](docs/usage.md)

---

## 📤 API Output Format

Each email is converted into JSON:

```json
{
  "fromAddress": "example@domain.com",
  "toAddress": "user@example.com",
  "subject": "Hello",
  "content": "Text body only"
}
```

See: [docs/api.md](docs/api.md)

---

## 🧪 Troubleshooting

See: [docs/troubleshooting.md](docs/troubleshooting.md)

---

## 🤝 Contributing

If you'd like to contribute or report an issue, please read [CONTRIBUTING.md](CONTRIBUTING.md).

---

## 📄 License

MIT — see [LICENSE](LICENSE)

---

## 🧠 Maintainer

Built and maintained by [Lacerda Informática](https://github.com/lacinf)
