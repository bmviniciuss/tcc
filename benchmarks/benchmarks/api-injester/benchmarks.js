const pino = require('pino')
const axios = require('axios')

const logger = pino({
  transport: {
    target: 'pino-pretty'
  },
})

const benchmarks = [
  {
    "id": "f889ac39-6948-4782-afa0-93f678ad1a98",
    "name": "create-card",
    "vus": 100,
    "duration": "1m",
  },
  {
    "id": "14db1832-359c-4226-a0e5-59ee9399888a",
    "name": "create-card",
    "vus": 50,
    "duration": "1m",
  },
  {
    "id": "f31c8249-3080-4a20-9c51-1a8d4fd982c8",
    "name": "create-card",
    "vus": 25,
    "duration": "1m",
  },
  {
    "id": "dfce9f52-4c55-4845-bbc2-5ec8279ccf6e",
    "name": "create-card",
    "vus": 10,
    "duration": "1m",
  },
  {
    "id": "8244654e-51ff-481d-bd43-bbc86715a505",
    "name": "create-card",
    "vus": 1,
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
