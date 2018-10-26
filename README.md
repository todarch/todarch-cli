# Todarch CLI

Command line utility to interact with the Todarch application.

## Install

You can download the single executable on [releases](https://github.com/todarch/todarch-cli/releases).

```shell
# download
# mkdir $TODARCH_HOME as $HOME/.todarch
# extract to $TODARCH_HOME
# rename if you wish
mv $TODARCH_HOME/todarch-cli $TODARCH_HOME/td
# add $TODARCH_HOME to your $PATH
# create your own configuration as explain in the section
# finally
td version
```

You may use the installation script as well, it will do all for you.

```shell
curl https://raw.githubusercontent.com/todarch/todarch-cli/master/install | bash
```

### Configuration

The application expects to read a config file at $HOME/.todarch/config.yml location.

```yaml
todarchApiBase: https://api.todarch.com
```

### Commands

You need a Todarch account to use the cli application. You can create one on [todarch.com](https://todarch.com/register).

```shell
td check
td help
```

* Create todos

```shell
td create
td create --interactive
td create --editor
td create --file yourtododefinition.yml
```

* List your todos

```shell
# shows todos with INITIAL status
td todo ls

# show todos with any status
td todo ls -a

td todo ls -l
```

* See the details of your todo

```shell
td todo inspect $ID
td todo inspect $ID --output yml
```

* Get your todo done

```shell
td todo done $ID
```

* Delete your todos

```shell
td todo rm $ID
```

* Filter your todos using [rsql](https://github.com/jirutka/rsql-parser)

```shell
td todo ls --rsql priority=gt=5
td todo ls --rsql priority=lt=5

td todo ls --rsql timeNeededInMin=lt=60

# why do think the following form will not work as expected?
td todo ls --rsql priority>5
td todo ls --rsql priority<5
```

## Development & Usage

If you already have golang environment setup, you can have your own build.

```shell
go get github.com/todarch/todarch-cli
cd $GOPATH/src/github.com/todarch/todarch-cli
go build -o $GOPATH/bin/td
td help
```
