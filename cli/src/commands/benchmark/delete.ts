import { Command } from '@oclif/core'
import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

export interface FileBenchmark {
  id: string
  name: string
  vus: number
  duration?: string
  iterations?: number
}

export default class BenchmarksDeleter extends Command {
  static description = 'Removes all benchmarks tests from the database '

  static examples = []
  static flags = {}
  static args = []

  async run (): Promise<void> {
    this.log('Removing all benchmarks tests from database')
    await prisma.benchmark.deleteMany({ where: {} })
    this.log('Done')
  }
}
