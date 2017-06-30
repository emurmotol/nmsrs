### Run
* 32-bit (bash) `go build ; ./nmsrs.exe`
* 32-bit (cmd) `go build && nmsrs.exe`
* 64-bit press key `F5`

### Todo
* add validations on modal wizard if already exists prevent adding
* fix basicInfoHeightInFeet responsive stack
* select all not working
* ~~ren empStatRadios regist create.gohtml~~
* check row index if working
* fix login redirect when web server is reset on loggedInOnly middleware
* ~~set mongod --directoryperdb --config "C:\Program Files\MongoDB\Server\3.2\bin\mongod.cfg" --install~~
* must retrieve auth user using template func instead of passing as data
* create a flash middleware
* ~~ren javascript to script on gohtml~~
* ~~trim all space `strings.Replace(str, " ", "", -1)`~~
* ~~check if bootstrap 3.3.7 and jquery version are compat~~
* login_form btn-block animate not working properly
* ~~on input success remove help-block text~~
* ~~create pagination struct~~
* repopulate input file on post fail
* vendor dependencies on production
* ~~add onSuccess event on parsley and remove help-block span~~
* add error logs
* ~~handle redirects~~
* ~~set in php.ini for mysql client `post_max_size=80M`, `upload_max_filesize=200M`, `max_execution_time=5000`, `max_input_time=5000`, `memory_limit=1024M`~~
* ~~set id for .json seeded obj~~
* ~~register/create.gohtml: Calling validate on a parsley form without passing arguments as an object is deprecated~~
* ~~remove console.log()~~
* ~~update a table row instead of replacing it~~
* pref occ err ui not adding attr has-error
* ~~block-0 modal shows when enter key is pressed~~

### Learn
* https://github.com/disintegration/imaging
* https://github.com/go-gomail/gomail
* https://github.com/arnauddri/algorithms
* https://github.com/knq/xo
* https://github.com/dustin/go-humanize
* https://github.com/vulcand/oxy
* https://github.com/montanaflynn/stats
* https://github.com/egonelbre/gophers
* https://github.com/jung-kurt/gofpdf
* https://github.com/0xAX/go-algorithms
* https://github.com/MaxHalford/gago
* https://github.com/renstrom/fuzzysearch
* https://play.golang.org/p/DIQ2XtC8C4
* https://support.microsoft.com/en-us/help/196271/when-you-try-to-connect-from-tcp-ports-greater-than-5000-you-receive-the-error-wsaenobufs-10055
* http://parsleyjs.org/doc/annotated-source/defaults.html