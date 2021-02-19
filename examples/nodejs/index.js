import axios from "axios";
import { promisify } from "util";
import { exec } from "child_process";
import {
  SENTINEL_DEVICE_AUTH_KEY,
  SENTINEL_DEVICE_ID,
  SENTINEL_URL,
  SENTINEL_POLL_FREQUENCY,
} from "./config.js";
const execAsync = promisify(exec);

(async () => {
  console.log("Running SentinelEngine Device File Manager...");
  const headers = {
    authorization: SENTINEL_DEVICE_AUTH_KEY,
    deviceid: SENTINEL_DEVICE_ID,
  };

  while (true) {
    try {
      console.log("Get pending commands...");
      const pendingCommands = (
        await axios.get(`${SENTINEL_URL}/device/commands`, { headers })
      ).data;

      if (pendingCommands.length < 1) {
        console.log("No commands.");
      } else {
        console.log("Executing new commands...");
        await Promise.all(
          pendingCommands.map(async (cmd) => {
            console.log(`Running command: [${cmd.command}]...`);
            const { stdout, stderr } = await execAsync(cmd.command);
            console.log("stdout:", stdout);
            console.error("stderr:", stderr);
            console.log("Uploading output...");
            const result = await axios.put(cmd.upload, { stdout, stderr });
            console.log(result.message);
          })
        );
      }
      await new Promise((r) => setTimeout(r, SENTINEL_POLL_FREQUENCY));
    } catch (err) {
      // Log Error
      console.log("Oops something is wrong:");
      if (err.isAxiosError) {
        console.error(err?.message);
        console.error(`URL: ${err?.config?.url}`);
        console.error(
          `Headers: ${JSON.stringify(err?.response?.headers, null, 2)}`
        );
      } else {
        console.error(err);
      }

      // Error cool-off
      await new Promise((r) =>
        setTimeout(r, SENTINEL_POLL_FREQUENCY * 2)
      );
    }
  }
})();
