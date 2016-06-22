# unmake - Display usage for Makefile


## Install

```
$ go get github.com/TakesxiSximada/unmake
```

## How to use

### Syntax

```
.PHONY: command
command:
    @## param1=value1 param2=value2
    @#
    @# document 1
    @# document 2
    @# document 3

    echo "EXECUTE COMMANDS..."
```

### Execute

```
$ unmake /PATH/TO/YOUR/Makefile
```

For example (See https://github.com/TakesxiSximada/unmake/blob/master/examples/Makefile)

```
$ unmake examples/Makefile
* help

      display usage

* run target=NO

      execute an application.


* get target=NO_TARGET

      Install a application.
      ex) unmake

```
