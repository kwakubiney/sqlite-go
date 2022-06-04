# http-filedb
A naïve implementation of a file storage which can be served over REST and terminal `stdin` solely for learning purposes and not for production.

# How was the storage engine built?

1) `http-filedb` uses an append only file and an index file.

2) When inserts of rows are done, I maintain a `map` in memory to store the row `ID` as the key and the byte offset of the encoded row as the value for faster lookups during a `select`

3) Index map is flushed to disk after the app is exited and loaded back in memory when the DB is up running again.

# How was serialization done?
1) Data is stored using the format `<ID>:<Name>:<Email>` along with their lengths (in `32` bits, `big endian` format).

2) The encoded lengths help to seek through the file to read specific rows.

# Why?
To learn the basics of serialization in disk databases, how indexes work and because I can.

# How to use?
Check out the API documentation.

# Flaws of this design

1) There is no control on how large the DB file can grow.
2) I use mutexes on the file to ensure one thread writes and reads at a time but atomicity is not guaranteed.
3) i use a default schema for writes:
```
struct Row {
    ID string
    Username string
    Email string
}
```


   
