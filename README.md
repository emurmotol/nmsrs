### NMSRS
National Manpower Skills Registration System For The Municipality Of Gasan Marinduque Philippines

### Tasks On Server
* CRUD for registrant
* Send email
* Search all
* Add CSRF token on forms
* Finish downgrade markup from BS4 - BS3
* Add parsley to each form
* Fix back button after logout
* Change controllers status 500 to panic
* Remove unused structs
* Validate on client side
* Remove go-validator look for alternative
* Remove validate tag on structs

### Completed
* Remove Errors field in JSON struct use Data field to store errors
* ~~Missing on server Installing gomodifytags FAILED and Installing dlv FAILED~~
* Load comboBox values
* Seed database
* Create application forms and backend logic
* ~~Migrate to CouchDB to support 32-bit host pc~~
* Screenshot adding of registrant
* JWT Middleware: f path is / and not authenticated set / to a different view'
* Fix [negroni] 2017-04-17T13:24:28+08:00 | 404 | 0s | localhost:8080 | GET /css/bootstrap.min.css.map
* Check laptop for virus, temp files, or whatever
* Fix VSCode extension "Go" not reloading (Delete package dir)
* ~~Clean your Google Chrome account~~
* Implement mgo usage
* Add struct annotations
* Process login
* Create authentication with JWT
* Parse templates
* ~~[Run MongoDB server as a service](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-windows/#configure-a-windows-service-for-mongodb-community-edition). Set `mongod --config "C:\Program Files\MongoDB\Server\3.2\bin\mongod.cfg" --journal --directoryperdb --dbpath=C:\Users\Public\OJT\MongoDB\db --storageEngine=mmapv1 --install`~~
* Add FontAwesome
* Use [Negroni](https://github.com/urfave/negroni)

### Notes
* go build & nmsrs.exe (PowerShell)
* go build ; ./nmsrs.exe (Bash)
* git fetch origin && git reset --hard origin/master && git clean -f -d
* Server architecture is 32-bit must be 64-bit
* Max DB size 2GB
* VSCode debugger only works on 64-bit

### Important
* Change key on production
* Always push source to remote

### Useful Links
* https://github.com/oliamb/cutter
* https://www.socketloop.com/tutorials/golang-crop-image
* https://github.com/tdewolff/minify
* http://www.chartjs.org/docs/
* https://www.google.com.ph/url?sa=t&rct=j&q=&esrc=s&source=web&cd=5&cad=rja&uact=8&ved=0ahUKEwjts_f3n-XTAhXMybwKHeeTDQ0QFgg8MAQ&url=http%3A%2F%2Fjsfiddle.net%2Fmilz%2FryQu5%2F1%2F&usg=AFQjCNFdjklYQzRr4p0SjxmdnFW9xXTlEQ&sig2=E_9Z9O9OXihfGpIeqAvcPw
* http://stackoverflow.com/questions/24844591/how-to-configure-parsley-remote-validator-to-handle-my-api-response
* http://stackoverflow.com/questions/22217735/getting-parsley-2-x-working-with-bootstrap-3-correctly
* https://godoc.org/github.com/mitchellh/mapstructure#example-Decode
* http://stackoverflow.com/questions/18487056/select2-doesnt-work-when-embedded-in-a-bootstrap-modal/19574076#19574076
* https://github.com/twitter/typeahead.js
* https://mholt.github.io/json-to-go/
* http://codepen.io/cfmatre/pen/LGOdjq
* http://stackoverflow.com/questions/19448939/how-to-disable-browser-from-regenerating-my-form-data-upon-pressing-back-button
* https://github.com/go-playground/validator
* https://github.com/rhinoman/couchdb-go
* https://stackoverflow.com/questions/17129797/golang-how-to-check-multipart-file-information
* https://github.com/asaskevich/govalidator
* https://github.com/moxiecode/plupload
* http://stackoverflow.com/questions/17129797/golang-how-to-check-multipart-file-information
* https://www.socketloop.com/tutorials/golang-convert-png-transparent-background-image-to-jpg-or-jpeg-image
* https://www.socketloop.com/tutorials/golang-save-image-to-png-jpeg-or-gif-format
* http://stackoverflow.com/questions/12430874/image-manipulation-in-golang
* http://stackoverflow.com/questions/33168973/converting-multipart-file-to-an-image-object-in-golang
