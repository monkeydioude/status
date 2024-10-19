# Dashboard for a quick check of localhost service' status

## FLAGS!

- `-c /path/to/config.json`: (**MANDATORY**) points to the json config file
  ([see example](./config_example/config.json))

## ENV VARS!

- `BASIC_AUTH_SHA256_LOGIN`: (OPTIONAL) if set along
  `BASIC_AUTH_SHA256_PASSWORD`, will activate basic authentification
- `BASIC_AUTH_SHA256_PASSWORD`: (OPTIONAL) if set along
  `BASIC_AUTH_SHA256_LOGIN`, will activate basic authentification

### BASIC AUTH!

Activates only if `BASIC_AUTH_SHA256_LOGIN` and `BASIC_AUTH_SHA256_PASSWORD` env
vars are set.
