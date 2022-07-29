import {Command} from '@oclif/core'
import {PrismaClient} from '@prisma/client'
import {parse} from 'json2csv'
import * as fs from 'node:fs/promises'
import path = require('path')

const prisma = new PrismaClient()

type TestResult = {
  vus: number
  type: string
  name: string
  http_req_duration_min: number
  http_req_duration_max: number
  http_req_duration_avg: number
  http_req_duration_med: number
  http_reqs_count: number
  http_reqs_rate: number
}

export default class ResultsCSVGEnerator extends Command {
  static description = 'Generate results csv to analysis'

  static examples = [
    '$ tcc-cli results csv-generator',
  ]

  static flags = {}

  static args = [
    {name: 'path', description: 'Path to the file that will be generated', required: true},
  ]

  async run(): Promise<void> {
    const {args} = await this.parse(ResultsCSVGEnerator)
    // CliUx.ux.action.start('Removing all results from database')
    const res: TestResult[] = await prisma.$queryRaw`
          SELECT b.vus, type, name,
            avg("httpReqDurationMin") as http_req_duration_min,
            avg("httpReqDurationMax") as http_req_duration_max,
            avg("httpReqDurationAvg") as http_req_duration_avg,
            avg("httpReqDurationMed") as http_req_duration_med,
            avg("httpReqsCount") as http_reqs_count,
            avg("httpReqsRate") as http_reqs_rate
      FROM public."Result" r
      JOIN public."Benchmark" b ON r."benchmarkId" = b.id
      GROUP BY type, vus, name
    `

    const fields = [
      'vus',
      'type',
      'name',
      'http_req_duration_min',
      'http_req_duration_max',
      'http_req_duration_avg',
      'http_req_duration_med',
      'http_reqs_count',
      'http_reqs_rate',
    ]
    const csv = parse(res, {fields})
    const absPath = path.resolve(process.cwd(), args.path)
    await fs.writeFile(absPath, csv, {encoding: 'utf8'})
  }
}
