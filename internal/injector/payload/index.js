/*
 * Veilcord — minimal loader (first/test build).
 *
 * This is intentionally tiny and SAFE: it does not touch Discord's preload or
 * its UI, so Discord starts exactly as normal. All it does is:
 *   1. drop a small proof-of-injection marker file, and
 *   2. re-point Electron at the original Discord app (_app.asar) and start it.
 *
 * Once injection is confirmed working, this whole folder is replaced by the
 * real Veilcord build (the Equicord fork's dist/), which is where plugins,
 * themes, privacy features, etc. live.
 */

const { join, dirname } = require("path");
const electron = require("electron");

// .../resources/app/index.js  ->  .../resources
const here = dirname(require.main.filename);
const realAsar = join(here, "..", "_app.asar");
const realPkg = require(join(realAsar, "package.json"));

// 1) proof of injection (best-effort; never blocks startup)
try {
    const fs = require("fs");
    const os = require("os");
    fs.writeFileSync(
        join(os.tmpdir(), "veilcord-injected.txt"),
        "Veilcord active\n" + new Date().toISOString() + "\nbranch: " + (realPkg.name || "discord") + "\n"
    );
} catch (e) {
    // ignore — the marker is only a convenience
}
console.log("[Veilcord] loader active — starting Discord");

// 2) hand control back to the real Discord app
electron.app.setAppPath(realAsar);
require.main.filename = join(realAsar, realPkg.main);
require(require.main.filename);
