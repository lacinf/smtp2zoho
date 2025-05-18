#!/bin/bash
# Exemplo de teste Zoho API - NÃO USAR CREDENCIAIS REAIS NESTE SCRIPT
# Substitua pelos seus valores apenas em ambiente local seguro.

# Definir variáveis temporárias
export ZOHO_API_URL="https://mail.zoho.com/api/accounts/SEU_ACCOUNT_ID/messages"
export ZOHO_FROM_ADDRESS="seu@email.com"
export ZOHO_TO_ADDRESS="destino@email.com"
export ZOHO_CLIENT_ID="YOUR_CLIENT_ID"
export ZOHO_CLIENT_SECRET="YOUR_CLIENT_SECRET"
export ZOHO_REFRESH_TOKEN="YOUR_REFRESH_TOKEN"

# Obter access token
ACCESS_TOKEN=$(curl -s -X POST "https://accounts.zoho.com/oauth/v2/token" \
  -d "refresh_token=$ZOHO_REFRESH_TOKEN" \
  -d "client_id=$ZOHO_CLIENT_ID" \
  -d "client_secret=$ZOHO_CLIENT_SECRET" \
  -d "grant_type=refresh_token" | jq -r '.access_token')

# Verificar se obteve token
if [[ "$ACCESS_TOKEN" == "null" || -z "$ACCESS_TOKEN" ]]; then
  echo "Erro ao obter Access Token."
  exit 1
fi

# Enviar email de teste
RESPONSE=$(curl -s -X POST "$ZOHO_API_URL" \
  -H "Authorization: Zoho-oauthtoken $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "fromAddress": "'"$ZOHO_FROM_ADDRESS"'",
    "toAddress": "'"$ZOHO_TO_ADDRESS"'",
    "subject": "Teste SMTP2API - Confirmação de credenciais",
    "content": "Este é um teste automático para validar as credenciais atuais via terminal."
  }')

# Verificar resposta
echo "$RESPONSE" | grep -q '"status":{"code":200'
if [[ $? -eq 0 ]]; then
  echo "Envio realizado com sucesso."
else
  echo "Falha no envio. Resposta:"
  echo "$RESPONSE"
fi

# Limpar variáveis
unset ZOHO_API_URL ZOHO_FROM_ADDRESS ZOHO_TO_ADDRESS ZOHO_CLIENT_ID ZOHO_CLIENT_SECRET ZOHO_REFRESH_TOKEN ACCESS_TOKEN RESPONSE
