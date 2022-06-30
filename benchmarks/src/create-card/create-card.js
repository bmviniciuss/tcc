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
        path: '/api/cards',
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
          if (status >= 200 && status <= 299) {
            context.card_token = JSON.parse(body).token
          } else {
            console.log('ERROR: ', status, body, context)
            throw new Error('Erro em request')
          }
        }
      }
    ],
    amount: 500_000
  })

  const s = autocannon.printResult(result, {
    renderLatencyTable: true,
    renderResultsTable: true
  })

  console.log(s)

  await writeFile(path.join('src', 'create-card', 'benchmarks', fileName), JSON.stringify(result))
}

async function main () {
  const grpc = false
  console.log(`${grpc ? 'gRPC' : 'HTTP'} Gateway Benchmarking`)
  const n = 1
  for (let i = 0; i < n; i++) {
    console.log(`Executing ${i + 1}/${n}`)
    await run(grpc)
  }
  console.log('Done')
}

main()
