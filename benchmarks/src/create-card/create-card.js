const { GATEWAY_HOST } = require('../config/env.js')
const autocannon = require('autocannon')
const { writeFile } = require('fs/promises')
const path = require('path')

const run = async (isGRPC = false) => {
  const startDate = new Date()
  const fileName = `${startDate.toISOString()}-create-card-${isGRPC ? 'grpc' : 'http'}.json`

  const result = await autocannon({
    url: `http://${GATEWAY_HOST}/api`,
    requests: [
      {
        path: `http://${GATEWAY_HOST}/api/cards`,
        method: 'POST',
        body: JSON.stringify({
          cardholder_name: 'Vinicius Barbosa',
          is_credit: true,
          is_debit: true
        }),
        headers: {
          'Content-Type': 'application/json'
        },
        onResponse: (status, body, context) => {
          if (status === 200) {
            context.card_token = JSON.parse(body).token
          } else {
            console.log('context: ', context)
            console.log(status, body)
            throw new Error('Erro em request')
          }
        }
      }
      // {

      //   method: 'POST',
      //   path: `http://${GATEWAY_HOST}/api/payments/card`,
      //   headers: {
      //     'Content-Type': 'application/json'
      //   },
      //   setupRequest: (req, context) => ({
      //     ...req,
      //     body: JSON.stringify({
      //       client_id: '79750f54-ef8f-48f8-acaf-6fd66cd9843f',
      //       payment_type: 'CREDIT_CARD',
      //       payment_date: '2022-06-29T03:55:18.436Z',
      //       amount: 1000,
      //       payment_info: {
      //         card_token: context.card_token
      //       }
      //     })

      //   })
      // }
    ],
    amount: 10000
  })

  const s = autocannon.printResult(result, {
    renderLatencyTable: true,
    renderResultsTable: true
  })

  console.log(s)

  // console.table(result)

  // await writeFile(path.join('src', 'create-card', 'benchmarks', fileName), JSON.stringify(result))
}

run(true)
