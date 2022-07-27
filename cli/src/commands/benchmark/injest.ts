import {Command} from '@oclif/core'
import { OutputArgs } from '@oclif/core/lib/interfaces'
import * as path from "node:path"

export default class BenchmarksInjest extends Command {
  static description = 'Say hello world'

  static examples = [
    `$ oex hello world
hello world! (./src/commands/hello/world.ts)
`,
  ]

  static flags = {}

  static args = [
    { name: "file", description: "JSON file that contains a list of benchmarks tests to injest to the database", required: true}
  ]

  async run(): Promise<void> {
    const {args} = await this.parse(BenchmarksInjest)
    const filePath = this.getFilePath(args)
    this.log(`Reading file from ${filePath}`)
  }

  private getFilePath(args: OutputArgs): string {
    const rawFilePath: string = args.file
    return path.resolve(process.cwd(), rawFilePath)
  }
}
