# NASA RETHi

### Database

MySQL should be installed on your machine.
Once MySQL is installed, use the latest SQL dump (mcvt-$(date_time).sql) to create the mcvt database.
The web server uses the root user to access the database.
Make sure that the root user uses port 3306 to access the database.
Set the password for the root user to an environment variable called "DBPasswd"

### Web server

Make sure a Go (Golang) distribution is installed on the machine
You also need to install the necessary git repositories.
##### Preparing the dev environment
1. Set your **$GOPATH** environment variable to the location of your **Go/** folder
1. Add **$GOPATH/bin** to your **PATH**
1. In the **$GOPATH/src/** folder, make a directory called **github.com**
    1. Inside the **github.com** folder make a folder called **AmyangXYZ** and inside this folder make a folder called **sweetygo** and execute this statement ```git pull https://github.com/AmyangXYZ/sweetygo.git ```
    1. Inside **github.com** make another folder called **dgrijalva** and inside this folder make a folder called **jwt-go** and execute this statement ```git pull https://github.com/dgrijalva/jwt-go.git ```
    1. Inside **github.com** make another folder called **go-sql-driver** and inside this folder make a folder called **mysql** and execute this statement ```git pull https://github.com/go-sql-driver/mysql.git ```
1. Do a ```git pull``` to install a copy of the NASA RETHi repository on the local machine
1. Build using the command ```$ go build``` in the ~/web folder
This will create an executable file called *web* (Linux) or *web.exe* (Windows)
1. Run this executable file to start the web service

### Front end

Currently the front end can only be deployed on a Linux machine.
This can be pulled from the *linux_deployment* branch
The web server can be run without the front end but it will only display raw JSON data
