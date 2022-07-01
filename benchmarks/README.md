// TODO: write docs

```bash
k6 run src/create-card/k6-create-card.js -e GENERATE_SUMMARY=true -e GRPC_ENABLED=true --vus 1500 --duration '30s'
```
