const pino = require('pino')
const axios = require('axios')
const fsp = require("node:fs/promises")
const path = require('node:path')

const logger = pino({
  transport: {
    target: 'pino-pretty'
  },
})

async function main() {
  logger.info("API-Injester process started")
  const fileRootPath = path.resolve('..', 'results', 'create-card-3')
  const files = await fsp.readdir(fileRootPath)

  logger.info(`Found ${files.length} files`);

  let index = 0;
  for (const fileName of files) {
    logger.info(`Processing ${index + 1}/${files.length}: ${fileName}`)
    const filePath = path.join(fileRootPath, fileName)
    const file = await fsp.readFile(filePath)
    const content = JSON.parse(file)
    await axios.post("http://localhost:3000/api/results", content)
    index++
  }
  
} 

main()
