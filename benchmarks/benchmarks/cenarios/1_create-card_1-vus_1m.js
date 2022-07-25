import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/helpers.js'

const testConfig = {
  id: "36c812b5-ce7e-46d6-8bf2-1b749fd2c4a6",
  name: "create-card-payment",
  vus: 1,
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
    "client_id": "{{$guid}}",
    "payment_type": "CREDIT_CARD",
    "payment_date": "{{$isoTimestamp}}",
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
