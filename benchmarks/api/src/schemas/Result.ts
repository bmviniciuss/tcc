import { Type } from '@sinclair/typebox'

export const MMAM = Type.Object({
  min: Type.Number(),
  max: Type.Number(),
  avg: Type.Number(),
  med: Type.Number()
})

export const CountRate = Type.Object({
  count: Type.Integer(),
  rate: Type.Number()
})

export enum CommunicationType {
  GRPC = 'gRPC',
  REST = 'REST'
}

const ResultSchema = Type.Object({
  metadata: Type.Object({
    type: Type.Enum(CommunicationType),
    testConfig: Type.Object({
      id: Type.String(),
      name: Type.String(),
      vus: Type.Integer({ minimum: 0 }),
      duration: Type.Optional(Type.String()),
      iterations: Type.Optional(Type.Integer({ minimum: 0 })),
      executedAt: Type.String({ format: 'date-time' })
    })
  }),
  state: Type.Object({
    testRunDurationMs: Type.Number()
  }),
  metrics: Type.Object({
    http_req_duration: Type.Object({
      values: MMAM
    }),
    iteration_duration: Type.Object({
      values: MMAM
    }),
    iterations: Type.Object({
      values: CountRate
    }),
    http_reqs: Type.Object({
      values: CountRate
    })
  })
})

export default ResultSchema
