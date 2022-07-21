/*
  Warnings:

  - You are about to alter the column `iterationsCount` on the `Result` table. The data in that column could be lost. The data in that column will be cast from `UnsignedBigInt` to `UnsignedInt`.
  - You are about to alter the column `httpReqsCount` on the `Result` table. The data in that column could be lost. The data in that column will be cast from `UnsignedBigInt` to `UnsignedInt`.

*/
-- AlterTable
ALTER TABLE `Result` MODIFY `iterationsCount` INTEGER UNSIGNED NOT NULL,
    MODIFY `httpReqsCount` INTEGER UNSIGNED NOT NULL;
