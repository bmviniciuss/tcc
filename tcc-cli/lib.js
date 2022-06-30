const path = require('path')
const fs = require('fs').promises
const { parse, stringify } = require('envfile')

function getEnvFilePath (rootPath, projectName) {
  return path.resolve(rootPath, projectName, '.env')
}

async function changeEnvForProject (path, mode) {
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

module.exports = {
  changeEnvForProject,
  getEnvFilePath
}
