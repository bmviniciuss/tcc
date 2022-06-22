import 'dotenv/config'

export const ENV = {
  PORT: Number(process.env?.PORT ?? '4000'),
  ENABLE_GRPC: process.env?.ENABLE_GRPC === 'true',
  CARD_GRPC_HOST: process.env?.CARD_GRPC_HOST ?? 'localhost:3333',
  CARD_PAYMENT_GRPC_HOST: process.env?.CARD_PAYMENT_GRPC_HOST ?? 'localhost:5555',
  CLIENT_WALLET_HOST: process.env?.CLIENT_WALLET_HOST ?? 'localhost:5004'
}
