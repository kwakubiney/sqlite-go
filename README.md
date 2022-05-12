# sqlite-go
A naïve implementation of a persistent disk database using Golang.

# Demo

![ezgif-5-dc6eb5f6d1](https://user-images.githubusercontent.com/71296367/168181177-3a1e3945-e187-4817-8aff-8cd2a3c78f2b.gif)


# How was it built?

1) Basically, I utilize an append only file named `db` and an index file named `index` which gets created when the DB is opened by running `go run main.go`

2) When inserts of rows are done, I maintain a `map` in memory to store the row `ID` as the key and the byte offset of the encoded row as the value for faster lookups during a `select`

3) Index map is flushed to disk after the app is exited and loaded back in memory when the DB is up running again.

# How was serialization done?
1) I encode strings using the format `<ID>:<Name>:<Email>` along with their lengths (in `32` bits, `big endian` format) during a single insert.

2) The encoded lengths help to seek through the file to read specific rows during a `select`
# Why?
To learn how disk based databases and indexes work.

# How to use
This database supports `insert` and `select` of rows with fields `ID, Name, Email`

# How to do an insert

`insert <ID> <Name> <Email>`

`example` : `insert 1 Kwaku k@mail.com`

# How to do a select of a specific row

`select <ID>`

`example` : `select 1`

# How to do a select all rows

`select`

`example` : `select`

# Things I will do in the future

1) There is no control on how large the DB file can grow so compaction must be done to limit the growth.
2) Implement `~4KB` paging to make DB reads on disk faster and more efficient.
3) Concurrency and all the other cool stuff

...when I understand them enough
   
