# Pagamento de Cartões

- [Pagamento de Cartões](#pagamento-de-cartões)
  - [Definição](#definição)
  - [Funcionalidades](#funcionalidades)
    - [Criar Pagamento](#criar-pagamento)
      - [Rest](#rest)
      - [gRPC](#grpc)
    - [Busca de pagamentos de cliente](#busca-de-pagamentos-de-cliente)
      - [Rest](#rest-1)
      - [gRPC](#grpc-1)

## Definição
Responsável por gerenciar (criar e listar) pagamentos que foram feitos através de cartões de crédito e débito para clientes do gateway.

## Funcionalidades
### Criar Pagamento
Simula o pagamento de valores pro cartão de uma portador para um cliente do gateway.

Para requisição de criação de pagamento o micro serviço **busca o cartão pelo token no micro serviço de cartões** para realizar validações, simulando assim um fluxo de validações que poderia acontecer em um sistema verdadeiro.

São cobradas taxas sobre o valor da transação que varia de acordo com o tipo de cartão utilizado no pagamento:

| Método  	| Porcentagem 	|
|---------	|-------------	|
| Crédito 	| 3%           	|
| Debito  	| 1%           	|

#### Rest
> POST {{CARD_PAYMENT_HOST}}/api/payment

Corpo da requisição:
```json
{
  "client_id": "5dea0dda-01ae-4c75-8b55-0065dfa2264b",
  "payment_type": "CREDIT_CARD",
  "payment_date": "2022-07-07T23:48:58.585Z",
  "amount": 1000,
  "payment_info": {
    "card_token": "45b6be6583224686a47b7a9202b2f8694ff5d7c177f047c2b6959a651e0f2758"
  }
}
```

Resposta:
```json
{
  "id": "4b6657cd-d133-468f-aeb0-ea648cd01fe4",
  "client_id": "5dea0dda-01ae-4c75-8b55-0065dfa2264b",
  "amount": 1000,
  "payment_type": "CREDIT_CARD",
  "payment_date": "2022-07-07T23:48:58.585Z",
  "payment_info": {
    "masked_number": "5268********6114"
  }
}
```

#### gRPC
> {{CARD_PAYMENT_HOST}} CardPayment.ProcessCardPayment
Corpo requisição:
```json
{
  "amount": 100.65,
  "client_id": "705bb173-13db-45f8-b9b2-bd0103bec559",
  "payment_date": "2022-07-07T23:54:35Z",
  "payment_info": {
    "card_token": "ede25113a0de4f8aa6e71a84704a2ee47f09ca8c4ddc45c897962c3b426be4ba"
  },
  "payment_type": "CREDIT_CARD"
}
```

Resposta:
```json
{
  "id": "c933ae0f-f88f-4ad7-a9ec-edcf5e589e19",
  "client_id": "705bb173-13db-45f8-b9b2-bd0103bec559",
  "payment_type": "CREDIT_CARD",
  "payment_date": "2022-07-07T23:54:35Z",
  "amount": 100.65,
  "payment_info": {
    "masked_number": "3578********6138"
  }
}
```

### Busca de pagamentos de cliente
O sistema deve ser capaz de listar os pagamentos de cartão para um cliente pelo seu identificador.

#### Rest
> GET {{CARD_PAYMENT_HOST}}/api/payment?client_id={{CLIENT_ID}}

Reposta:

```json
{
  "content": [
    {
      "id": "c6c2dc63-49a6-4e2d-a8f1-9a4780bba31b",
      "client_id": "15c9de42-c721-43d1-9fbd-447bf0a74b9b",
      "amount": 100.65,
      "payment_type": "CREDIT_CARD",
      "payment_date": "2022-07-07T23:57:50.539Z",
      "payment_info": {
        "masked_number": "3578********6138"
      }
    },
    {
      "id": "118b5058-406a-4034-b2cb-469b0582f9ac",
      "client_id": "15c9de42-c721-43d1-9fbd-447bf0a74b9b",
      "amount": 50,
      "payment_type": "CREDIT_CARD",
      "payment_date": "2022-07-07T23:58:01.504Z",
      "payment_info": {
        "masked_number": "3578********6138"
      }
    }
  ]
}
```


#### gRPC
> {{CARD_PAYMENT_HOST}} CardPayment.GetPaymentsByClientId
Mensagem requisição:

```json
{
  "client_id": "15c9de42-c721-43d1-9fbd-447bf0a74b9b"
}
```


Reposta:

```json
{
  "content": [
    {
      "id": "c6c2dc63-49a6-4e2d-a8f1-9a4780bba31b",
      "client_id": "15c9de42-c721-43d1-9fbd-447bf0a74b9b",
      "payment_type": "CREDIT_CARD",
      "payment_date": "2022-07-07T23:57:50Z",
      "amount": 100.65,
      "payment_info": {
        "masked_number": "3578********6138"
      }
    },
    {
      "id": "118b5058-406a-4034-b2cb-469b0582f9ac",
      "client_id": "15c9de42-c721-43d1-9fbd-447bf0a74b9b",
      "payment_type": "CREDIT_CARD",
      "payment_date": "2022-07-07T23:58:01Z",
      "amount": 50,
      "payment_info": {
        "masked_number": "3578********6138"
      }
    }
  ]
}
```
