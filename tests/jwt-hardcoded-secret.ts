declare const jwt: { sign(...a: unknown[]): string; verify(...a: unknown[]): unknown };

// ruleid: jwt-hardcoded-secret
const token = jwt.sign({ sub: "42" }, "my-super-secret", { expiresIn: "1h" });

// ruleid: jwt-hardcoded-secret
jwt.verify(token, "my-super-secret");

// ok: jwt-hardcoded-secret
jwt.sign({ sub: "42" }, process.env.JWT_SECRET, { expiresIn: "1h" });
