# clz-xml-data-translator-cli

## Development

Small CLI utility to translate CLZ game collection datafrom XML to JSON. Additionally will hydrate additional game information from IGDB.

Update/Install Dependencies:
`make deps`

Test:
`make test`

Build:
`go build -C src -o ../build/main`

## Use

Translate provided CLZ game collection data in XML format to JSON

**Usage:** `CLZTranslate translate [flags]`

**Flags:**

- `-h, --help`: help for translate
- `-i, --igdbSupplement`: whether to supplement data with IGDB data
- `-s, --seedFile`: string seed data file to translate (CLZ collection XML export)
- `-w, --writeFileName` string filename to write JSON data to

## References

- https://api-docs.igdb.com/#getting-started
- https://clz.com/games
