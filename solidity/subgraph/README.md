# Injective Peggo Subgraph 🤖 🚀

This is the Subgraph for the Injective Peggo.

---

## Local Subgraph development

### 0. Install dependencies

```bash
$ yarn install
```

### 1. Run local ganache

```bash
$ yarn run ganache -h 0.0.0.0
```

### 2. Deploy contracts

See deploy script.

### 3. Run local Graph Node

Follow the instructions here: https://github.com/InjectiveLabs/graph-node#quick-start. You might have to use an older IPFS version, see [here](https://github.com/graphprotocol/graph-node/issues/1799#issuecomment-661433084).

```bash
$ cd $GOPATH/src/github.com/InjectiveLabs
$ git clone https://github.com/InjectiveLabs/graph-node.git
$ cd graph-node
# NOTE: if you do not already have the dependencies required installed (Rust, IPFS, Postgres), run this after completing 3a.
$ cargo build
```

### 3a (OPTIONAL). Install The Graph dependencies.

First install Rust .

```bash
$ curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
$ source $HOME/.cargo/env
```

Next, install PostgreSQL.

On **MacOS**, you can install using brew.

```bash
$ brew install postgresql
$ brew services start postgresql
```

Alternatively, you can use the [graphical installer](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads).

On **Linux**, run the following:

```bash
# See full instructions on https://www.postgresqltutorial.com/install-postgresql-linux/
# First, execute the following command to create the file repository configuration:
$ sudo sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
# Second, import the repository signing key:
$ wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | sudo apt-key add -
# Third, update the package list:
$ sudo apt-get update
# Finally, install the latest version of PostgreSQL:
$ sudo apt-get install postgresql
```

Next, start the PostgreSQL database

```bash
$ initdb -D .postgres
$ pg_ctl -D .postgres -l logfile start
$ createdb graph-node
```

Next, install **IPFS**. Note: you might have to use an older IPFS version, see [here](https://github.com/graphprotocol/graph-node/issues/1799#issuecomment-661433084).

```bash
# Download older version of IPFS until this issue is fixed
$ wget https://github.com/ipfs/go-ipfs/releases/download/v0.4.23/go-ipfs_v0.4.23_darwin-amd64.tar.gz
$ tar -xvzf go-ipfs_v0.4.23_darwin-amd64.tar.gz
$ cd go-ipfs
$ ./install.sh
> Moved ./ipfs to /usr/local/bin
$ ipfs --version
> ipfs version 0.4.23
$ ipfs init
$ ipfs daemon
# If you get Error: serveHTTPGateway: manet.Listen(/ip4/127.0.0.1/tcp/8080) failed: listen tcp4 127.0.0.1:8080: bind: address already in use, run the following:
# sudo lsof -i tcp:8080
# kill -9 <PID>
```

The final command to run the node will be:

```bash
$ cargo run -p graph-node --release -- --postgres-url postgresql://$USER@localhost:5432/graph-node --ethereum-rpc development:http://localhost:8545 --ipfs 127.0.0.1:5001
# Or, if your PostgresDB has a password:
export $PASSWORD=<insert_postgres_password>
$ cargo run -p graph-node --release -- --postgres-url postgresql://$USER[:$PASSWORD]@localhost:5432/graph-node --ethereum-rpc development:http://localhost:8545 --ipfs 127.0.0.1:5001

```

### 4. Create Subgraph

```bash
$ yarn codegen-local
$ yarn create-local
```

### 5. Deploy Subgraph

```bash
$ yarn deploy-local
```

## Injective Subgraph deployment

Optional: Deploy Peggy contracts if required. Make sure to update the hard-coded peggy state id inside `mapping.ts`.

1. Update Subgraph

```bash
$ yarn codegen:mainnet
$ yarn deploy:mainnet
```

## ⛑ Support

Reach out to us at one of the following places!

- Website at <a href="https://injectiveprotocol.com" target="_blank">`injectiveprotocol.com`</a>
- Twitter at <a href="https://twitter.com/InjectiveLabs" target="_blank">`@InjectiveLabs`</a>

---

## 🔐 License
