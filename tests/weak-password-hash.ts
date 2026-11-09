import crypto from "node:crypto";

// ruleid: weak-password-hash
const weak = crypto.createHash("md5").update("x").digest("hex");

// ok: weak-password-hash
const strong = crypto.createHash("sha256").update("x").digest("hex");

export { weak, strong };
