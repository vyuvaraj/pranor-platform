# Pipeline Status Monitor Showcase App

This is a showcase application written entirely in the **Serv Programming Language (`pranor`)**. It demonstrates how to build web applications with persistent database storage, background cron/every loops, and HTTP services using zero boilerplate.

## Architecture

- **HTTP Server**: Runs natively on port `9000` exposing a web dashboard.
- **SQLite Persistence**: Stores the pipeline status of the 15 repositories in a local `pipeline.db` using declarative schema migrations.
- **Daemon Worker**: Runs a background loop every 30 seconds (`every 30s`) that queries the public GitHub Actions runs API to pull the latest build conclusions.
- **UI Presentation**: Serves a premium glassmorphic, responsive HTML dashboard showing live build statuses (success, failure, or in-progress) and triggered author credits.

## Running the Showcase App

To run the application locally in watch/hot-reload mode:

```bash
# Navigate to the showcase directory
cd pipeline-dashboard

# Run the compiler dev server
pranor run main.pnr --watch
```

To compile the application into a standalone native binary:

```bash
pranor build main.pnr -o pipeline-dashboard.exe
./pipeline-dashboard.exe
```

Open [http://localhost:9000](http://localhost:9000) in your browser to view the live dashboard!
