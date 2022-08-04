import { Command, Flags } from '@oclif/core'
import * as fs from 'node:fs/promises'
import * as path from 'node:path'
import { parse, stringify } from 'envfile'

export default class EnvModifier extends Command {
  static description = 'Change the microsservices mode'

  static examples = [
    '$ tcc-cli env ../ --mode grpc'
  ]

  static flags = {
    mode: Flags.enum({ options: ['grpc', 'http'] })
  }

  static args = [
    { name: 'path', description: 'Path that microsservices folders', required: true }
  ]

  async run (): Promise<void> {
    const { args, flags } = await this.parse(EnvModifier)

    const allowedValuesSet = new Set(['grpc', 'http'])
    const mode: 'grpc' | 'http' = flags.mode
    const { path: rootPath } = args

    if (!allowedValuesSet.has(mode)) {
      console.error(`${mode} is not allowed. Must be one of the following [ ${[...allowedValuesSet].join(', ')} ]`)
    }

    const PROJECTS_NAME = ['card', 'card-payment', 'gateway-go']
    const projectPaths = PROJECTS_NAME.map(projectName => this.getEnvFilePath(rootPath, projectName))

    for (const envFilePath of projectPaths) {
      // eslint-disable-next-line no-await-in-loop
      await this.changeEnvForProject(envFilePath, mode)
    }
  }

  private getEnvFilePath (rootPath: string, projectName: string) {
    return path.resolve(rootPath, projectName, '.env')
  }

  private async changeEnvForProject (path: string, mode: 'grpc' | 'http') {
    const grpcEnabledFlag = mode === 'grpc'
    console.log(`Setting ${path} to ${grpcEnabledFlag ? 'GRPC' : 'HTPP'}`)

    const content = await fs.readFile(path, {
      encoding: 'utf-8',
      flag: 'r'
    })

    const envVariables = parse(content)

    if (String(envVariables.GRPC_ENABLED) === String(grpcEnabledFlag)) {
      return
    }

    envVariables.GRPC_ENABLED = String(grpcEnabledFlag)
    const strContent = stringify(envVariables)
    await fs.writeFile(path, strContent)
  }
}
