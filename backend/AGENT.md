# Backend Agent Guide

## Scope
This `backend/` directory is a Go web application (Echo + Ent ORM + Pagoda-style service container), not Django.

## Quick Start
```bash
cd backend
make run
```

Useful targets:
```bash
make watch      # hot reload via air
make test       # go test ./...
make css        # build Tailwind CSS into public/static/main.css
make build      # css + compile binary into ./tmp/main
make ent-gen    # regenerate Ent code after schema changes
```

## Entry Points
- `cmd/web/main.go`: app bootstrap, router build, task runner start, HTTP server lifecycle.
- `cmd/admin/main.go`: admin CLI (ex: create admin user).

## Architecture Map
- `pkg/handlers/`: route registration and HTTP handlers.
  - `router.go` wires middleware and all handlers.
  - `api*.go` contains API endpoints used by frontend.
- `pkg/middleware/`: auth/session/config/logging/cache middleware.
- `pkg/services/`: shared services container (`NewContainer`), auth, mail, cache, validation.
- `pkg/tasks/`: async/background task registration and handlers.
- `ent/`: ORM generated code + schemas in `ent/schema/`.
- `config/`: typed config loader (`config.go`) and defaults (`config.yaml`).
- `public/static/` + `tailwind.css`: backend-side CSS asset pipeline.

## Config and Environment
- Config is loaded through Viper from `config/config.yaml` (and env overrides without a prefix, e.g. `HTTP_PORT`).
- Default app port in config is `8000`; verify your local override before assuming API URL.
- DB defaults to Postgres on `localhost:5434`.

## Data + DB Notes
- Entity/schema source of truth: `ent/schema/*`.
- After schema changes:
  1. update schema file
  2. run `make ent-gen`
  3. run tests

## Frontend Contract Touchpoints
Main API surfaces consumed by SPA live in:
- `pkg/handlers/api.go`
- `pkg/handlers/api_course.go`
- `pkg/handlers/api_lms.go`
- `pkg/handlers/auth.go`

When response shapes change, update frontend types/service calls in `../web/src/types` and `../web/src/services/api.service.ts`.

## Safety Rules for Future Agents
- Do not use root `Makefile` backend commands; they reference a different (Django) workflow.
- Prefer `backend/Makefile` targets for this service.
- Avoid editing generated Ent files manually; edit `ent/schema/*` and regenerate.
- Run `make test` after backend changes.
