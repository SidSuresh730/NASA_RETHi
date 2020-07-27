# NASA RETHi

### Database

MySQL should be installed on your machine.
1. To install MySQL follow the instructions for your specific operating system here: https://dev.mysql.com/doc/mysql-installation-excerpt/8.0/en/
The main requirements are the MySQL Server and MySQL Workbench, the other features, while helpful, may not be necessary. The default settings can be used for the rest of the installation steps.
    1. During installation, you should be prompted to create a password for the **root** user of the MySQL database. By default the root user connects to the server using tcp/ip port **3306**.
    1. Create an environment variable called **DBPasswd** and store the password in this variable. The web server looks for this environment variable and uses it when accessing the database.
1. Create a folder called ‘NASA_RETHi’ or similar and use the command ```git init``` to initialize the folder as an empty Git repository. If Git is not installed on your machine, follow the instructions here: https://git-scm.com/book/en/v2/Getting-Started-Installing-Git.
If working with a Windows machine, the **Git Bash** is a useful tool to have as it provides a more “Linux-like” interface.
1. Once the Git repository is initialized, go to https://github.com/SidSuresh730/NASA_RETHi and press the green “Code” button. Select the option to clone with HTTPS and copy the URL.
1. Use the command ```git pull <URL> ``` to copy the remote Git repository into the local empty Git repository
    1. If working on a Linux machine, use the command ```git pull <URL> linux_deployment ``` to copy the repository. It includes the front-end.
1. Now that the Git repository is copied into your folder, there should be multiple **.sql** files. These begin with *mcvt-* followed by a date and time. Enter the MySQL Workbench. There should already be a connection for the root user, but if not, a connection can be created by pressing the **+** button next to **MySQL Connections**. Using this connection to access the server, go to *File --> Run SQL Script...* select the **.sql** file from ~/NASA_RETHi/ folder with the latest time. Then in the next window click **Run** to create the **mcvt** database on your local machine.

### Web server

Make sure a Go (Golang) distribution is installed on the machine.
Go can be installed by following the instructions here: https://golang.org/doc/install
You also need to install the necessary git repositories.
##### Preparing the dev environment
1. Set your **$GOPATH** environment variable to the location of your **Go/** folder
1. Add **$GOPATH/bin** to your **PATH**
1. In the **$GOPATH/src/** folder, make a directory called **github.com**
    1. Inside the **github.com** folder make a folder called **AmyangXYZ** and inside this folder make a folder called **sweetygo** and execute this statement ```git pull https://github.com/AmyangXYZ/sweetygo.git ```
    1. Inside **github.com** make another folder called **dgrijalva** and inside this folder make a folder called **jwt-go** and execute this statement ```git pull https://github.com/dgrijalva/jwt-go.git ```
    1. Inside **github.com** make another folder called **go-sql-driver** and inside this folder make a folder called **mysql** and execute this statement ```git pull https://github.com/go-sql-driver/mysql.git ```
1. Build using the command ```$ go build``` in the ~/web folder
This will create an executable file called *web* (Linux) or *web.exe* (Windows)
1. Run this executable file to start the web service
    1. ```$ ./web ``` in Linux or ``` > web.exe ``` in Windows

### Front end

Currently the front end can only be deployed on a Linux machine.
The following instructions are **LINUX ONLY**
1. In the ~/web folder there should be two files: **dist.tar.xz** and **deploy.sh**.
1. Run deploy.sh: ```$ ./deploy.sh ```
1. Restart the web service.

The web service should be running at *< Your IP >:8888*. If running on your local machine it runs at *localhost:8888*
If running on a Windows machine without front end, use these paths to observe JSON data for each subsystem:
* /api/eclss/wrs
* /api/eclss/ogs
* /api/eclss/fms
* /api/c2/decisions
* /api/human/agents
* /api/human/setpoints
* /api/interior/temperature
* /api/interior/humidity
* /api/interior/pressure
* /api/structural/temperature
* /api/structural/strain
* /api/structural/pressure
* /api/structural/acceleration
* /api/structural/displacement
* /api/structural/damageinfo
* /api/robot/agents
* /api/robot/interventions
* /api/power/powergeneration
* /api/power/powerconsumption
* /api/power/storage
* /api/power/health
* /api/inventory
* /api/external/disturbances
