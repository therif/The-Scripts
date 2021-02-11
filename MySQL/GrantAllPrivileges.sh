#!/bin/sh

mysql -e "GRANT ALL PRIVILEGES ON *.* TO 'therif'@'%';FLUSH PRIVILEGES;"
