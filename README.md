# repo structure
you run, build, and test everything through a makefile

pls do NOT use chatgippity to genuinely make code beyong human comprehension
group concepts by module broadly

the entry is main.go in the root

- data/ stores sqlite databases and their extra stuff (wal, locks)
- bruno/ is for bruno requests
- mod db/: responsible for handling data through gorm, also contains methods for interacting with that data
- mod handlers/:
    http handlers. every handler should be a separate file with some documentation on top (should prolly switch to swagger at some point idk)
    files should be names as their route and method to compensate for go's handler initialization syntax being very non-descriptive ("GET /poll/:id" is "poll_id_get.go", both the request and upgraded websocket at "/polls/:id/live should be "poll_live_ws.go")

middleware should be stored inside relevant modules

some notes:
- the database is passed around using a context

architecture things to solve:

since every instance is separate and relatively small, they should all have their own sqlite databases.
do you just open a database dynamically when a user connects or store them all in a map and close on a cron job?

tap in with reverse proxies
