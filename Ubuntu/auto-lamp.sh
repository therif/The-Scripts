#!/usr/bin/env bash

#UBUNTU VERSION ARRAYS
VERSIONARRAY=(16.04 16.04.1 16.04.2 16.04.3 16.04.4 16.04.5 16.04.6 18.04 18.04.1 18.04.2)
SIXTEENARRAY=(16.04 16.04.1 16.04.2 16.04.3 16.04.4 16.04.5 16.04.6)
EIGHTEENARRAY=(18.04 18.04.1 18.04.2)

#CHECKING UBUNTU VERSION
echo -e "######## \e[32mCHECKING UBUNTU VERSION\e[39m ########"
VERSION=$(lsb_release -rs)

if [[ ! " ${VERSIONARRAY[@]} " =~ " ${VERSION} " ]]; then
    echo -e ""
    echo -e "This script does nopt support your OS/Distribution."
    echo -e "Goodbye!"
    echo -e ""
    exit 1
fi

echo -e ""
echo -e "Found Ubuntu Version: $VERSION"

#INSTALLING UPDATES
sudo apt-get update -y -qq > /dev/null
echo -e ""

#INSTALLING APACHE2
echo -e "######## \e[32mINSTALLING APACHE2\e[39m ########"
sudo apt install -y apache2 > /dev/null

if [[ " ${SIXTEENARRAY[@]} " =~ " ${VERSION} " ]]; then
    #CONFIGURING APACHE FOR UBUNTU 16
    echo -e "######## \e[32mCONFIGURING APACHE FOR UBUNTU 16\e[39m ########"
    echo -e ""
    read -s -p "Enter your Domain Name or Press ENTER to skip: " servername
    echo -e ""
    if [ -z "$servername" ];
    then
    echo -e "Skipping setting Server Name in Apache Config."
    else
    echo "ServerName $servername" >> /etc/apache2/apache2.conf
    fi
    echo -e ""
fi

#CONFIGURING UFW
echo -e "######## \e[32mCONFIGURING UFW\e[39m ########"
echo -e ""
sudo ufw allow in "Apache Full"
echo -e ""

#INSTALLING MYSQL
if [[ " ${SIXTEENARRAY[@]} " =~ " ${VERSION} " ]]; then
    echo -e "######## \e[32mINSTALLING MYSQL FOR UBUNTU 16\e[39m ########"
    echo -e ""
    read -s -p "What Password Do You Want To Set For Mysql?: " mysqlpass
    echo -e ""
    echo "mysql-server mysql-server/root_password password $mysqlpass" | sudo debconf-set-selections
    echo "mysql-server mysql-server/root_password_again password $mysqlpass" | sudo debconf-set-selections
    sudo apt-get install -y mysql-server -qq > /dev/null
    echo -e ""
fi

if [[ " ${EIGHTEENARRAY[@]} " =~ " ${VERSION} " ]]; then
    echo -e "######## \e[32mINSTALLING MYSQL FOR UBUNTU 18\e[39m ########"
    echo -e ""
    read -s -p "What Password Do You Want To Set For Mysql?: " mysqlpass
    echo -e ""
    sudo apt install -y mysql-server -qq > /dev/null
    mysql -e "ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '$mysqlpass';"
fi

#CONFIGURING MYSQL
echo -e "######## \e[32mCONFIGURING MYSQL\e[39m ########"
echo -e ""
mysql -u root -p$mysqlpass -e "DELETE FROM mysql.user WHERE User='';"
mysql -u root -p$mysqlpass -e "DELETE FROM mysql.user WHERE User='root' AND Host NOT IN ('localhost', '127.0.0.1', '::1');"
mysql -u root -p$mysqlpass -e "DROP DATABASE IF EXISTS test;"
mysql -u root -p$mysqlpass -e "DELETE FROM mysql.db WHERE Db='test' OR Db='test\\_%';"
mysql -u root -p$mysqlpass -e "FLUSH PRIVILEGES;"
echo -e ""

#INSTALLING PHP
echo -e "######## \e[32mINSTALLING PHP\e[39m ########"
sudo apt install -y php libapache2-mod-php php-mysql -qq > /dev/null
sed -i "s/index.html index.cgi/index.php index.html index.cgi/g" /etc/apache2/mods-enabled/dir.conf
sed -i "s/index.php index.xhtml/index.xhtml/g" /etc/apache2/mods-enabled/dir.conf
sudo systemctl restart apache2

#INSTALLATION FINISHED
echo -e "######## \e[32mINSTALLATION FINISHED\e[39m ########"
echo -e ""
echo -e "Congratulations, installion is complete!"
echo -e ""
