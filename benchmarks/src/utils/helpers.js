const padNumber = (number, n = 2, filler = '0') => String(number).padStart(n, filler)
const format2D = (digit) => padNumber(digit)
const format3D = (digit) => padNumber(digit, 3)

export function getNowTimestamp () {
  const now = new Date()
  return `${now.getUTCFullYear()}${format2D(now.getUTCMonth() + 1)}${format2D(now.getDate())}${format2D(now.getHours())}${format2D(now.getMinutes())}${format3D(now.getMilliseconds())}`
}

export function getBenchmarkSummaryFileName (testName, isGRPC) {
  const name = `src/${testName}/benchmarks`
  const mode = isGRPC ? 'grpc' : 'http'
  return `${name}-${mode}.json`
}

export function generateData (testName, data) {
  console.log('__ENV: ', JSON.stringify(__ENV, null, 2))
  const GENERATE_SUMMARY = __ENV.GENERATE_SUMMARY === 'true'
  if (!GENERATE_SUMMARY) {
    console.log('Not generating summary report for this benchmark')
    return {}
  }

  const IS_GRPC = __ENV.GRPC_ENABLED === 'true'
  const summaryOutputFileName = getBenchmarkSummaryFileName(testName, IS_GRPC)
  return {
    [`./${summaryOutputFileName}`]: JSON.stringify(data)
  }
}
