<a href="https://github.com/burntcarrot/hashsearch">
    <img align="left" src="assets/hashsearch.png" width="256">
</a>

#### **`hashsearch`** — Reverse image search using perceptual hashes. 🔍

**hashsearch** is a lightweight, easy-to-use reverse image search engine which makes use of perceptual hashes.

[![justforfunnoreally.dev badge](https://img.shields.io/badge/justforfunnoreally-dev-9ff)](https://justforfunnoreally.dev)

<br>
<br>
<br>
<br>
<br>

## Installation

The easiest way is to download from the [releases](https://github.com/burntcarrot/hashsearch/releases).

You could also build `hashsearch` from the source code:

```sh
git clone https://github.com/burntcarrot/hashsearch
cd hashsearch
go build -o ./bin/hashsearch api/v1/main.go
```

## Usage

```
hashsearch -config <CONFIG_FILE_PATH>
```

If `-config` isn't provided, `hashsearch` defaults to `<HOME_DIR>/.hashsearch/config.yml`.

`hashsearch` runs the server on the configured address and exposes an API to interact with.

## API

The API is very simple. Two routes, one for searching and one for getting the list of images.

### `/v1/search`

Post an image using form data, get list of images (sorted by least to most distance):

```sh
curl --location --request POST 'localhost:8081/v1/search' \
--form 'file=@"star.png"'
```

Response:

```json
[
  {
    "path": "files/star.png",
    "distance": 0,
    "hash": "0000000000010000111100001111110011111100111100000001000000000000"
  },
  {
    "path": "files/star-new.png",
    "distance": 4,
    "hash": "0001000000110000111100001111110011111100111100000011000000010000"
  },
  {
    "path": "files/random.png",
    "distance": 28,
    "hash": "0000000110000000110000100010001111110010010001100000011110000110"
  }
]
```

### `/v1/list`

Get list of all images:

```
curl --silent 'localhost:8081/v1/list'
```

Response:

```json
[
  {
    "path": "files/random.png",
    "distance": 0,
    "hash": "0000000110000000110000100010001111110010010001100000011110000110"
  },
  {
    "path": "files/star-new.png",
    "distance": 0,
    "hash": "0001000000110000111100001111110011111100111100000011000000010000"
  },
  {
    "path": "files/star.png",
    "distance": 0,
    "hash": "0000000000010000111100001111110011111100111100000001000000000000"
  }
]
```

## Configuration

The configuration file is a simple `.yaml` file:

```yaml
db:
  url: "data.db" # Database URL.

server:
  addr: "localhost:8081" # Server address.

files:
  dir: "/files" # Directory where the images would be saved.

cors:
  allow_origin: "*" # CORS Allow-Origin value.

logging:
  file: "/hashsearch.log" # Log file path.
```

## How does this work?

You upload an image using `/v1/search` route:

- `hashsearch` makes a copy of your image
- `hashsearch` stores the copied image in `FILES_DIR`, which is configurable
- `hashsearch` generates the hash when you post the image, and saves it to the database
- `hashsearch` computes the distances between the posted image and other images and returns the result as a JSON response
- This response is sorted on the basis of the distance; and you should get the *most similar* images at the beginning of the response.

## Where can I use this?

If you have a small-scale application, and you don't want to make use of large dependencies/systems, this should work fine.

**Is it the best solution?** Not really, but if you want a quick and easy solution, this should be good enough.

**Is it blazing fast?** Again, not sure about this; I haven't tested it out on large sets of images.

## License

`hashsearch` is licensed under the [MIT license](./LICENSE).

## References

[Looks Like It](https://www.hackerfactor.com/blog/index.php?/archives/432-Looks-Like-It.html) is one of the inspirations behind this project.

## What's next?

This isn't how reverse image search is implemented in most areas; I just wanted to have some fun with perceptual hashes.

Average hash is fine for most cases, but it struggles in some areas, so the better option would be to use dHash/pHash.

I'm actively working on **reverse video search**; expect it to be a part of the future releases.

A nice, little web UI would also be added soon.
