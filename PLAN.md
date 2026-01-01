# things3-cli TODO

Goal: deliver a Go implementation that is functionally equivalent to the
reference `things-cli` project, including help/man docs and full unit +
integration test coverage.

## Assumptions / decisions

- [x] Binary name: `things` (parity with reference)
- [x] CLI framework: Cobra
- [x] Target platform: macOS only (Things URL scheme + `open`)
- [x] No persistent config file for v1 (flags/env only)

## Project scaffolding

- [x] Initialize Go module (`go.mod` + `go.sum`)
- [x] Add Cobra dependency
- [x] Create baseline repo structure:
  - `cmd/things/` main entry
  - `internal/cli/` command tree + flag parsing
  - `internal/things/` URL builder + helpers
  - `internal/open/` launcher abstraction
- [x] Add Makefile targets: `build`, `test`, `install`
- [x] Define version metadata (ldflags + `--version` output)

## Core helpers + URL scheme layer

- [x] Implement URL encoding to match reference behavior
- [x] Implement input parsing helpers:
  - `findTitle` (first line)
  - `findNotes` (rest of lines, trim blank edges)
  - `join` for checklist/title lists (`%0A` join)
- [x] Implement launcher abstraction:
  - default `open`
  - respect `OPEN` override
  - error if command not found
- [x] Implement URL builder for `add`
- [x] Implement URL builders for:
  - [x] `add-project`
  - [x] `update`
  - [x] `update-project`
  - [x] `show`
  - [x] `search`

## Commands (Cobra)

- [x] Wire root command + `help` command
- [x] Implement `add` command (flags + stdin `-` handling)
- [x] Implement `add-project` command (flags + stdin `-` handling)
- [x] Implement `update` command (requires `--auth-token` + `--id`)
- [x] Implement `update-project` command (requires `--auth-token` + `--id`)
- [x] Implement `show` command (`--id` or query)
- [x] Implement `search` command (optional query)
- [x] Implement `--version` output (CLI + Things app version)
- [x] Match reference help text formatting
- [x] Match reference error messages
- [x] Implement `--debug` behavior (reference-like diagnostics)

## Docs + distribution

- [x] Update README with reference repo link and parity goals
- [x] Add usage + install instructions for Go binary
- [x] Write `doc/man/things.1.md` (markdown source)
- [x] Generate `share/man/man1/things.1` (man page build)
- [x] Add install/uninstall steps to Makefile
- [x] Add `CHANGELOG.md` with Unreleased section
- [x] Add release checklist (`docs/RELEASING.md`)
- [x] Add release scripts (build artifacts + release notes + gh release)
- [x] Add GitHub Actions release workflow
- [x] Add minimal GitHub Pages site (`docs/site/index.html`)
- [x] Add GitHub Pages workflow

## Database access (read-only)

- [x] Implement Things DB path resolution (`THINGSDB`, ThingsData-* fallback)
- [x] Add SQLite read-only store (pure Go driver)
- [x] Add `projects`, `areas`, `tags`, `tasks` commands
- [x] Add JSON/table output for DB commands
- [x] Document DB access + Full Disk Access note
- [x] Expand DB schema coverage (notes, tags, checklist items, dates, metadata)
- [x] Implement Things date decoding (start_date, deadline, stop_date, created, modified)
- [x] Match Things sidebar list logic (inbox, today, upcoming, anytime, someday, logbook, completed, canceled, trash)
- [x] Add today prediction rules (unconfirmed scheduled + overdue) to match things.py
- [x] Implement deadlines list (deadline-only view)
- [x] Implement createdtoday / logtoday / logbook filters
- [x] Add search across title/notes/area (match things.py)
- [x] Add recursive tree output (areas → projects → headings → todos)
- [x] Add --recursive flag and context filters (project/area/tag) consistent with reference
- [x] Add output formats beyond table/JSON (not needed for v1)
- [x] Add “all” aggregation view (Inbox/Today/Upcoming/Anytime/Someday/Logbook/Areas)
- [x] Add tags ordered by usage (match reference behavior)
- [x] Add CLI command parity with thingsapi/things-cli:
  - [x] projects (basic)
  - [x] areas (basic)
  - [x] tags (basic list)
  - [x] tasks/todos (basic)
  - [x] today (basic)
  - [x] inbox
  - [x] upcoming
  - [x] anytime
  - [x] someday
  - [x] completed
  - [x] canceled
  - [x] trash
  - [x] logbook
  - [x] logtoday
  - [x] createdtoday
  - [x] deadlines
  - [x] all (aggregate)
  - [x] search (DB-backed)
  - [x] show (DB-backed)
- [x] Align flags with reference CLI:
  - [x] Add `-d/--database` alias for `--db`
  - [x] Add `-p/--filter-project`, `-a/--filter-area`, `-t/--filter-tag`
  - [x] Add `--filtertag` alias for `--filter-tag`
  - [x] Add `-e/--only-projects`
  - [x] Add `-r/--recursive`
  - [x] Add `-j/--json`
- [x] Add parity tests with a fixture DB (ported from thingsapi tests)

## Tests (tracked)

- [x] Unit tests for URL encoding + helpers
- [x] Unit tests for `add` URL builder behavior
- [x] CLI tests for `add` command (stdin + args)
- [x] Unit tests for `show` + `search` URL builders
- [x] CLI tests for `show` + `search` commands
- [x] Unit tests for `update` + `update-project` URL builders
- [x] CLI tests for `update` + `update-project` commands
- [x] CLI tests for `add-project` command (stdin + args)
- [x] Port reference bats tests as Go integration tests:
  - [x] `add-project`
  - [x] `update`
  - [x] `update-project`
  - [x] `show`
  - [x] `search`
- [x] Add tests for help/version output and error handling
- [x] Ensure tests pass without Things.app (stub `open` + version)
- [x] Run `go test ./...`

## Hybrid: URL scheme writes + DB reads

- [x] Keep URL-scheme write commands (`add`, `update`, `add-project`, `update-project`)
  - [x] Add background open (`open -g`) as default to avoid stealing focus
  - [x] Add `--foreground` flag to allow activating Things when desired
  - [x] Add `--dry-run` to print the URL without opening
- [x] Replace read commands with DB-backed equivalents where possible
  - [x] `show` → DB lookup + output
  - [x] `search` → DB search (title/notes/area), no `open`
- [x] Keep launcher abstraction but support background open toggle
- [x] Update root help/man/README to document hybrid behavior (writes via URL scheme, reads via DB)
- [x] Update tests:
  - [x] Write commands should not require Things (mock `OPEN`)
  - [x] Read commands use fixture DB with CLI-level tests
- [x] Ensure read commands never open Things.app during normal operation

## QA + parity validation

- [x] Compare CLI behavior vs reference for every command/flag (not required)
- [x] Verify Things URL scheme links match expected output
- [x] Verify man page content mirrors reference sections (incl. DB read commands)
- [x] Run full test suite on macOS

## Definition of done

- [x] All commands and options match reference behavior (manual parity sweep not required)
- [x] CLI help + man page present and accurate
- [x] Unit + integration tests cover all commands and edge cases (fixture + integration coverage sufficient for v1)
- [x] `make test` passes without Things.app installed
