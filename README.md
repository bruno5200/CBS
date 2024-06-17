# Core Secrets Manager

## Intrduction

This service is a simple secrets manager that allows you to store and retrieve secrets in a secure way. It uses a simple REST API to store and retrieve secrets. The secrets are encrypted using AES-256 encryption and stored in a database. The service uses a master key to encrypt and decrypt the secrets. The master key is stored in an environment variable and is used to encrypt and decrypt the secrets.

## Requirements

- [Go](go.dev/dl) 1.22 or higher
- [GNU Make](https://www.gnu.org/software/make/) 4.3 or higher
- [K8s](https://kubernetes.io/) 1.22 or higher
- [Memcached](https://memcached.org/) 1.6.9 or higher
- [PostgreSQL](https://www.postgresql.org/) 13.14 or higher

> [!NOTE]
> The service uses a PostgreSQL database to store the secrets. The database connection data is stored in an environment variables. The service also uses Memcached to cache the secrets. The Memcached connection string is stored in an environment variable.

> [!IMPORTANT]
> The service uses a master key to encrypt and decrypt the secrets. The master key is stored in an environment variable. Make sure to keep the master key secure and do not expose it to the public.

> [!CAUTION]
> For deployment in Google Cloud, make sure to use the latest version of gcloud cli and kubectl to configure the cluster.

## DB Docker (Required)

To run the database in a docker container, use the following command:
```bash
docker run -dp 5432:5432 --name=db -e POSTGRES_PASSWORD=lEy9gfGEqbdYxl1fWcqd -e POSTGRES_DB=vault -v ./vault-db:/var/lib/postgresql/data postgres:alpine
```

## Memcached Docker (Required)

To run the Memcached server in a docker container, use the following command:
```bash
docker run -dp 11211:11211 --name=memcached memcached:1.6.28-alpine -p 11211 -m 64 -vv
```