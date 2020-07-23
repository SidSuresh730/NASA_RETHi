# NASA RETHi

### Database

MySQL should be installed on your machine.
Once MySQL is installed, use the latest SQL dump (mcvt-$(date_time).sql) to create the mcvt database.
The web server uses the root user to access the database.
Make sure that the root user uses port 3306 to access the database.
Set the password for the root user to an environment variable called "DBPasswd"

### Web server

Make sure a Go (Golang) distribution is installed on the machine
Do a git pull to install a copy of the NASA RETHi repository on the local machine
Build using the command ```$ go build``` in the ~/web folder
This will create an executable file called *web* (Linux) or *web.exe* (Windows)
Run this executable file to start the web service

### Front end

Currently the front end can only be deployed on a Linux machine.
This can be pulled from the *linux_deployment* branch
The web server can be run without the front end but it will only display raw JSON data
