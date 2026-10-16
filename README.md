<div align="center">
  <h1>🐹 Vole</h1>
  <p><b>Keep your Linux clean, light, and running like new. Fast, native Linux system cleaner, app uninstaller, and live system monitor.</b></p>
  <p><i>A Linux-native fork of <a href="https://github.com/tw93/Mole">tw93/Mole</a></i></p>

  <a href="https://github.com/wfzyx/vole/stargazers"><img src="https://img.shields.io/github/stars/wfzyx/vole?style=flat-square" alt="Stars"></a>
  <a href="https://github.com/wfzyx/vole/releases"><img src="https://img.shields.io/github/v/release/wfzyx/vole?label=version&style=flat-square" alt="Version"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-GPL_v3-blue.svg?style=flat-square" alt="License"></a>
  <a href="https://ko-fi.com/wfzyx"><img src="https://img.shields.io/badge/Support%20on-Ko--fi-ff5e5b.svg?style=flat-square&logo=kofi&logoColor=white" alt="Support on Ko-fi"></a>
</div>

---

## What is Vole?

**Vole** (`vo`) brings the fast, beautiful terminal maintenance suite of macOS [Mole](https://github.com/tw93/Mole) to **Linux** (with first-class support for Fedora, RHEL, Ubuntu/Debian, and Arch).

It combines disk cleaning, application removal, maintenance optimization, an interactive disk usage treemap, and a live hardware monitor in a single zero-dependency terminal command.

```text
 __     __     _      
 \ \   / /___ | | ___ 
  \ \ / // _ \| |/ _ \
   \ V /| (_) | |  __/  https://ko-fi.com/wfzyx
    \_/  \___/|_|\___|  Deep clean and optimize your Linux.
```

---

## Features

- **🧹 Deep Clean (`vo clean`)**: Reclaims gigabytes by removing user caches, thumbnail caches, browser data (Firefox, Chrome, Brave, Edge), dev tooling (`npm`, `pip`, `go-build`, `cargo`, `uv`, `bun`), DNF / APT package caches, Flatpak runtimes, and systemd journals.
- **🗑️ Smart Uninstaller (`vo uninstall`)**: Discovers installed GUI apps (RPM packages, Flatpaks, and local desktop entries), calculates exact disk footprints, detects last-used timestamps, and hunts down leftover configs in `~/.config`, `~/.cache`, and `~/.local/share`.
- **⚡ System Optimizer (`vo optimize`)**: Runs safe Linux maintenance tasks: systemd journal vacuuming, SSD block trimming (`fstrim`), clearing failed systemd units, rebuilding font caches (`fc-cache`), refreshing MIME desktop databases, and cleaning orphaned dependencies (`dnf autoremove`).
- **🗺️ Interactive Disk Analyzer (`vo analyze`)**: Ultra-fast Go + BubbleTea disk usage explorer. Visualizes disk space with real-time navigation, folder drill-downs, and safe trash deletion.
- **📊 Live Hardware Dashboard (`vo status`)**: Real-time terminal dashboard tracking CPU per-core load, memory & swap pressure, NVMe/Btrfs/Ext4 mounts, network throughput, battery thermals, and high-CPU/zombie process watchdogs.
- **📦 Project Build Purge (`vo purge`)**: Recursively cleans bloated project artifacts: `node_modules`, `.venv`, Rust `target`, `__pycache__`, `.pytest_cache`, `.next`, and build directories.
- **📥 Leftover Installer Cleanup (`vo installer`)**: Finds orphaned `.rpm`, `.AppImage`, `.run`, and `.deb` files in Downloads and Desktop.

---

## Quick Start

### Installation

Install via curl:

```bash
curl -fsSL https://raw.githubusercontent.com/wfzyx/vole/main/install.sh | bash
```

Or build from source:

```bash
git clone https://github.com/wfzyx/vole.git
cd vole
make build
# Binary is in ./vole with symlink ./vo
./vo
```

### Basic Usage

```bash
vo                           # Interactive main menu
vo clean                     # Deep cleanup
vo clean --dry-run           # Preview what would be cleaned without touching files
vo uninstall                 # Interactive application uninstaller
vo uninstall slack           # Uninstall Slack and purge leftovers
vo optimize                  # Run safe system optimizations
vo optimize --dry-run        # Preview optimizations
vo analyze ~                 # Explore disk usage in home directory
vo status                    # Live system health dashboard
vo purge                     # Interactively purge project dependencies (node_modules, target, etc.)
vo installer                 # Find and remove leftover .rpm, .AppImage, .deb files
```
---

## 🗺️ Distro Roadmap

Vole is expanding native package manager, service management, and cache layout support across all major Linux families:

- [x] **RPM-based**: Fedora, RHEL, CentOS Stream, Rocky Linux, AlmaLinux (`dnf`, `rpm`)
- [ ] **Arch-based**: Arch Linux, Manjaro, EndeavourOS (`pacman`, `yay`, `paru`)
- [ ] **Deb-based**: Debian, Ubuntu, Linux Mint, Pop!_OS (`apt`, `dpkg`)
- [ ] **openSUSE / SUSE-based**: Tumbleweed, Leap (`zypper`, `rpm`)
- [ ] **NixOS / Guix**: Declarative package and store pruning
- [ ] **Desktop GUI**: Native GTK4/Qt6 or Tauri desktop companion app

---
## Support & Sponsorship

If Vole saves you disk space or makes managing your Linux system easier, please consider supporting development!

☕ **[Support on Ko-fi](https://ko-fi.com/wfzyx)**

---

## Credits & Attribution

Vole is a Linux-native fork of [Mole](https://github.com/tw93/Mole), originally created by [Tw93](https://github.com/tw93).
Licensed under the [GNU General Public License v3.0](LICENSE).
