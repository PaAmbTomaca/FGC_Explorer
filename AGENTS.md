# AGENTS.md for FGC_explorer

A Go project for exploring FGC (Ferrocarrils de la Generalitat de Catalunya) transit data via GTFS and OpenData APIs.

---

## 1. Project Overview
- **Language**: Go (1.27.1).
- **Purpose**: Fetch, parse, and display FGC transit data (routes, stops, trips) from:
  - [FGC OpenData API](https://dadesobertes.fgc.cat/api/explore/v2.1/)
  - GTFS feeds (via `google_transit.zip`).
- **Entrypoint**: `fgc_get.go` (CLI tool with a simple menu-driven interface).
- **Key Files**:
  - `go.mod`: Dependency manifest.
  - `fgc_get.go`: Core logic (data fetching, parsing, and display).
  - Downloaded GTFS files: `trips.txt`, `stop_times.txt` (extracted from `google_transit.zip`).

---

## 2. Setup & Dependencies
- **Dependencies** (from `go.mod`):
  ```bash
  go mod tidy  # Install dependencies.
  ```
  - `github.com/jackc/pgx/v5`: PostgreSQL driver (unused in current code but listed).
  - `github.com/jackc/pgpassfile`, `github.com/jackc/pgservicefile`: PostgreSQL config helpers (unused).
  - `golang.org/x/text`: Text processing (unused).

- **Environment**:
  - No `.env` or config files. Hardcoded URLs/paths in `fgc_get.go`.
  - GTFS data is downloaded dynamically from `https://www.fgc.cat/google/google_transit.zip`.

---

## 3. Key Commands
| Task                     | Command                          | Notes                                                                 |
|--------------------------|----------------------------------|-----------------------------------------------------------------------|
| Build                    | `go build`                       | Outputs `FGC_explorer` binary.                                        |
| Run                      | `go run fgc_get.go`              | Starts the CLI menu.                                                  |
| Download GTFS data       | Triggered via `download_data()`  | Fetches `google_transit.zip` and extracts `trips.txt`/`stop_times.txt`. |
| Test a function          | N/A                              | No tests or test framework configured.                                |

---

## 4. Architecture & Workflow
- **Data Flow**:
  1. **Download**: `download_data()` fetches `google_transit.zip` and extracts `trips.txt`/`stop_times.txt`.
  2. **Parse**:
     - `load_trips()`: Loads route IDs and `trip_id`s from `trips.txt`.
     - `load_stop(trip_id)`: Loads stop sequences for a given `trip_id` from `stop_times.txt`.
  3. **Display**: `print_linea(id)` prints stops for a route (e.g., "S1").

- **GTFS Quirks**:
  - `trip_id` in `trips.txt` uses a pipe delimiter (e.g., `6c4bdae602747613ef|6f2dc7e303`).
  - `stop_times.txt` maps `trip_id` → `stop_id` + `arrival_time`.

- **API Notes**:
  - FGC OpenData API supports `where`, `limit`, etc. (e.g., `?where=route_short_name="S1"`).
  - No authentication required.

---

## 5. Gotchas & Constraints
- **GTFS Parsing**:
  - `trip_id` in `trips.txt` is split by `|` (not a comma).
  - `stop_times.txt` uses `trip_id` to map stops (column 0).
- **Error Handling**:
  - Uses `panic()` for critical errors (e.g., file I/O). Avoid in production.
  - No retry logic for HTTP requests.
- **State**:
  - `lineas` (global `map[string]linea`) stores route data. Reset on restart.
- **Unused Code**:
  - PostgreSQL dependencies (`pgx`) are listed but unused.
  - `case 2` in the CLI menu is a stub.

---

## 6. Development Workflow
1. **Modify Code**:
   - Edit `fgc_get.go` (e.g., add new routes, fix parsing).
   - Test changes via `go run fgc_get.go`.
2. **Add Features**:
   - To query the FGC API, use `http.Get()` with `fgcbaseurl`.
   - To extend GTFS parsing, add new fields to the `linea` struct.
3. **Debugging**:
   - Use `fmt.Println()` for logging (no debugger configured).
   - GTFS files are overwritten on each `download_data()` call.

---

## 7. Future Improvements (Optional)
- Replace `panic()` with graceful error handling.
- Add tests (e.g., `go test`).
- Cache GTFS files to avoid repeated downloads.
- Use environment variables for URLs/paths.

---

## 8. References
- [FGC OpenData API Docs](https://dadesobertes.fgc.cat/api/explore/v2.1/)
- [GTFS Reference](https://gtfs.org/reference/static)