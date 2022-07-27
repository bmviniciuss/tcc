/* eslint-disable no-await-in-loop */
import {Command} from '@oclif/core'
import {OutputArgs} from '@oclif/core/lib/interfaces'
import * as path from 'node:path'
import * as fs from 'node:fs/promises'
import {PrismaClient} from '@prisma/client'
import {Type as TestType} from '.prisma/client'

const prisma = new PrismaClient()

export type FileResult = {
  options: {
    summaryTrendStats: Array<string>
    summaryTimeUnit: string
    noColor: boolean
  }
  state: {
    isStdOutTTY: boolean
    isStdErrTTY: boolean
    testRunDurationMs: number
  }
  metrics: {
    vus_max: {
      type: string
      contains: string
      values: {
        value: number
        min: number
        max: number
      }
    }
    http_req_sending: {
      type: string
      contains: string
      values: {
        max: number
        'p(90)': number
        'p(95)': number
        avg: number
        min: number
        med: number
      }
    }
    checks: {
      type: string
      contains: string
      values: {
        rate: number
        passes: number
        fails: number
      }
    }
    http_reqs: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    http_req_waiting: {
      type: string
      contains: string
      values: {
        avg: number
        min: number
        med: number
        max: number
        'p(90)': number
        'p(95)': number
      }
    }
    http_req_failed: {
      type: string
      contains: string
      values: {
        rate: number
        passes: number
        fails: number
      }
    }
    http_req_connecting: {
      type: string
      contains: string
      values: {
        med: number
        max: number
        'p(90)': number
        'p(95)': number
        avg: number
        min: number
      }
    }
    iteration_duration: {
      type: string
      contains: string
      values: {
        max: number
        'p(90)': number
        'p(95)': number
        avg: number
        min: number
        med: number
      }
    }
    iterations: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    http_req_duration: {
      type: string
      contains: string
      values: {
        med: number
        max: number
        'p(90)': number
        'p(95)': number
        avg: number
        min: number
      }
    }
    http_req_blocked: {
      type: string
      contains: string
      values: {
        med: number
        max: number
        'p(90)': number
        'p(95)': number
        avg: number
        min: number
      }
    }
    http_req_tls_handshaking: {
      type: string
      contains: string
      values: {
        avg: number
        min: number
        med: number
        max: number
        'p(90)': number
        'p(95)': number
      }
    }
    'http_req_duration{expected_response:true}': {
      type: string
      contains: string
      values: {
        avg: number
        min: number
        med: number
        max: number
        'p(90)': number
        'p(95)': number
      }
    }
    http_req_receiving: {
      type: string
      contains: string
      values: {
        avg: number
        min: number
        med: number
        max: number
        'p(90)': number
        'p(95)': number
      }
    }
    vus: {
      type: string
      contains: string
      values: {
        value: number
        min: number
        max: number
      }
    }
    data_sent: {
      type: string
      contains: string
      values: {
        count: number
        rate: number
      }
    }
    data_received: {
      contains: string
      values: {
        count: number
        rate: number
      }
      type: string
    }
  }
  metadata: {
    type: TestType
    testConfig: {
      id: string
      name: string
      vus: number
      duration: string
      executedAt: string
    }
  }
  root_group: {
    name: string
    path: string
    id: string
    groups: Array<any>
    checks: Array<{
      name: string
      path: string
      id: string
      passes: number
      fails: number
    }>
  }
}

export default class ResultsInjest extends Command {
  static description = 'Injets results from a folder'

  static examples = [
    '$ tcc-cli results injest ../benchmarks/benchmarks/api-injester/benchmarks-tests.json',
  ]

  static flags = {}

  static args = [
    {name: 'path', description: 'Path that contains results to be injests.', required: true},
  ]

  async run(): Promise<void> {
    const {args} = await this.parse(ResultsInjest)
    const resultsPath = this.getResultsPath(args)
    this.log(`Reading results from ${resultsPath}`)

    const files = await fs.readdir(resultsPath)
    this.log(`Found ${files.length} files`)

    let index = 0
    for (const fileName of files) {
      this.log(`Processing ${index + 1}/${files.length}: ${fileName}`)
      const filePath = path.join(resultsPath, fileName)
      const file = await fs.readFile(filePath, {encoding: 'utf8'})
      const content: FileResult = JSON.parse(file)
      await prisma.result.create({
        data: {
          fileName: fileName,
          type: content.metadata.type,
          testDuration: content.state.testRunDurationMs,
          httpReqDurationMin: content.metrics.http_req_duration.values.min,
          httpReqDurationMax: content.metrics.http_req_duration.values.max,
          httpReqDurationAvg: content.metrics.http_req_duration.values.avg,
          httpReqDurationMed: content.metrics.http_req_duration.values.med,
          iterationDurationMin: content.metrics.iteration_duration.values.min,
          iterationDurationMax: content.metrics.iteration_duration.values.max,
          iterationDurationAvg: content.metrics.iteration_duration.values.avg,
          iterationDurationMed: content.metrics.iteration_duration.values.med,
          iterationsCount: content.metrics.iterations.values.count,
          iterationsRate: content.metrics.iterations.values.rate,
          httpReqsCount: content.metrics.http_reqs.values.count,
          httpReqsRate: content.metrics.http_reqs.values.rate,
          executedAt: new Date(content.metadata.testConfig.executedAt),
          benchmark: {
            connect: {
              id: content.metadata.testConfig.id,
            },
          },
        },
      })
      index++
    }
  }

  private getResultsPath(args: OutputArgs): string {
    return path.resolve(process.cwd(), args.path)
  }
}
