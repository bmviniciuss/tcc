import { textSummary } from 'https://jslib.k6.io/k6-summary/0.0.1/index.js'

const padNumber = (number, n = 2, filler = '0') => String(number).padStart(n, filler)
const format2D = (digit) => padNumber(digit)
const format3D = (digit) => padNumber(digit, 3)

export function getNowTimestamp () {
  const now = new Date()
  return `${now.getUTCFullYear()}${format2D(now.getUTCMonth() + 1)}${format2D(now.getDate())}${format2D(now.getHours())}${format2D(now.getMinutes())}${format3D(now.getMilliseconds())}`
}

export function getBenchmarkSummaryFileName (testName, isGRPC, basePath = "/home/bmviniciuss/Repos/tcc/benchmarks/benchmarks/results") {
  const mode = isGRPC ? 'grpc' : 'http'
  const timestamp = getNowTimestamp()
  return `${basePath}/${timestamp}-${testName}-${mode}.json`
}

export function generateData (fileName, testConfig, data) {
  const GENERATE_SUMMARY = __ENV.GENERATE_SUMMARY === 'true'
  const IS_GRPC = __ENV.GRPC_ENABLED === 'true'
  const type = IS_GRPC ? 'GRPC' : 'REST'

  if (!GENERATE_SUMMARY) {
    console.log('Not generating summary report for this benchmark')
    return {
      stdout: textSummary(data, { indent: ' ', enableColors: true })
    }
  }

  const summaryOutputFileName = getBenchmarkSummaryFileName(fileName, IS_GRPC, __ENV.OUT_PATH)

  data.metadata = { type, testConfig }
  return {
    stdout: textSummary(data, { indent: ' ', enableColors: true }),
    [`${summaryOutputFileName}`]: JSON.stringify(data)
  }
}
