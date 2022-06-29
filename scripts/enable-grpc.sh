#!/bin/bash
node ../env-modifier/index.js ../card/.env true
node ../env-modifier/index.js ../card-payment/.env true
node ../env-modifier/index.js ../client-wallet/.env true
node ../env-modifier/index.js ../gateway/.env true
