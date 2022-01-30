# ClickD

A super-simple analytics system.

## Config

```toml
# List of hosts that are allowed to send tracking data
AllowedHosts = [
    "http://localhost:8080"
]

# Configure your postgres database here
DbDsn = "host=localhost user=postgres password=postgres dbname=clickd port=5432 sslmode=disable"

# Host that the server is exposed as
Host = "http://localhost:8080"
```
