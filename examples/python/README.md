# Sentinel Engine Python Examples

This repo contains runnable examples for sending data and uploading files to Sentinel Engine using Python. You should first have a [Sentinel Engine account](https://app.sentinelengine.ai/signup) and must have created an auth key at https://app.sentinelengine.ai/authkeys

## Getting Started

### Clone this Repo

```
git clone https://github.com/Reach-Industries/sentinel-python-example.git
```

### Configure environment variables

1. Create a `.env` file in root of the project
2. Add the following parameters to the env file in the format `KEY=VALUE`:
   - `SENTINEL_DEVICE_ID` - This should be the unique, human readable name for your device.
   - `SENTINEL_AUTH_KEY` - Use an auth key generated at app.sentinelengine.ai/authkeys
   - `CUSTOM_DOMAIN` - If using a premium account you should have a custom ingest domain url.
   
Your file should look something like this:

```
SENTINEL_DEVICE_ID=MyNewDevice
SENTINEL_AUTH_KEY=YourAuthKeyHere
CUSTOM_DOMAIN=https://http-decoder.XXXXX.sentinelengine.ai
```

### Install dependencies

*!! Note: where the `python` command is used - python 3 is assumed as the default version. This may be different for your system so make sure all commands are run with python 3   * 

Install [Pipenv](https://github.com/pypa/pipenv):

```
python -m pip install --user pipenv
```

Change your working directory to the project folder and install the required project dependencies with pipenv:

```
cd sentinel-python-example
```

```
python -m pipenv install
```

### Run example scripts using pipenv

You can now run the examples scripts in the example folder:

```
python -m pipenv run python ./examples/sendPayload.py
```

```
python -m pipenv run python ./examples/fileUpload.py
```

## View your data in Sentinel Engine

You should now see data in your Sentinel Engine account. You can view this by going to the Devices section and selecting the device ID configured in your repo
