/*
  Warnings:

  - A unique constraint covering the columns `[fileName]` on the table `Result` will be added. If there are existing duplicate values, this will fail.
  - Added the required column `fileName` to the `Result` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "Result" ADD COLUMN     "fileName" TEXT NOT NULL;

-- CreateIndex
CREATE UNIQUE INDEX "Result_fileName_key" ON "Result"("fileName");
