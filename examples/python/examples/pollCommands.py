import json
import os
import time
from subprocess import Popen, PIPE, TimeoutExpired
from sentinel import Sentinel

# your Auth key generated from app.sentinelengine.ai/authkeys
AUTH_KEY = os.getenv("SENTINEL_AUTH_KEY")
# your Device ID
DEVICE_ID = os.getenv("SENTINEL_DEVICE_ID")

# Initialise a Sentinel API client
SentinelClient = Sentinel(AUTH_KEY)

while True:
    # 1. Get latest commands.
    commands = SentinelClient.pollCommands(DEVICE_ID)

    # 2. Execute commands.
    for command in commands:
        print("Executing command: ["+command['command']+"]")
        process = Popen(command['command'], stdout=PIPE, stderr=PIPE)
        try:
            # will raise error and kill any process that runs longer than 60 seconds
            stdout, stderr = process.communicate(timeout=10)
        except TimeoutExpired as e:
            process.kill()
            stdout, stderr = process.communicate()
        # 3. Upload command output.
        print("Uploading output...")
        SentinelClient.uploadOutput(command['upload'], stdout, stderr)

    print("No commands left to execute.")
    # 4. Remember to sleep before checking for new commands!
    time.sleep(60)


# Failing to add a pause before checking for new commands wait result in your device's requests getting throttled by SentinelEngine.
