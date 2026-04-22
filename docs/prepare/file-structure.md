# File Structure

A standard Go project file structure combined with early documentation.

Early project documentation, as shown the `docs/` directory, is produced because, as a growing engineer in trade, documentation is just as important. Plus, to get more attached to *design briefs* and *decision matrices*.

```txt
golang-stl-format-converter
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── gui/
│       └── main.go
├── docs/
│   └── prepare/
│       └── file-structure.md
├── go.mod
├── internal/
│   ├── api/
│   │   └── handlers.go
│   └── converter/
│       └── converter.go
├── LICENSE
├── pkg/
└── README.md
```

> **Note on Decision Matrix**: Omitted to prioritize implementation. As this is a focused utility with a defined tech stack (Go), a Design Brief is sufficient to guide development without redundant cross-solution comparisons.

