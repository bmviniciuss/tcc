import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/src/utils/helpers.js'

const GATEWAY_HOST = 'localhost:5000'
const VUS = 1000
const DURATION = '30s'

export const options = {
  vus: VUS,
  duration: DURATION
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
  return generateData(`${VUS}-vus-${DURATION}-create-card`, data)
}
