import logger from '../utils/logger'
import app from './app'
import { ENV } from './config/env'

async function main (): Promise<void> {
  try {
    await app.listen(ENV.PORT)
    logger.info(`Gateway server started at port ${ENV.PORT}`)
  } catch (error) {
    app.log.error(error)
    process.exit(1)
  }
}

main()
