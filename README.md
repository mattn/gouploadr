# gouploadr

flickr uploader

## Usage

gouploadr require access-token to upload photos.

```
$ gouploadr
```

You'll see flickr web page for authenticate. Make confirmation to the authentication, then type `Enter`.

```
GOUPLOADR_TOKEN=XXXXXXXXXXXX
```

You put following like into your `.bashrc`. On Windows, set your environment variable to the system environment dialog.

```
export GOUPLOADR_TOKEN=XXXXXXXXXXXX
```

To upload photos, you can specify filename and title like below.

```
$ gouploadr /path/to/the/photo.jpg#TITLE
```

## Installation

```
$ go get github.com/mattn/gouploadr
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a. mattn)
