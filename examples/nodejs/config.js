import dotenv from "dotenv";
dotenv.config();

export const SENTINEL_URL = process.env.SENTINEL_URL;
export const SENTINEL_FILE_FOLDER =
  process.env.SENTINEL_FILE_FOLDER;
export const SENTINEL_DEVICE_ID = process.env.SENTINEL_DEVICE_ID;
export const SENTINEL_DEVICE_AUTH_KEY =
  process.env.SENTINEL_DEVICE_AUTH_KEY;
export const SENTINEL_POLL_FREQUENCY = 5000;
