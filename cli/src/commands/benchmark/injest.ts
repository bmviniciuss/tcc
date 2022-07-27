import {Command, CliUx} from '@oclif/core'
import {OutputArgs} from '@oclif/core/lib/interfaces'
import * as path from 'node:path'
import * as fs from 'node:fs/promises'
import {PrismaClient} from '@prisma/client'
import {Prisma} from '.prisma/client'

const prisma = new PrismaClient()

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
    '$ tcc-cli benchmark injest ../benchmarks/benchmarks/api-injester/benchmarks-tests.json',
  ]

  static flags = {}

  static args = [
    {name: 'file', description: 'Path to a JSON file that contains a list of benchmarks tests to injest to the database', required: true},
  ]

  async run(): Promise<void> {
    const {args} = await this.parse(BenchmarksInjest)
    const filePath = this.getFilePath(args)
    this.log(`Reading file from ${filePath}`)
    const file = await fs.readFile(filePath, {encoding: 'utf8'})
    const content: FileBenchmark[] = JSON.parse(file)
    this.log(`${content.length} Benchmarks test were found.`)

    const createManyData: Prisma.BenchmarkCreateManyInput[] = content.map(b => ({
      id: b.id,
      name: b.name,
      vus: b.vus,
      duration: b?.duration ?? null,
      iterations: b?.iterations ?? null,
    }))

    CliUx.ux.action.start('Inserting benchmarks tests to the database')

    await prisma.benchmark.createMany({
      data: createManyData,
    })
    CliUx.ux.action.stop('Done')
  }

  private getFilePath(args: OutputArgs): string {
    const rawFilePath: string = args.file
    return path.resolve(process.cwd(), rawFilePath)
  }
}
