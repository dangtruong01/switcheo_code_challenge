# switcheo
**switcheo** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/switcheo@latest! | sudo bash
```
`username/switcheo` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

# Consensus-Breaking Changes Documentation

## Change Description
This update introduces a new mandatory `memo` field to the `MsgSend` transaction message. This field is intended for users to attach arbitrary data to their transactions.

## What Does Breaking Consensus Mean?
A consensus-breaking change is a change that is not compatible with the previous version of the blockchain protocol. Existing nodes will reject transactions that do not include the new `memo` field because it does not conform to their expected transaction structure.

## Why Would This Change Break Consensus?
This change would break consensus because transactions are the fundamental elements of the blockchain that all nodes must agree upon. The addition of a new mandatory `memo` field changes the transaction validation rules. Nodes running older versions of the software will consider transactions without a `memo` as invalid, while nodes that have upgraded will reject transactions that lack this field.

## Steps for Network Upgrade
1. Notify all node operators about the upcoming mandatory `memo` field.
2. Provide a software upgrade that includes the new transaction structure.
3. Define a block height when the new transaction structure will be enforced.
4. Coordinate with the node operators to ensure they have upgraded before the specified block height.
