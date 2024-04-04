# from-dwca

Converts DwCA data into SQLite database according to sfgma schema.

## Install

Copy from the [latest release] compressed file for your operating system,
extract `from-dwca` file and place it somewhere in your `PATH`.

## Usage

Use DwCA file or URL to convert it to sqlite SQL dump:

```bash
from-dwca http://opendata.globalnames.org/dwca/174-mammal-sp-2018-08-04.tar.gz db.sql.zip
unzip db.sql.zip
## it will extract a sfga.sql file
```

To open file with sqlite

```bash
$ sqlite3 :memory:
SQLite version 3.45.2 2024-03-12 11:06:23
Enter ".help" for usage hints.
sqlite> .read sfga.sql
sqlite> .schema
```

[latest release]: https://github.com/sfborg/from-dwca/releases/latest


