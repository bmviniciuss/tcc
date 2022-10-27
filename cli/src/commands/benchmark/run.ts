import { Command, Flags } from '@oclif/core'
import { OutputArgs } from '@oclif/core/lib/interfaces'
import * as path from 'node:path'
import * as fs from 'node:fs/promises'
import * as cp from 'node:child_process'
import { format } from 'date-fns'

export interface FileBenchmark {
  id: string
  name: string
  vus: number
  duration?: string
  iterations?: number
}

export default class BenchmarksInjest extends Command {
  static description = 'Injets benchmarks tests to the database'

  static examples = [
    '$ tcc-cli benchmark injest ../benchmarks/benchmarks/api-injester/benchmarks-tests.json'
  ]

  static flags = {
    grpc: Flags.boolean({}),
    rest: Flags.boolean({}),
    createCard: Flags.boolean({}),
    createCardPayment: Flags.boolean({})
  }

  static args = [
    { name: 'cenariosPath', description: 'cenarios path', required: true },
    { name: 'composePath', description: 'docker compose path', required: true },
    { name: 'outPath', description: 'docker compose path', required: true }
  ]

  async run (): Promise<void> {
    const { args, flags } = await this.parse(BenchmarksInjest)
    const values = this.getGrpcEnabledValue(flags)

    const filesPath = this.getFilePath(args)

    const files = await (await fs.readdir(filesPath)).filter(file => {
      let shouldBeExecuted = false

      if (flags?.createCard) {
        shouldBeExecuted = file.includes('create-card_')
      }

      if (flags?.createCardPayment) {
        shouldBeExecuted = shouldBeExecuted || file.includes('create-card-payment')
      }

      return shouldBeExecuted
    })

    const composePath = path.resolve(process.cwd(), args.composePath)
    const outDate = format(new Date(), 'yyyy-MM-dd-HH-mm')
    const outPath = path.resolve(process.cwd(), args.outPath, outDate)

    await fs.mkdir(outPath).catch(() => {
      this.log('Directory already exists')
      throw new Error('Directory already exists')
    })

    console.log({
      args, flags, values, files, composePath, outPath
    })

    for (const GRPC_ENABLED of values) {
      cp.spawnSync('docker compose -f ../docker-compose.yml down -v', { stdio: 'inherit', shell: true })
      cp.spawnSync('sleep 2', { stdio: 'inherit', shell: true })
      for (const file of files) {
        this.log(`Processing file ${file} with GRPC_ENABLED=${GRPC_ENABLED}`)
        const filePath = path.join(filesPath, file)
        console.log(filePath)
        cp.spawnSync(`GRPC_ENABLED=${GRPC_ENABLED} docker compose -f ../docker-compose.yml up -d`, { stdio: 'inherit', shell: true })
        cp.spawnSync('sleep 5', { stdio: 'inherit', shell: true })
        cp.spawnSync(`k6 run ${filePath} -e GENERATE_SUMMARY=true -e GRPC_ENABLED=${GRPC_ENABLED} -e OUT_PATH=${outPath}`, { stdio: 'inherit', shell: true })
        cp.spawnSync('echo', { stdio: 'inherit', shell: true })
        cp.spawnSync('sleep 5', { stdio: 'inherit', shell: true })
        cp.spawnSync('docker compose -f ../docker-compose.yml down -v', { stdio: 'inherit', shell: true })
        cp.spawnSync('sleep 5', { stdio: 'inherit', shell: true })
      }
    }
  }

  private getGrpcEnabledValue (flags: { grpc: boolean; rest: boolean } & { [flag: string]: any } & { json: boolean | undefined }) {
    let values: boolean[] = []

    if (!flags?.grpc && !flags?.rest) {
      values = [false, true]
    }

    if (flags?.rest) {
      values = [...values, false]
    }

    if (flags?.grpc) {
      values = [...values, true]
    }

    return values
  }

  private getFilePath (args: OutputArgs): string {
    const rawFilePath: string = args.cenariosPath
    return path.resolve(process.cwd(), rawFilePath)
  }
}
