# Cartões

- [Cartões](#cartões)
  - [Definição](#definição)
  - [Funcionalidades](#funcionalidades)
    - [Geração de Cartões](#geração-de-cartões)
      - [REST](#rest)
      - [gRPC](#grpc)
    - [Busca de cartão por token](#busca-de-cartão-por-token)
      - [REST](#rest-1)
      - [gRPC](#grpc-1)

## Definição
O micro serviço de cartões tem como objetivo gerenciar os cartões que são utilizados em toda a aplicação. Permite a geração de cartões aleatórios e a busca de cartão por token único.


## Funcionalidades
Cada funcionalidade do sistema é exposta em APIs REST e gRPC.

### Geração de Cartões
Geração - mock - de um cartão de crédito e/ou débito para um usuário. Sendo necessário informar o nome do portador e especificar se o cartão é de crédito e/ou débito. Como o objetivo desse micro serviço é simular uma geração de cartão, o processo de tokenização do cartão é simples e consiste na concatenação de dois UUID's V4 mantendo apenas os caracteres alfanuméricos.

#### REST
> POST {{CARD_HOST}}/api/cards

Corpo da requisição:
```json
{
  "cardholder_name": "Vinicius Barbosa",
  "is_credit": true,
  "is_debit": true
}
```

Resposta da requisição:
```json
{
  "id": "d4c56470-662c-41ca-a172-40eb64c8dcf5",
  "cardholder_name": "Vinicius Barbosa",
  "token": "dbf499ad94d94fe7817b2d9d2de182ef0d2b1d47ba32417aa1cd3e24c1ccb244",
  "masked_number": "3438*******9301",
  "expiration_year": 2027,
  "expiration_month": 7,
  "active": true,
  "is_credit": true,
  "is_debit": true
}
```

#### gRPC
> Host: {{CARD_HOST}} - Cards.GenerateCard

Mensagem da requisição:
```json
{
  "cardholder_name": "Vinicius Barbosa",
  "is_credit": true,
  "is_debit": true
}
```

Resposta:
```json
{
  "id": "e4af7a66-d31e-4510-bde1-dd5b8a56d959",
  "cardholder_name": "Vinicius Barbosa",
  "number": "343428111325241",
  "cvv": "713",
  "token": "64d54b8d24e0467b804e29e7f47dda957ef0b2b39e9c4a2dac2ff28077d359db",
  "masked_number": "3434*******5241",
  "expiration_year": "2027",
  "expiration_month": "7",
  "active": true,
  "is_credit": true,
  "is_debit": true
}
```

### Busca de cartão por token
Permite a busca de um cartão pelo seu token gerado no processo de tokenização da geração.

#### REST
> GET {{CARD_HOST}}/api/cards?token={{CARD_TOKEN}}

Resposta:
```json
{
  "id": "d4c56470-662c-41ca-a172-40eb64c8dcf5",
  "cardholder_name": "Vinicius Barbosa",
  "token": "dbf499ad94d94fe7817b2d9d2de182ef0d2b1d47ba32417aa1cd3e24c1ccb244",
  "masked_number": "3438*******9301",
  "expiration_year": 2027,
  "expiration_month": 7,
  "active": true,
  "is_credit": true,
  "is_debit": true
}
```

#### gRPC
> Host: {{CARD_HOST}} - Cards.GetCardByToken
Request:
```json
{
  "token": "dbf499ad94d94fe7817b2d9d2de182ef0d2b1d47ba32417aa1cd3e24c1ccb244"
}
```

```json
{
  "id": "d4c56470-662c-41ca-a172-40eb64c8dcf5",
  "cardholder_name": "Vinicius Barbosa",
  "token": "dbf499ad94d94fe7817b2d9d2de182ef0d2b1d47ba32417aa1cd3e24c1ccb244",
  "masked_number": "3438*******9301",
  "expiration_year": "2027",
  "expiration_month": "7",
  "active": true,
  "is_credit": true,
  "is_debit": true
}
```
