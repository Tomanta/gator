# gator
boot.dev blog aggregator project


To start the postgres server: `sudo service postgresql start`

Enter shell: `sudo -u postgres psql`

`psql "postgres://postgres:postgres@localhost:5432/gator"`

SQLC docs:
https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html

You will need Postgres and Go installed to run this program. Update `.gatorconfig.json` to point to the postgres install.