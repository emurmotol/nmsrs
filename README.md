### NMSRS Lookup
National Manpower Skills Registration System Lookup For The Municipality Of Gasan Marinduque Philippines

### Tasks On Server
* Create application forms and backend logic
* Populate database
* ~~Migrate to CouchDB to support 32-bit host pc~~
* Missing on server Installing gomodifytags FAILED and Installing dlv FAILED

### Completed
* Fix [negroni] 2017-04-17T13:24:28+08:00 | 404 | 0s | localhost:8080 | GET /css/bootstrap.min.css.map
* Check laptop for virus, temp files, or whatever
* Fix VSCode extension "Go" not reloading (Delete package dir)
* ~~Clean your Google Chrome account~~
* Implement mgo usage
* Add struct annotations
* Process login
* Create authentication with JWT
* Parse templates
* [Run MongoDB server as a service](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/#configure-a-windows-service-for-mongodb-community-edition). Set `mongod --config "C:\Program Files\MongoDB\Server\3.2\bin\mongod.cfg" --journal --directoryperdb --dbpath=C:\Users\Public\OJT\MongoDB\db --storageEngine=mmapv1 --install`
* Add FontAwesome
* Use [Negroni](https://github.com/urfave/negroni)

### Notes
* go build & nmsrs-lookup.exe (PowerShell)
* go build ; ./nmsrs-lookup.exe (Bash)
* Server architecture is 32-bit must be 64-bit
* Max DB size 2GB
* VSCode debugger only works on 64-bit

### Important
* Always push source to remote

### Useful Links
* [validator](https://github.com/go-playground/validator)
* [couchdb-go](https://github.com/rhinoman/couchdb-go)
