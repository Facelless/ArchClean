#  Go Disk Cleaner

A **Golang-based disk cleanup utility** designed to **scan the file system**, identify **large or non-essential files**, and **safely delete them** to free up disk space.

>  **Warning:** This tool permanently deletes files. Use with caution and always test before running in production environments.

---

##  Features

-  Recursive directory scanning
-  Detection of files above a configurable size threshold
-  Protection against deletion of critical system paths and files
-  Automatic removal of non-essential large files
-  Final summary report including:
  - Number of deleted files
  - Total disk space reclaimed (MB / GB)
-  Compatible with **Windows** and **Linux** (path adjustments required)

---

##  Requirements

- Go **1.18+**
- Administrative / root privileges (recommended)

---

##  Getting Started

### 1️ Clone the repository
```bash
git clone https://github.com/facelless/ArchClean.git

