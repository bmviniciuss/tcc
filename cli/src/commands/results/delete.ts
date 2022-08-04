import { Command, CliUx } from '@oclif/core'
import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

export default class ResultsDeleter extends Command {
  static description = 'Injets results from a folder'

  static examples = [
    '$ tcc-cli results delete'
  ]

  static flags = {}

  static args = []

  async run (): Promise<void> {
    CliUx.ux.action.start('Removing all results from database')
    await prisma.result.deleteMany({ where: {} })
    CliUx.ux.action.stop('Done')
  }
}
