# Description
(com-poh-sih-ler) A simple golang templating program for docker-compose, taking in
modular compose service, volume, network, and secret definitions,
merging them together into a single docker-compose.yml, and 
writing out any values from an associated config file for that
applications various environments. Modular compose component templates
are stored in `.tmpl` files, while environmental configs are
stored in `.json` files.

# Usage
## Basic
`composiler <environment>`
Composiler takes the associated environment argument and merges the
configs in the <environment>.conf file with all available templates
to deploy the full application stack as defined by your compose
templates. Great for deploying to a Swarm where all services are
run on the same logical docker daemon.
##### Example
`composiler load`


## Targeted
`composiler [--service <svc1>,...] [--network [<net1>,...]]
[--volume [<vol1>,...]] [--secret [<sec1>,...]] <environment>`
When top-level docker-compose keys are specified (in a
comma-separated list when multiples are involved) composiler will
load any of the named docker-compose template components from the
associated directory when building the final docker-compose.yml
file. If the `--service` option is specified, at least one service
name must be supplied. If any of the other top-level compose keys
are specified, a comma-separated list of named volumes, networks,
etc. may be supplied to perform a similar targeted loading of the
compose file component, or these flags may be supplied without
any arguments to load all of the components for that key-type.
##### Example
`composiler --service webapp,redis --network --volume redis load`

Unless directed otherwise (`--conf` flag), composiler will look
for the compose component templates and the environment configuration
files to be in the `/composiler` directory. The directory structure
for composiler to find the necessary components to build the final
compose file is:
```
composiler
    ├── configs
    │   ├── load.json
    │   ├── production.json
    │   └── test.json
    └── templates
        ├── networks
        │   └── redis.tmpl
        ├── secrets
        ├── services
        │   ├── apache.tmpl
        │   ├── redis.tmpl
        │   └── web.tmpl
        └── volumes
            └── redis.tmpl
```