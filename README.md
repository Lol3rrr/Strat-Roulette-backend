# Strat-Roulette Backend

## Env-Variables
 * stratDB_URL
 * stratDB_PORT
 * stratDB_DATABASE
 * stratDB_COLLECTION
 
 * sessionDB_URL
 * sessionDB_PORT
 * sessionDB_DATABASE
 * sessionDB_COLLECTION
 
 * sessionDuration (in Minutes)

 * PORT

 * VAULT_URL: The base url for the Vault instance
 * APPROLE_ID: The ID of the Approle to use
 * APPROLE_SECRET: The Secret of the Approle to use

## Routes
 * GET /strat/random
 * GET /strat/single
 * POST /admin/login
 * GET /admin/strat/all
 * POST /admin/strat/add
 * POST /admin/strat/delete