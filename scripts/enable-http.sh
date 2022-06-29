#!/bin/bash
node ../env-modifier/index.js ../card/.env false
node ../env-modifier/index.js ../card-payment/.env false
node ../env-modifier/index.js ../client-wallet/.env false
node ../env-modifier/index.js ../gateway/.env false
