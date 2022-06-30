#!/usr/bin/env node

const { Command } = require('commander')
const program = new Command()
const { getEnvFilePath, changeEnvForProject } = require('./lib')

const PROJECTS_NAME = ['card', 'card-payment', 'client-wallet', 'gateway-go']

program
  .name('tcc-cli')
  .version('1.0.0')

program.command('env')
  .description('Change the GRPC_ENABLE mode for microsservices')
  .argument('<string>', 'paths for microsservices')
  .option('--mode <string>', 'Microsservices execution mode. Must be \'grpc\' or \'http\'')
  .action(async (path, options) => {
    if (!options?.mode) {
      console.error("Options 'mode' must be provided")
      return
    }

    const { mode } = options
    const parsedMode = mode.trim().toLowerCase()
    const allowedValuesSet = new Set(['grpc', 'http'])

    if (!allowedValuesSet.has(parsedMode)) {
      console.error(`${mode} is not allowed. Must be one of the following [ ${Array.from(allowedValuesSet).join(', ')} ]`)
      return
    }

    const projectPaths = PROJECTS_NAME.map(projectName => getEnvFilePath(path, projectName))
    for (const envFilePath of projectPaths) {
      await changeEnvForProject(envFilePath, parsedMode)
    }
  })

program.parse()
