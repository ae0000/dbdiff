DBDiff
================================================

Version 0.2
-----------



DBDiff takes two (MySQL) database structures and works out the differences between the 
two. It then uses these diffs to create a list of SQL updates you would need to 
make to align them.

Notes:
------
The following things are not yet handled.

- [ ] triggers
- [ ] stored procedures
- [ ] views
- [ ] other dbs besides mysql?

Comments are a little dicey as well.... 

How to use:
--------------

Firstly get your DB structures, something like this:
```
mysqldump --user tt -p  \
--comments=FALSE \
--compact=TRUE \
--routines=TRUE \
--no-data \
--triggers=TRUE \
databasename \
 | sed -e 's/DEFINER[ ]*=[ ]*[^*]*\*/\*/' > original.sql
```

You should have _original.sql_ and _current.sql_.

Then install this thing..
```
go get github.com/ae0000/dbdiff
cd to/where/ever/the/dbdiff/dir/ends/up
go install
```

Then to actually run the diff:
```
dbdiff diff --original original.sql --current current.sql
```

You should then get something like this printed out:
```
-- :::::::::: Changes for  Addresses
-- .....................................................
-- Column   :  Town
-- Original :  char(128) NOT NULL
-- New      :  char(128) NOT NULL DEFAULT ''
ALTER TABLE `Addresses` MODIFY COLUMN Town char(128) NOT NULL DEFAULT '';
-- .....................................................
-- Column   :  Description
-- Original :  text NOT NULL
-- New      :  mediumtext NOT NULL
ALTER TABLE `Addresses` MODIFY COLUMN Description mediumtext NOT NULL;

```

