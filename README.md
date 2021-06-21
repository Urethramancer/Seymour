# Seymour
A command line podcast downloader.

## Install
Ensure you have Go 1.16 or newer, and run this from a shell:

```
go get -u github.com/Urethramancer/Seymour
```

## Run

### Add feeds
Seymour requires podcast feeds to operate on. Add a feed with the `add` tool command. For example:

```
Seymour add https://example.com/podcast/rss
```

Ypu may also add a list of feeds:

```
Seymour add podcasts.txt
```

The file must be a list of feed URLs or absolute file paths for the feeds, one per line.

### Update
The `update` tool command updates all podcasts from the latest feed. Specify the (partial) name of a podcast to update only one. NOTE: Some podcasts may rotate out

### Download
The `download` tool command fetches episodes for poddcasts. Specify a (partial) name to download only specific podcasts. The supported flags are as follows:
- -e, --episode: Episode number to start from.
- -m, --mark: Also mark the episodes before the downloaded ones as downloaded.
- -f, --force: Force re-download of episodes.
- -l, --latest: Only download the latest episode.

### Remove
The `remove` tool command removes a specific podcast. The name must be an exact match.

### List
The `list` tool command lists all podcasts, or episodes for a specific podcast if a partial name is supplied.
