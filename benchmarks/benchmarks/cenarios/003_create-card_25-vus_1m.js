import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/helpers.js'

const testConfig = {
  id: "f31c8249-3080-4a20-9c51-1a8d4fd982c8",
  name: "create-card",
  vus: 25,
  duration: '1m',
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


export default function () {
  const url = `http://${GATEWAY_HOST}/api/cards`
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

  const res = http.post(url, payload, params)
  check(res, {
    'is status 201': (r) => r.status === 201
  })
}

export function handleSummary (data) {
  return generateData(`${testConfig.vus}-vus-${testConfig.duration}-${testConfig.name}`, testConfig, data)
}
