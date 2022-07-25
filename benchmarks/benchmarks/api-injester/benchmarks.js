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
  },
  {
    "id": "240a7dee-fed8-4fa0-b279-a79cdad6331d",
    "name": "create-card-payment",
    "vus": 100,
    "duration": "1m",
  },
  {
    "id": "c4e77cc4-0103-4abc-ac9e-da0f78901f4e",
    "name": "create-card-payment",
    "vus": 50,
    "duration": "1m",
  },
  {
    "id": "43631c23-72c8-4ce4-ae7d-b50cbb5e56e7",
    "name": "create-card-payment",
    "vus": 25,
    "duration": "1m",
  },
  {
    "id": "8ad8795d-8995-4bc5-b526-1914eddaf994",
    "name": "create-card-payment",
    "vus": 10,
    "duration": "1m",
  },
  {
    "id": "dad9c90f-2411-4200-9356-eeb9588680e0",
    "name": "create-card-payment",
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
