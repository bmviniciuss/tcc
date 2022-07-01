import http from 'k6/http'
import { check } from 'k6'
import { generateData } from '/home/bmviniciuss/Repos/tcc/benchmarks/src/utils/helpers.js'

const GATEWAY_HOST = 'localhost:5000'

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
  return generateData('create-card', data)
}
