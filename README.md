<a href="https://www.velocityworks.io/home">Velocity Works Coding Demo</a>
# Golang-Demo

This Golang application consumes a JSON payload from [data.gov](https://www.data.gov/), populates [expense Actuals](https://catalog.data.gov/dataset/expense-actuals/resource/3ba4662e-1a4c-4121-8590-2e6c6961b7a4) database and displays the database contents on a web page.

Frameworks used:

- Echo to build the website: https://github.com/labstack/echo
- Resty to retrieve data from Data.gov: https://github.com/go-resty/resty
- Jsonparser to process the data: https://github.com/buger/jsonparser
- GORM to populate the database: https://github.com/jinzhu/gorm


## Setup

### Configuration file

Create a velocity_works.ini configuration file and place it wherever you want. 
Configuration should look like this
```
[main]
templates_directory = /home/hostname/path/to/proj-dir/Golang-Demo/contrib/templates


[database]
host = localhost
name = velocity_works
password =postgres
port = 5432
user = username
```

### Export config path

put this line into ~/.bashrc file

```
export VELOCITY_WORKS_SETTINGS=path/to/config/velocity_works.ini 
```

**change** ``path/to/config/`` with your config path.

### Run migration

in order to run migration you need to install [golang-migrate/migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

after installing run this command from project's directory
```
 migrate -path ./contrib/migrations/ -database "postgres://USERNAME:@localhost:5432/DBNAME?sslmode=disable" up
```

**change** USERNAME and DBNAME with you username and dbname

### Build 
Run make command it will take care of rest
```
make
```

### Help

``./bin/velocity-worker --help``

### Sync database

``./bin/velocity-worker sync_db``

### Serve UI api

``./bin/velocity-worker serve_ui``

