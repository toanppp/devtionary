# devtionary

Developer's dictionary CLI

## Usage

Read

```shell
devtionary -k [keyword]
```

Upsert

```shell
devtionary -k [keyword] -v [value]
```

Delete

```shell
devtionary -k [keyword] -d
```

Help

```shell
devtionary -h
```

## Installation

### Install executable

```shell
go install github.com/toanppp/devtionary@latest
```

### GOPATH

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:

- Linux or Mac
    ```shell
    export PATH=$PATH:$GOPATH/bin
    ```
  
- Windows
    ```shell
    set PATH=%PATH%;%PATHGOPATH%\bin
    ```

### MONGODB_URI

Set `MONGODB_URI` environment variable:

- Linux or Mac
    ```shell
    export MONGODB_URI=your-mongodb-connection-string
    ```

- Windows
    ```shell
    set MONGODB_URI=your-mongodb-connection-string
    ```
