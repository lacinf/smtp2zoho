# Zoho OAuth Setup: Getting Your Refresh Token and Account ID

This guide explains how to obtain the necessary credentials to use `smtp2zoho` by creating a Zoho Self Client and retrieving:

- `refresh_token` — used to authenticate API requests
- `account_id` — needed to build the Zoho API URL

> 💡 This guide uses `jq` to extract values from JSON responses.

> If not installed, run: `sudo apt install jq` (Ubuntu/Debian) or `apk add jq` (Alpine).

---

## 🔧 Step 1: Create a Self Client in Zoho API Console

1. Visit: [https://api-console.zoho.com/](https://api-console.zoho.com/)
2. Click **"Add Client"**
3. Choose **"Self Client"**
4. Give it any name, like `smtp2zoho`, and create it

Once created, click on the Self Client to open its details panel.

---

## 🎯 Step 2: Generate the authorization code

In the Self Client panel, click **"Generate Code"**, then:

- Set **Scope** to:

```
ZohoMail.accounts.READ,ZohoMail.messages.CREATE
```

- Set **Time Duration** to `10 minutes`
- Click **Generate**

You will get a `code` string that expires in 10 minutes.

---

## 🔁 Step 3: Exchange the code for a refresh token

Use the following `curl` command — replace the values accordingly:

```bash
curl -s -X POST "https://accounts.zoho.com/oauth/v2/token" \
  -d "grant_type=authorization_code" \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=YOUR_CODE" \
  -d "redirect_uri=http://localhost" | jq -r '.refresh_token'
```
> 🔒 `redirect_uri` must match the value configured when registering the client (even if unused).

> ✅ This will print only the `refresh_token` string.

---

## 🆔 Step 4: Get your Zoho Mail account ID

Now use your refresh token to get an access token:

```bash
curl -s -X POST "https://accounts.zoho.com/oauth/v2/token" \
  -d "refresh_token=YOUR_REFRESH_TOKEN" \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "grant_type=refresh_token" | jq -r '.access_token'
```

Then use the `access_token` to fetch your `account_id`:

```bash
curl -s -X GET "https://mail.zoho.com/api/accounts" \
  -H "Authorization: Zoho-oauthtoken YOUR_ACCESS_TOKEN" | jq -r '.data[0].accountId'
```
> 📘 The account ID is usually at position `[0]` in the `data` array.

> ✅ This will return your `account_id`, which is needed in your `ZOHO_API_URL`.

Example URL:

```
https://mail.zoho.com/api/accounts/123456789/messages
```

---

## ✅ Step 5: Add credentials to your `.env`

```env
ZOHO_CLIENT_ID=your_client_id
ZOHO_CLIENT_SECRET=your_client_secret
ZOHO_REFRESH_TOKEN=your_refresh_token
ZOHO_API_URL=https://mail.zoho.com/api/accounts/your_account_id/messages
```

> 🎯 You're now ready to run `smtp2zoho` with valid credentials.

