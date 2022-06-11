-- CreateEnum
CREATE TYPE "transactionTypeEnum" AS ENUM ('CREDIT_CARD_PAYMENT', 'DEBIT_CARD_PAYMENT', 'WITHDRAWAL');

-- CreateEnum
CREATE TYPE "serviceTypeEnum" AS ENUM ('CARD_PAYMENT', 'INTERNAL');

-- CreateTable
CREATE TABLE "transactions" (
    "id" TEXT NOT NULL,
    "clientId" TEXT NOT NULL,
    "amount" DOUBLE PRECISION NOT NULL,
    "type" "transactionTypeEnum" NOT NULL,
    "transactionServiceId" TEXT,
    "service" "serviceTypeEnum" NOT NULL,
    "transactionDate" TIMESTAMP(3) NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "transactions_pkey" PRIMARY KEY ("id")
);
