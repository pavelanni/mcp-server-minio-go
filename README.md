# MinIO / AIStor MCP server

This is a Model Context Protocol (MCP) for interacting with MinIO servers.
The server provides functions for listing buckets and objects, uploading and downloading files, getting and setting object tags.
Also, it provides the admin level functions such as getting information about the server's status.

## Requirements

- Running MinIO server (alternatively, you can test it with the Playground MinIO server located at play.min.io)
- MinIO credentials to access the server: `MINIO_ACCESS_KEY` and `MINIO_SECRET_KEY`
- macOS or Linux host

## Installation

Download the executable for your OS and CPU architecture from the Releases page and install it in your local directory that is included in yout `PATH`.
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
         "command": "/PATH/TO/YOUR/mcp-server-minio-go",
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
- "Create bucket my-mcp-test on the MinIO server"
- "Upload the file DOCUMENT.PDF from my Documents directory to the bucket my-mcp-test"
- "Get metadata of the file DOCUMENT.PDF in the bucket my-mcp-test"
- etc.

## Development

This MCP server uses the [mcp-golang](https://mcpgolang.com/) library from [Metoro](https://metoro.io/).
Take a look at the docs and [examples](https://github.com/metoro-io/mcp-golang/tree/main/examples).

Feel free to add more MinIO functions and open a PR.
