import https from "node:https";

// ruleid: tls-verification-disabled
const insecureAgent = new https.Agent({ keepAlive: true, rejectUnauthorized: false });

// ruleid: tls-verification-disabled
process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";

// ok: tls-verification-disabled
const agent = new https.Agent({ keepAlive: true, rejectUnauthorized: true });

export { insecureAgent, agent };
