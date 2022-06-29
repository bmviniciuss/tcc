#!/usr/bin/env node

const process = require('process')
const path = require('path')
const fs = require('fs').promises
const { parse, stringify } = require('envfile')

function getEnvFilePath () {
  const fileName = process.argv[2]
  const filePath = path.join(__dirname, fileName)
  return filePath
}

async function main () {
  if (!process?.argv?.[3]) {
    throw new Error('You must provide GRPC_ENABLED flag')
  }

  const grpcEnabledFlag = process?.argv?.[3] === 'true'
  const path = getEnvFilePath()
  console.log(`Setting ${path} to ${grpcEnabledFlag}`)

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

main()
