const fs = require('node:fs/promises')
const ffs = require('fs')
const path = require('node:path')

const logger = require('pino')({
  transport: {
    target: 'pino-pretty',
    options: {
      colorize: true
    }
  }
})

async function main () {
  const DELIMITER = ';'
  const outputFile = 'out.csv'
  const resultsPath = path.resolve('./src/results')
  logger.info(`Reading results path = ${resultsPath}`)

  const files = await fs.readdir(resultsPath)
  logger.info(`${files.length} files found in ${resultsPath}`)

  const writeStream = ffs.createWriteStream(outputFile)

  const fields = {
    testName: {
      header: 'Teste',
      key: 'metadata.testName'
    },
    type: {
      header: 'Tipo',
      key: 'metadata.type'
    },
    testDurationMs: {
      header: 'Duração Teste (ms)',
      key: 'state.testRunDurationMs'
    },
    httpReqsCount: {
      header: 'Número Requisições',
      key: 'metrics.http_reqs.values.count'
    },
    httpReqsRate: {
      header: 'Vazão  (reqs/s)',
      key: 'metrics.http_reqs.values.rate'
    },
    httpReqDurationMin: {
      header: 'Tempo requisição (min)',
      key: 'metrics.http_req_duration.values.min'
    },
    httpReqDurationMax: {
      header: 'Tempo requisição (max)',
      key: 'metrics.http_req_duration.values.max'
    },
    httpReqDurationAvg: {
      header: 'Tempo requisição (média)',
      key: 'metrics.http_req_duration.values.avg'
    }
  }

  const rowConfig = [
    fields.testName,
    fields.type,
    fields.testDurationMs,
    fields.httpReqsCount,
    fields.httpReqsRate,
    fields.httpReqDurationMin,
    fields.httpReqDurationMax,
    fields.httpReqDurationAvg
  ]
  const headerRow = rowConfig.map(config => config.header).join(DELIMITER)
  await writeStream.write(headerRow + '\n')

  function accessField (obj, fieldKey) {
    const keys = fieldKey.split('.')
    return keys.reduce((acc, key) => acc[key], obj)
  }

  for (const file of files) {
    logger.info(`Processing file: ${file}`)
    const rawFileData = await fs.readFile(path.join(resultsPath, file))
    const fileData = JSON.parse(rawFileData)
    const row = rowConfig.map(rowConfig => {
      return accessField(fileData, rowConfig.key)
    }).join(DELIMITER) + '\n'
    await writeStream.write(row)
  }

  writeStream.end()
}

main()
