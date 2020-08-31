# uptimed

This is uptimed. The poor man's dyndns.
It calls a webserver under your control regularly and comes with a service template.

## Installation

You can also check out the repository and build it with a uptodate golang toolchain, like so:

```
git clone git@github.com:4thel00z/uptimed.git
cd uptimed
make build
# This will create the service template for you
env INTERVAL=30 USER=other_username URL=https://your.service make service > /tmp/uptimed.service
```

## How does it work

The url to be called can be configured using the `-url` flag or the `UPTIMED_URL` environment variable.
The sleeping interval defaults to 30 seconds and can be set via the `-interval` flag.
These are populated in the service file for you, as can be seen in the installation section
