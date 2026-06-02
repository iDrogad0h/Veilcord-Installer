# Veilcord

**privacy-first vibes for your client.**

Veilcord is a privacy-focused Discord client mod by [sensurare
](https://github.com/iDrogad0h). It is a rebrand of [Equicord](https://github.com/Equicord/Equicord) — itself an enhanced fork of [Vencord](https://github.com/Vendicated/Vencord) bundling 300+ community plugins — tuned for a clean, quiet experience: no telemetry, trackers blocked out of the box, and crash reporting disabled by default.

## Features

- **Privacy by default** — no telemetry; Discord analytics and crash reporting blocked out of the box.
- **300+ plugins** from the Equicord/Vencord ecosystem.
- **Custom CSS and a built-in theme editor** — including support for BetterDiscord themes.
- **Works on any Discord branch** — Stable, PTB and Canary.
- **Lightweight** and quick to start.

## Installing Veilcord

### Windows (easiest)

Download the latest installer from the [Releases page](https://github.com/iDrogad0h/Veilcord-Installer/releases) and run it. Press **Inject**, then open Discord.

### Linux / macOS — quick install

```
wget https://raw.githubusercontent.com/iDrogad0h/Veilcord/refs/heads/main/misc/install.sh && chmod +x install.sh && ./install.sh
```

The script checks dependencies, clones the repo, builds, and injects into Discord automatically.

### Manual install (build from source)

Works on Windows, macOS, Linux and BSD — anything with Git, Node and the Discord desktop app. `pnpm inject` auto-detects your Discord install (Stable / PTB / Canary / Dev).

**Dependencies:** [Git](https://git-scm.com/download) and [Node.JS LTS](https://nodejs.org).

Install `pnpm`:

> ❗ This next command may need to be run as admin/root depending on your system, and you may need to close and reopen your terminal for `pnpm` to be in your PATH.

```
npm i -g pnpm
```

> ❗ IMPORTANT: From here onwards, make sure you are **not** using an admin/root terminal. It can break your Discord/Veilcord install and you would likely have to reinstall.

Clone Veilcord:

```
git clone https://github.com/iDrogad0h/Veilcord
cd Veilcord
```

Install dependencies:

```
pnpm install --frozen-lockfile
```

Build Veilcord:

```
pnpm build
```

Inject Veilcord into your desktop client:

```
pnpm inject
```

Build Veilcord for web:

```
pnpm buildWeb
```

After building the web extension, find the appropriate ZIP in the `dist` directory and follow your browser's guide for installing custom extensions, if supported. (Firefox's extension ZIP requires Firefox Developer Edition.)

## Support

Need help, want to report a bug, or just hang out? Join the Veilcord Discord server:

https://discord.gg/YOUR_INVITE

## Credits

Built on the shoulders of:

- [Vencord](https://github.com/Vendicated/Vencord) by [Vendicated](https://github.com/Vendicated) — the original Discord client mod.
- [Equicord](https://github.com/Equicord/Equicord) — the enhanced fork Veilcord is based on.

Huge thanks to everyone who contributed to those projects. Individual plugin authors are credited in their respective source files.

## Disclaimer

Discord is a trademark of Discord Inc., and is mentioned solely for descriptive purposes. This does not imply any affiliation with or endorsement by Discord Inc. Veilcord is not connected to Vencord or Equicord.

Using Veilcord, like any client modification, violates Discord's Terms of Service. Discord generally does not enforce this against client mods, but use it at your own risk and avoid any plugin with abusive behavior. If your account is critical to you, you may prefer not to use any client mod.

## License

[GPL-3.0-or-later](https://www.gnu.org/licenses/gpl-3.0.en.html). Veilcord is a fork of Equicord/Vencord; all upstream copyright and attribution is preserved (see `NOTICE.md`).
