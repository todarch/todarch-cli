# Todarch CLI

Command line utility to interact with the Todarch application.

## Install

You can download the single executable on [releases](https://github.com/todarch/todarch-cli/releases).

```shell
# extract
# rename
mv todarch-cli todarch
# add it to $PATH
# finally
todarch version
```

## Development & Usage

If you already have golang environment setup, you can have your own build.

```shell
cd $GOPATH/src/github.com/todarch/todarch-cli
go build -o $GOPATH/bin/todarch
todarch help
```

### Configuration

The application expects to read a config file at $HOME/.todarch/config.yml location.

```yaml
todarchApiBase: http://localhost:7004
```

** Currently we do not have a publicly deployed version of Todarch Application, you may refer to [Todarch Docs](https://github.com/todarch/todarch-docs) to deploy your own version.

### Commands

You need a Todarch account to use the cli application.

```shell
todarch check
todarch help
```

* Create todos

```shell
todarch create
todarch create --interactive
todarch create --editor
todarch create --file yourtododefinition.yml
```

* List your todos

```shell
todarch todo ls
todarch todo ls -l
```

* See the details of your todo

```shell
todarch todo inspect $ID
todarch todo inspect $ID --output yml
```

* Get your todo done

```shell
todarch todo done $ID
```

* Delete your todos

```shell
todarch todo rm $ID
```
