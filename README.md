# me-api

> Let your agent know you better - aggregate your digital footprint into one REST API

A personal context aggregation API that exposes `/me` with data from various macOS and cloud sources.

**macOS only** - Your Mac serves as the bridge, aggregating local system data (Health, presence, Music) and cloud services (Calendar).

## Run

```bash
go run .
```

## Build

```bash
go build -o bin/me-api .
```

## Endpoint

`GET /me` - Returns aggregated personal context data as JSON

## Data Sources

- [ ] Apple Health (sleep, HRV, heart rate, steps, calories)
- [x] Mac presence (status, focused app)
- [ ] Apple Music (current track, recently played)
- [ ] Apple Calendar (current/next meeting, availability)

## Security Warning

**Never expose this API publicly.** The `/me` endpoint contains sensitive personal data.

Use private network solutions like:
- [Tailscale](https://tailscale.com) - recommended, zero-config mesh VPN
- SSH tunnels for remote access
- WireGuard or other VPN solutions

## Architecture

- Aggregates data from local sources
- Exposes single JSON endpoint
- OpenCode polls every 30 minutes for context-aware decisions
