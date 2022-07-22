/*
  Warnings:

  - Added the required column `type` to the `Result` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE `Result` ADD COLUMN `type` ENUM('GRPC', 'REST') NOT NULL;
