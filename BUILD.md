# Veilcord Installer — cómo sacar el .exe

Hay dos caminos. El **fácil** (GitHub te lo compila) y el **manual** (lo compilás vos).

---

## Camino fácil (recomendado): que GitHub haga el .exe

No necesitás instalar nada en tu PC.

1. Creá un repositorio en GitHub (por ejemplo `Veilcord-Installer`).
2. Subí **toda esta carpeta** (`installer/`) a ese repositorio.
3. Andá a la pestaña **Actions** del repositorio. Vas a ver que se ejecuta solo.
4. Cuando termine (tarda 1-2 minutos), entrá al trabajo que corrió y bajá el archivo **`Veilcord-windows`**. Adentro está:
   - **`Veilcord.exe`** → el de los dos botones (Inyectar / Restaurar).
   - **`Veilcord-cli.exe`** → el de consola (por las dudas).

Para hacer una **Release** pública (la página de descarga, como Vencord):
- Creá un *tag* que empiece con `v`, por ejemplo `v1.0.0`. GitHub publica el `.exe` solo en la sección **Releases**.

---

## Camino manual (en tu PC con Windows)

1. Instalá Go: https://go.dev/dl  (apretá Download, instalá, listo).
2. Abrí una terminal (PowerShell) dentro de la carpeta `installer/`.
3. El de los dos botones:
   ```powershell
   go get github.com/lxn/walk github.com/lxn/win
   go build -tags gui -ldflags "-H windowsgui -s -w" -o Veilcord.exe .
   ```
4. (Opcional) El de consola:
   ```powershell
   go build -ldflags "-s -w" -o Veilcord-cli.exe .
   ```

Te queda `Veilcord.exe` ahí mismo. Es portable: lo podés mover a donde quieras.

---

## Cómo se usa (el usuario final)

1. Abre `Veilcord.exe`.
2. **Cerrá Discord** (el botón Inyectar igual lo cierra solo).
3. Apretá **Inyectar**.
4. Abrí Discord. Listo.

Para volver a Discord normal: abrir `Veilcord.exe` y apretar **Restaurar**.

---

## Cosas importantes (sin vueltas)

- **Windows va a avisar** ("Windows protegió tu PC") porque el `.exe` no está firmado todavía. Es normal en este tipo de programas. Se arregla más adelante pagando una firma digital.
- Si Discord se actualiza, hay que apretar **Inyectar** otra vez (Discord se reescribe en sus updates).
- Esta primera versión solo deja una **prueba** de que funciona (un archivo en `%TEMP%\veilcord-injected.txt`) y abre Discord normal. Las funciones de verdad (plugins, temas, privacidad) vienen del Mod, que se conecta en el siguiente paso.
- Por ahora anda en **Windows** (y macOS). Linux más adelante.
