import app from './app'
import { ENV } from './config/env'

async function main (): Promise<void> {
  try {
    await app.ready()
    app.swagger()
    await app.listen(ENV.PORT)
    console.log('Server listening on port: ', ENV.PORT)
  } catch (error) {
    console.error(error)
    app.log.error(error)
    process.exit(1)
  }
}

main()
