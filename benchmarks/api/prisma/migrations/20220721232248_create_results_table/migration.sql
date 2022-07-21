-- CreateTable
CREATE TABLE `Result` (
    `id` VARCHAR(191) NOT NULL,
    `testDuration` DOUBLE NOT NULL,
    `httpReqDurationMin` DOUBLE NOT NULL,
    `httpReqDurationMax` DOUBLE NOT NULL,
    `httpReqDurationAvg` DOUBLE NOT NULL,
    `httpReqDurationMed` DOUBLE NOT NULL,
    `iterationDurationMin` DOUBLE NOT NULL,
    `iterationDurationMax` DOUBLE NOT NULL,
    `iterationDurationAvg` DOUBLE NOT NULL,
    `iterationDurationMed` DOUBLE NOT NULL,
    `iterationsCount` BIGINT UNSIGNED NOT NULL,
    `iterationsRate` DOUBLE NOT NULL,
    `httpReqsCount` BIGINT UNSIGNED NOT NULL,
    `httpReqsRate` DOUBLE NOT NULL,
    `executedAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `createdAt` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    `updatedAt` DATETIME(3) NOT NULL,

    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
