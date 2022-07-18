apt update
apt install mysql-server
mysql_secure_installation
nano /etc/mysql/mysql.conf.d/mysqld.cnf
mysql
GRANT ALL PRIVILEGES ON *.* TO 'room'@'*' IDENTIFIED BY 'therif' WITH GRANT OPTION;
FLUSH PRIVILEGES;
exit
