# Configuration Guide

For advanced setups, environments where a Web UI is undesirable, or fully automated deployments, you can configure `gPhotos2Immich` strictly via a `config.json` file.

If you don't use the Web UI, you can disable it entirely by setting the `DISABLE_WEBUI=true` environment variable.

## API Permissions

Your Immich API key needs these permissions (or use "All" for simplicity):

`asset.read` · `asset.upload` · `album.create` · `album.read` · `album.update` · `albumAsset.create` · `user.read`

## Example `config.json`

```json
{
  "apiKey": "YOUR_IMMICH_API_KEY",
  "apiURL": "http://your-immich-ip:2283/api",
  "debug": false,
  "workers": 4,
  "albumWorkers": 3,
  "strictMetadata": false,
  "skipVideos": false,
  "googlePhotos": [
    {
      "url": "https://photos.app.goo.gl/YourAlbumLink1",
      "albumName": "Vacation 2023",
      "syncInterval": "12h"
    },
    {
      "url": "https://photos.app.goo.gl/ExistingAlbumLink",
      "immichAlbumId": "existing-album-uuid",
      "syncInterval": "1h"
    }
  ]
}
```

## Options

| Key | Type | Default | Description |
| --- | --- | --- | --- |
| `apiKey` | string | — | Immich API key (required). |
| `apiURL` | string | — | Immich API URL, e.g. `http://localhost:2283/api` (required). |
| `debug` | bool | `false` | Enable verbose debug logging. When disabled, displays clean progress bars with speed and ETA. |
| `workers` | int | `1` | Number of concurrent download/upload workers **per album**. Controls how many photos within a single album are downloaded and uploaded in parallel. Higher values speed up large albums but use more bandwidth and memory. |
| `albumWorkers` | int | `1` | Number of albums processed **concurrently**. Controls how many albums are synced at the same time. Useful when you have many albums configured and want to process several in parallel. |
| `strictMetadata` | bool | `false` | Skip items with missing/invalid dates instead of uploading with current date. Skipped URLs are logged for manual review. |
| `skipVideos` | bool | `false` | Skip all video items entirely. Useful if you only want photos. |

## Album Options

| Key | Type | Default | Description |
| --- | --- | --- | --- |
| `googlePhotos[].url` | string | — | Google Photos shared album link (required). |
| `googlePhotos[].albumName` | string | auto-detected | Override the album name in Immich. If omitted, uses the album title from Google Photos. |
| `googlePhotos[].syncInterval` | string | `24h` | How often to re-check this album (e.g. `12h`, `60m`, `1h30m`). |
| `googlePhotos[].immichAlbumId` | string | — | Link to an existing Immich album by UUID instead of creating a new one. |
