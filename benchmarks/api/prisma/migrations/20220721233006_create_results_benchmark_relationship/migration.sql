/*
  Warnings:

  - Added the required column `benchmarkId` to the `Result` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `Result` ADD COLUMN `benchmarkId` VARCHAR(191) NOT NULL;
