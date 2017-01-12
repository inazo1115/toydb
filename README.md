# toydb

The toydb is a Relational Database Management System which isn't prectical.

## Unit test
```
$ make test
```

## Use REPL
```
$ make repl
```

## Example
```
$ make repl
version: 0.1.0
toydb> create table people (name string(20), age int);
ok
toydb> insert into people (name, age) values ('alice', 10);
ok
toydb> insert into people (name, age) values ('bob', 20);
ok
toydb> select * from people;
+---------+-----+
| name    | age |
+---------+-----+
| 'alice' |  10 |
| 'bob'   |  20 |
+---------+-----+
```
