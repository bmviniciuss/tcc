# TCC Cli
CLI to help develop TCC

## How to run
- Install dependencies
```bash
npm install
```
- Link CLI
```bash
npm link
```

## Commands
### Env
Change the GRPC_ENABLED env variable in the microsservices

To enable GRPC:
```bash
tcc-cli env --mode grpc <microsservices root folder>
```

To enable HTPP (Disable gRPC):
```bash
tcc-cli env --mode http <microsservices root folder>
```
