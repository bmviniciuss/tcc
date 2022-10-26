import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/helpers.js'

const testConfig = {
  id: "c4e77cc4-0103-4abc-ac9e-da0f78901f4e",
  name: "create-card-payment",
  vus: 50,
  duration: '5m',
  executedAt: new Date().toISOString()
}

const GATEWAY_HOST = 'localhost:5000'

export const options = {
  vus: testConfig.vus,
  duration: testConfig.duration,
  summaryTrendStats: ["min", "med", "avg", "max", "p(90)", "p(95)", "p(99)" ],
  thresholds: {
    checks: [{ threshold: 'rate>0.99', abortOnFail: true }],
  },
  discardResponseBodies: true
}



export function setup() {
  const cardUrl = `http://${GATEWAY_HOST}/api/cards`

  const payload = JSON.stringify({
    cardholder_name: 'Vinicius Barbosa',
    is_credit: true,
    is_debit: true
  })

  const params = {
    headers: {
      'Content-Type': 'application/json'
    }
  }

  const res = http.post(cardUrl, payload, params)
  return { card: res.json() };
}


export default function ({card}) {
  const url = `http://${GATEWAY_HOST}/api/payments/card`
  const payload = JSON.stringify({
    "client_id": "8fd14b79-956f-4261-a509-2efe63c6de39",
    "payment_type": "CREDIT_CARD",
    "payment_date": "2022-07-25T19:52:42.972Z",
    "amount": 1000,
    "payment_info": {
        "card_token": card.token
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
