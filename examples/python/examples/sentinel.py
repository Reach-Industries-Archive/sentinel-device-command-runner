import os
import requests
import json
from datetime import datetime

# Class holding helper methods to interact with the SentinelEngine platform.


class Sentinel:
    def __init__(self, authKey):
        self.authkey = authKey
        self.baseurl = os.getenv("SENTINEL_URL")

    def pollCommands(self, deviceId):
        headers = {'Authorization': self.authkey, 'DeviceId': deviceId}
        commandsUrl = f'{self.baseurl}/device/commands'
        print(commandsUrl)
        response = requests.request("GET", commandsUrl, headers=headers)
        print(response)
        return json.loads(response.text)

    def uploadOutput(self, url, stdout, stderr):
        payload = json.dumps(
            {
                "stdout": stdout.decode("utf-8"),
                "stderr": stderr.decode("utf-8")
            })
        response = requests.request("PUT", url, data=payload)
