# Veilcord Installer

The official installer for **Veilcord** — a privacy-first Discord client mod.
A small, portable Windows executable with two actions: **Inyectar** (install) and
**Restaurar** (uninstall).

It uses the same safe, file-based method as Vencord/Equicord: it renames Discord's
`app.asar` to `_app.asar` and adds a small `app/` loader that boots the original
Discord with Veilcord on top. Uninstalling reverses this exactly.

- No memory injection, no DLL injection — just files.
- No telemetry. It never sends your data anywhere.
- Build it yourself or let GitHub Actions build it. See **BUILD.md**.

## Status

First build: the bundled loader currently just confirms injection works and starts
Discord normally. The full Veilcord experience (plugins, themes, privacy hardening)
comes from the Veilcord mod (a fork of Equicord), which plugs into this loader next.

## Credits & license

Veilcord builds on the work of:
- **Vencord** by Vendicated and contributors — https://github.com/Vendicated/Vencord (GPL-3.0)
- **Equicord** — https://github.com/Equicord/Equicord (GPL-3.0)

The injection method follows the approach used by these projects.

Veilcord is **not affiliated with or endorsed by Discord Inc.** "Discord" is a
trademark of Discord Inc., used here only descriptively. Client modifications
technically violate Discord's Terms of Service.
