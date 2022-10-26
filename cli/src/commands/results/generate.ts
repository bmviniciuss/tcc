/* eslint-disable no-await-in-loop */
import { Command } from '@oclif/core'
import { OutputArgs } from '@oclif/core/lib/interfaces'
import * as path from 'node:path'
import * as fs from 'node:fs/promises'
import { Result } from '../../types/result'
import { benchmarks } from '../../benchmarks'
import { parse, transforms } from 'json2csv'

export default class ResultsInjest extends Command {
  static description = 'Injets results from a folder'

  static examples = [
    '$ tcc-cli results injest ../benchmarks/benchmarks/api-injester/benchmarks-tests.json'
  ]

  static flags = {}

  static args = [
    { name: 'path', description: 'Path that contains results to be injests.', required: true },
    { name: 'outPath', description: 'destination file', required: true }
  ]

  async run (): Promise<void> {
    const benchmarksMap = new Map(benchmarks.map((b) => [b.id, b]))
    const { args } = await this.parse(ResultsInjest)
    const resultsPath = this.getResultsPath(args)
    this.log(`Reading results from ${resultsPath}`)

    const files = await fs.readdir(resultsPath)
    this.log(`Found ${files.length} files`)
    const data = []

    let index = 0
    for (const fileName of files) {
      this.log(`Processing ${index + 1}/${files.length}: ${fileName}`)
      const filePath = path.join(resultsPath, fileName)
      const file = await fs.readFile(filePath, { encoding: 'utf8' })
      const content: Result = JSON.parse(file)
      const benchmark = benchmarksMap.get(content.metadata.testConfig.id)

      if (!benchmark) throw new Error('No Benchmark found')

      data.push({
        fileName,
        type: content.metadata.type,
        testDuration: content.state.testRunDurationMs,
        http_req_duration: {
          min: content.metrics.http_req_duration.values.min,
          max: content.metrics.http_req_duration.values.max,
          avg: content.metrics.http_req_duration.values.avg,
          med: content.metrics.http_req_duration.values.med,
          p90: content.metrics.http_req_duration.values['p(90)'],
          p95: content.metrics.http_req_duration.values['p(95)'],
          p99: content.metrics.http_req_duration.values['p(99)']
        },
        http_reqs_count: content.metrics.http_reqs.values.count,
        http_reqs_rate: content.metrics.http_reqs.values.rate,

        iteration_duration: {
          min: content.metrics.iteration_duration.values.min,
          max: content.metrics.iteration_duration.values.max,
          avg: content.metrics.iteration_duration.values.avg,
          med: content.metrics.iteration_duration.values.med,
          p90: content.metrics.iteration_duration.values['p(90)'],
          p95: content.metrics.iteration_duration.values['p(95)'],
          p99: content.metrics.iteration_duration.values['p(99)']

        },
        iterations_count: content.metrics.iterations.values.count,
        iterations_rate: content.metrics.iterations.values.rate,
        benchmark,
        executedAt: new Date(content.metadata.testConfig.executedAt)

      })
      index++
    }

    const csv = parse(data, {
      transforms: [transforms.flatten({ separator: '_' })]
    })
    const absPath = path.resolve(process.cwd(), args.outPath)
    await fs.writeFile(absPath, csv, { encoding: 'utf8' })
  }

  private getResultsPath (args: OutputArgs): string {
    return path.resolve(process.cwd(), args.path)
  }
}
