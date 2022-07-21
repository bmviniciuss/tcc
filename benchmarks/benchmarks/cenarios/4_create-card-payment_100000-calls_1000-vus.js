import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/helpers.js'

const GATEWAY_HOST = 'localhost:5000'

export const options = {
  vus: 1000,
  iterations: 100000
}

export default function () {
  const url = `http://${GATEWAY_HOST}/api/payments/card`

  const payload = JSON.stringify({
    client_id: '56c4a18b-f67a-48b9-a914-98381bd995b9',
    payment_type: 'CREDIT_CARD',
    payment_date: '2022-07-14T16:32:21.147Z',
    amount: 1000,
    payment_info: {
      card_token: '8b6ad026d5784b06b8e0ff30b1fcaeec589d0a9783224fd9b90b546bcc2ec965'
    }
  })

  const params = {
    headers: {
      'Content-Type': 'application/json'
    }
  }

  const res = http.post(url, payload, params)
  check(res, {
    'is status 201': (r) => r.status === 201
  })
}

export function handleSummary (data) {
  return generateData('100000calls-1000vus-create-card-payment', data)
}
