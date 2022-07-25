import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/helpers.js'

const testConfig = {
  id: "c4e77cc4-0103-4abc-ac9e-da0f78901f4e",
  name: "create-card-payment",
  vus: 50,
  duration: '1m',
  executedAt: new Date().toISOString()
}

const GATEWAY_HOST = 'localhost:5000'

export const options = {
  vus: testConfig.vus,
  duration: testConfig.duration
}

export default function () {
  const url = `http://${GATEWAY_HOST}/api/payments/card`
  const payload = JSON.stringify({
    "client_id": "8fd14b79-956f-4261-a509-2efe63c6de39",
    "payment_type": "CREDIT_CARD",
    "payment_date": "2022-07-25T19:52:42.972Z",
    "amount": 1000,
    "payment_info": {
        "card_token": "ede2131f7b354ad98a8ff60f4ba529cb783b21025ed14f448af669a119fa9957"
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
  return generateData(`${testConfig.vus}-vus-${testConfig.duration}-${testConfig.name}`, testConfig, data)
}