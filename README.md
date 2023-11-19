# ddx - DynamoDB Key Schema Viewer

`ddx` is a CLI tool designed to help you quickly and conveniently view the key
schemas of your DynamoDB tables and their indexes. With `ddx`, you can get a
summary of the partition keys and sort keys for your DynamoDB tables, with
support for various output formats.

## Features

- View key schema of DynamoDB tables and indexes.
- Output in human-readable text or machine-readable JSON format.
- Option to view key schema of a specific table.
- Ability to limit the number of tables listed.

## Installation

```sh
go install github.com/kamilturek/ddx
```

## Usage

```sh
ddx --help
A tool for quickly viewing your DynamoDB table and index key schemas

Usage:
  ddx [flags]

Flags:
  -a, --all                 list all tables
  -f, --format string       output format, available: "text", "json" (default "text")
  -h, --help                help for ddx
  -l, --limit int           maximum number of tables listed (default -1)
  -t, --table-name string   table name whose key schema should be shown
```

## Examples

View key schema for a specific table:

```sh
$ ddx --table-name Users
Table Name:     Users
Table Type:     BASE
Parition Key:   UserID
```

View key schemas for all tables in text format:

```sh
$ ddx --all
Table Name:     Books
Table Type:     BASE
Parition Key:   AuthorID
Sort Key:       BookID

Table Name:     BooksLSI
Table Type:     LSI
Partition Key:  AuthorID
Sort Key:       PublisherID

Table Name:     BooksGSI
Table Type:     GSI
Partition Key:  BookID

Table Name:     Users
Table Type:     BASE
Parition Key:   UserID
```

Output key schemas in JSON format:

```sh
$ ddx --all --format json
[
    {
       "TableName": "Books",
       "TableType": "BASE",
       "PartitionKey": "AuthorID",
       "SortKey": "BookID"
    },
    {
       "TableName": "BooksLSI",
       "TableType": "BASE",
       "PartitionKey": "AuthorID",
       "SortKey": "PublisherID"
    },
    {
       "TableName": "BooksGSI",
       "TableType": "BASE",
       "PartitionKey": "BookID"
    },
    {
       "TableName": "Users",
       "TableType": "BASE",
       "PartitionKey": "UserID"
    }
]
```

## License

`ddx` is open-sourced software licensed under the [MIT license](LICENSE).
