# sentinel-device-command-runner

Repository used to hold examples in various languages for using the remote device command runner functionality.

## Examples

The examples provided are scripts that should poll for pending commands, it will then execute and return the results to SentinelEngine.

### NodeJS

- Navigate to `examples/nodejs`.
- Create `.env` file with required configuration values (default values can be found in .env.example):

  - `SENTINEL_URL` - Standard SE accounts should use the default, otherwise `https://http-decoder.SUBDOMAIN.sentinelengine.ai`
  - `SENTINEL_FILE_FOLDER` - the folder you want to be synchronised with the remote files.
  - `SENTINEL_DEVICE_ID` - The deviceId to be synchronised.
  - `SENTINEL_DEVICE_AUTH_KEY` - A valid AuthKey for your account.

- Install dependencies:

```
npm i
```

- Run script:

```
node index.js
```

### Python3

- Navigate to `examples/python`.
- Create `.env` file with required configuration values:

  - `SENTINEL_URL` - Standard SE accounts should use the default, otherwise `https://http-decoder.SUBDOMAIN.sentinelengine.ai`
  - `SENTINEL_DEVICE_ID` - The deviceId to be synchronised.
  - `SENTINEL_DEVICE_AUTH_KEY` - A valid AuthKey for your account.

- Install [Pipenv](https://github.com/pypa/pipenv):

```
python -m pip install --user pipenv
```

- Install dependencies:

```
python -m pipenv install
```

- Run script:

```
python -m pipenv run ./pollCommands.py
```

## Troubleshooting

Below are some common issue that users can run into when setting up the examples.

### The call to the /device/commands endpoint keeps timing out.

Ensure that the SENTINEL_URL is correct, be careful to check the url begins with `https` and not `http`.
