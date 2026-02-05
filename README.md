### Argus — simple file integrity checker (Go CLI)

Argus is a tiny command‑line tool that creates a snapshot of file hashes for a folder and later compares the current state against that snapshot to detect tampering or unexpected changes.

It’s built with Go and uses Cobra for the CLI.

#### Features
- Recursively scans a folder and computes hashes for files
- Saves a snapshot to a plain text `.ags` file
- Compares current files to a previous snapshot and reports:
  - Changed files (hash mismatch)
  - New files not present in the snapshot

---

### Requirements
- Go 1.20+ (recommended if you are building it yourself)
- Windows, macOS, or Linux

---

### Build

Clone the repo and build the CLI:

```
go mod download
go build -o out/argus.exe ./main
```

Notes:
- The project’s main package is in `main/main.go`.
- On non‑Windows systems you can change the output name, e.g. `-o out/argus`.


### Running Argus 

Run the executable you've built or found in the Release section.

---

### Usage

Argus has two subcommands: `snap` and `compare`.

1) Create a snapshot of a folder

```
argus snap [folder path]
```

This scans the folder and writes a snapshot file named `argus_snapshot.ags` into the target folder.

Important path note: the current implementation concatenates the folder path and the file name. If your folder path does not end with a path separator, the snapshot file name may be concatenated directly. As a safe practice, include a trailing separator when you run `snap`.

Examples
- Windows:
```
out\argus.exe snap C:\data\projects\
```
- macOS/Linux:
```
./out/argus snap /home/user/projects/
```

2) Compare current folder state against a snapshot

```
argus compare [folder path] [snapshot path]
```

Examples
- Windows:
```
out\argus.exe compare C:\data\projects C:\data\projects\argus_snapshot.ags
```
- macOS/Linux:
```
./out/argus compare /home/user/projects /home/user/projects/argus_snapshot.ags
```
