# backend
you run, build, and test everything through a makefile

```bash
cp ./.env.example ./.env # then edit the file to set everything
make run
```

pls do NOT use chatgippity to genuinely make code beyong human comprehension
group concepts by module broadly

the entry is main.go in the root

- data/ stores sqlite databases and their extra stuff (wal, locks)
- bruno/ is for bruno requests
- mod db/: responsible for handling data through gorm, also contains methods for interacting with that data
- mod handlers/:
    http handlers. every handler should be a separate file with swagger docs on top

middleware should be stored inside relevant modules

some notes:
- the database is passed around using a context

architecture things to solve:

since every instance is separate and relatively small, they should all have their own sqlite databases.
do you just open a database dynamically when a user connects or store them all in a map and close on a cron job?

tap in with reverse proxies

## swagger
didnt go with it, it was kinda silly, just use bruno at bruno/

but if it happens at some point
```
go install github.com/swaggo/swag/cmd/swag@latest
```

docs are at [](http://localhost:3333/swagger/index.html/)
you will have to regenerate them using `make swag`
