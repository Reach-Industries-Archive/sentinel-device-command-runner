# sentinel-device-command-runner

Repository used to hold examples in various languages for using the remote device command runner functionality.

## Examples

The examples provided are scripts that should poll for pending commands, it will then execute and return the results to SentinelEngine.

### NodeJS

- Navigate to examples/nodejs
- Create .env file with required configuration values (default values can be found in .env.example):

  - `SENTINEL_ENGINE_URL` - Standard SE accounts should use the default, otherwise `https://http-decoder.SUBDOMAIN.sentinelengine.ai`
  - `SENTINEL_ENGINE_FILE_FOLDER` - the folder you want to be synchronised with the remote files.
  - `SENTINEL_ENGINE_DEVICE_ID` - The deviceId to be synchronised.
  - `SENTINEL_ENGINE_DEVICE_AUTH_KEY` - A valid AuthKey for your account.

- Install dependencies: `npm i`
- Run script: `node index.js`
