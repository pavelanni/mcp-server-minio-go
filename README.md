# MinIO / AIStor MCP server

This is a Model Context Protocol (MCP) for interacting with MinIO servers.
The server provides functions for listing buckets and objects, uploading and downloading files, getting and setting object tags.
Also, it provides the admin level functions such as getting information about the server's status.

## Security considerations

This MCP server has access to the MinIO/AIStor server via the credentials (`MINIO_ACCESS_KEY` and `MINIO_SECRET_KEY`) you provide it.
That means it can perform destructive operations such as deleting objects and buckets if you give it too much privileges.

We strongly suggest creating a separate user with read-only access to the MinIO server and use that user's credentials for the MCP server.
If you're interested in getting diagnostics information from the server, add the `diagnostics` policy to that user.

Here is how to do it. We use the playground server (`play.min.io`) as an example.
In the example below, `mcpserver` is the access key, and `mcppassword` is the secret key.

```shell
mc admin user add play mcpserver mcppassword
mc admin policy attach play readonly --user mcpserver
mc admin policy attach play diagnostics --user mcpserver
```

Create a new alias with the provided credentials and test if you can create a bucket.

```shell
mc alias set playmcp https://play.min.io mcpserver mcppassword
mc mb playmcp/mcp-bucket-test
```

Expected output:

```none
mc: <ERROR> Unable to make bucket `playmcp/mcp-bucket-test`. Access Denied.
```

But you should be able to get the server information with this alias:

```shell
mc admin info playmcp
```

Expected output:

```none
●  play.min.io
   Uptime: 8 hours
   Version: 2024-12-21T04:24:45Z
   Network: 1/1 OK
   Drives: 4/4 OK
   Pool: 1

┌──────┬──────────────────────┬─────────────────────┬──────────────┐
│ Pool │ Drives Usage         │ Erasure stripe size │ Erasure sets │
│ 1st  │ 1.8% (total: 80 GiB) │ 4                   │ 1            │
└──────┴──────────────────────┴─────────────────────┴──────────────┘

686 MiB Used, 266 Buckets, 5,135 Objects, 69 Versions, 1 Delete Marker
4 drives online, 0 drives offline, EC:2
```

If you want to create/delete buckets and upload/delete objects with the MCP server, use the user with the `readwrite` policy attached to it.

## Requirements

- Running MinIO server (alternatively, you can test it with the Playground MinIO server located at play.min.io)
- MinIO credentials to access the server: `MINIO_ACCESS_KEY` and `MINIO_SECRET_KEY` (see "Security considerations" above)
- macOS or Linux host

## Installation

Download the executable for your OS and CPU architecture from the [Releases page](https://github.com/pavelanni/mcp-server-minio-go/releases) and install it in your local directory that is included in your `PATH`.
The most common location is `$HOME/.local/bin/` both on macOS and Linux.

## Configuration

### Claude Desktop setup

1. Add the following lines to your Claude Desktop configuration file.
   Config file location: on MacOS: `~/Library/Application\ Support/Claude/claude_desktop_config.json`,
   on Windows: `%APPDATA%/Claude/claude_desktop_config.json`

   The easiest way to find the config file is the following:

   - In the main Claude Desktop menu (at the top of the screen on macOS and at the top of the window on Windows) click **Settings...**
   - In the next window, click **Developer** (you may have to enable it if it's not enabled yet)
   - In the next window, click **Edit Config**
   - On macOS, it opens a Finder window pointing to the `claude_desktop_config.json` file in the config location. On Windows it opens an Explorer window
   - Double-click that file, and it will be opened in your default text editor
   - Edit and save it
   - Exit Claude Desktop and start it again.

   Replace the path to the executable (`/PATH/TO/YOUR/mcp-server-minio-go`) with the actual location of the MCP server executable.

   Replace the access key and secret key with the actual values.
   If you are going to use `play.min.io` you can get these values by running `mc alias list`.

   ```json
   {
    . . . .
     "mcpServers": {
       . . . .
       "minio": {
         "command": "mcp-server-minio-go",
         "args": [
           "--allowed-directories",
           "~/Desktop",
           "~/Documents"
         ],
         "env": {
           "MINIO_ENDPOINT": "play.min.io",
           "MINIO_ACCESS_KEY": "REPLACE_WITH_ACCESS_KEY",
           "MINIO_SECRET_KEY": "REPLACE_WITH_SECRET_KEY",
           "MINIO_USE_SSL": "true"
         }
       }
     }
   }
   ```

## Usage

Start the Claude Desktop application. Look at the bottom-right of the screen: you should see a small icon of a hammer with a number next to it.
Click it and check the list of installed tools for this client. There should be several tools with "from server: minio" lines in them.
That means you installed the MinIO MCP server successfully.

Start with simple prompts in the Claude chat. For example:

- "List all buckets on my MinIO server"
- "List the contents of bucket test" (or whatever bucket you see in the list)
- "Download the file FILE.PDF to the Desktop directory on my computer"
- "Create bucket my-mcp-test on the MinIO server" (if your credentials allow that)
- "Upload the file DOCUMENT.PDF from my Documents directory to the bucket my-mcp-test" (if your credentials allow that)
- "Get metadata of the file DOCUMENT.PDF in the bucket my-mcp-test"
- etc.

You can also ask the server to get technical information about the MinIO server:

- "Give me admin information about the MinIO server"

You should expect a concise summary about your MinIO server (cluster) with the number of nodes and drives, amount of space, used and available, and similar information.
It's the output of the `mc admin info` command presented in human language.

## Development

This MCP server uses the [mcp-golang](https://mcpgolang.com/) library from [Metoro](https://metoro.io/).
Take a look at the docs and [examples](https://github.com/metoro-io/mcp-golang/tree/main/examples).

Feel free to add more MinIO functions and open a PR.
