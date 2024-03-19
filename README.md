# Office Days Agent

This is the background agent for the automated office days tracker.

## What does it do?

   1. it periodically checks if you are connected to the office wifi, and records the days (it does not store any other data)
   2. it provides a REST API for other processes to use this information

## Installation

You can use the included `Makefile` to build and install the agent:

```
sudo make clean install
```

`sudo` is required to install the Launch Daemon on Mac.

In order to check if the agent is running properly, you can use `launchctl`:

```
sudo launchctl list | grep office-days-agent
```

You should see the agent and the associated process id listed in the response.
If you don't see it, check the Troubleshooting section of this readme.


## API

| Path     | Method | Description                                                                    |
| -------- | ------ | ------------------------------------------------------------------------------ |
| `/list`  | GET    | Get the stored days. Example response: `{"20240101": true, "20240102": false}` |
| `/clear` | POST   | Clear the stored days.                                                         |

By default the API can be reached on http://localhost:23460/.


## Troubleshooting

### I can't see the agent in the list

Make sure to use `sudo` for installing __and__ querying `launchctl`, otherwise the agent will be installed as a user-agent instead of a system daemon.

If you accidentally installed the agent incorrectly:

   1. `sudo make uninstall` to remove everything
   2. `launchctl unload hu.okki.office-days-agent` (without `sudo`) to unload the _user_ agent

### How can I check if the agent is running properly?

   1. check if it is installed correctly to `/usr/local/office-days-agent`. The directory should contain the binary, and may contain a couple of other files (db, logs, etc).
   2. check if the launch daemon is installed properly: `/Library/LaunchDaemons/hu.okki.office-days-agent.plist`.
   3. check that the agent has started up (see previous step)
   4. check if there is anything in the log files:
      1. `/usr/local/office-days-agent/stdout.log`
      2. `/usr/local/office-days-agent/stderr.log`


## FAQ

### How can I change the Office Wifi Name?

In case you want the agent to monitor a different SSID, you can use the `OFFICE_DAYS_WIFI_NAME` environment variable.
If you used the Mac Launch Daemon, you will need to change this environment variable in `/Library/LaunchDaemons/hu.okki.office-days-agent.plist`.

### What configuration options are supported?

Here is a comprehensive list:

| Environment Variable                | Mandatory/Default                 | Description                                   |
| ----------------------------------- | --------------------------------- | --------------------------------------------- |
| `OFFICE_DAYS_API_SERVER_PORT`       | `23460`                           | the port the api server is opened on          |
| `OFFICE_DAYS_DB_PATH`               | `/usr/local/office-days-agent/db` | database for the agent                        |
| `OFFICE_DAYS_POLL_INTERVAL_SECONDS` | `600`                             | how frequently the agent checks the wifi ssid |
| `OFFICE_DAYS_WIFI_NAME`             | *                                 | the office wifi ssid                          |
| `OFFICE_DAYS_WIFI_DEVICE_NAME`      | `en0`                             | the wireless network adapter name             |

