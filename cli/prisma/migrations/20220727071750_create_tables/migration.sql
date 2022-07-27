-- CreateEnum
CREATE TYPE "Type" AS ENUM ('GRPC', 'REST');

-- CreateTable
CREATE TABLE "Benchmark" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "vus" INTEGER NOT NULL,
    "duration" TEXT,
    "iterations" INTEGER,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "Benchmark_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "Result" (
    "id" TEXT NOT NULL,
    "testDuration" DOUBLE PRECISION NOT NULL,
    "type" "Type" NOT NULL,
    "httpReqDurationMin" DOUBLE PRECISION NOT NULL,
    "httpReqDurationMax" DOUBLE PRECISION NOT NULL,
    "httpReqDurationAvg" DOUBLE PRECISION NOT NULL,
    "httpReqDurationMed" DOUBLE PRECISION NOT NULL,
    "iterationDurationMin" DOUBLE PRECISION NOT NULL,
    "iterationDurationMax" DOUBLE PRECISION NOT NULL,
    "iterationDurationAvg" DOUBLE PRECISION NOT NULL,
    "iterationDurationMed" DOUBLE PRECISION NOT NULL,
    "iterationsCount" INTEGER NOT NULL,
    "iterationsRate" DOUBLE PRECISION NOT NULL,
    "httpReqsCount" INTEGER NOT NULL,
    "httpReqsRate" DOUBLE PRECISION NOT NULL,
    "executedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,
    "benchmarkId" TEXT NOT NULL,

    CONSTRAINT "Result_pkey" PRIMARY KEY ("id")
);

-- AddForeignKey
ALTER TABLE "Result" ADD CONSTRAINT "Result_benchmarkId_fkey" FOREIGN KEY ("benchmarkId") REFERENCES "Benchmark"("id") ON DELETE RESTRICT ON UPDATE CASCADE;
