const pino = require('pino')
const axios = require('axios')

const logger = pino({
  transport: {
    target: 'pino-pretty'
  },
})

const benchmarks = [
  {
    "id": "6623f5bc-fa13-464c-ac10-ffc0f564fc05",
    "name": "create-card",
    "vus": 1000,
    "duration": "1m",
  },
  {
    "id": "b681f416-b0ca-46ca-b775-1b1d56aad45c",
    "name": "create-card",
    "vus": 750,
    "duration": "1m",
  },
  {
    "id": "9d9e7cc3-2820-4086-92c9-3a2e5addbb69",
    "name": "create-card",
    "vus": 500,
    "duration": "1m",
  },
  {
    "id": "652f8a51-30af-4b57-a5e5-f7e3eb5a3a4c",
    "name": "create-card",
    "vus": 250,
    "duration": "1m",
  },
  {
    "id": "ca5c45f1-0535-42e3-bf83-248a31c58eaa",
    "name": "create-card",
    "vus": 100,
    "duration": "1m",
  }
]

async function main() {
  for(const benchmark of benchmarks) {
    logger.info(`Processing Benchmark ${benchmark.id}`)
    await axios.post("http://localhost:3000/api/benchmarks", benchmark)
  }
}

main()
