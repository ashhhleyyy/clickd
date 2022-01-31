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

## Usage

### Tracker script

Add the following `<script>` tag to your site:

```html
<script src="https://analytics.example.org/tracker.js" async defer>
```

Replace the `analytics.example.org` with the domain you are hosting `clickd` on.

### Dashboard

You can access the dashboard at `https://analytics.example.org/dashboard/index.html`. By default, it is available to everyone, however if you need it to be private, you can configure a reverse proxy.

Replace `analytics.example.org` with the domain you are hosting `clickd` on.
