
SERVER_IP="127.0.0.1"
APACHE_DIR_PATH="/etc/apache2/sites-available"
BLUE='\033[0;34m'
DEFAULT='\033[0m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'

echo "Hosting...."
echo "Enter the ${YELLOW} hostname${DEFAULT}"
read -p "http://" HOST_NAME

echo "Your hostname: $HOST_NAME"

if ! which apache2 >/dev/null 2>&1; then
  echo "Apache2 is not installed. Installing..."
  sudo apt install apache2
else
  echo "Apache2 is already installed."
fi
sudo a2enmod proxy proxy_http

INPUT=$(cat <<EOF
<VirtualHost *:80>
	ServerName ${HOST_NAME}
	ServerAdmin cosign@mail.com
	ProxyPreserveHost On
	ProxyPass / http://127.0.0.1:8000/
	ProxyPassReverse / http://127.0.0.1:8000/
	TransferLog /var/log/apache2/mvc_access.log
	ErrorLog /var/log/apache2/mvc_error.log
</VirtualHost>
EOF
)

sudo touch ${APACHE_DIR_PATH}/${HOST_NAME}.conf
sudo chmod 777 ${APACHE_DIR_PATH}/${HOST_NAME}.conf
sudo echo "${INPUT}" >> ${APACHE_DIR_PATH}/${HOST_NAME}.conf
sudo a2ensite ${APACHE_DIR_PATH}/${HOST_NAME}.conf
sudo chmod 777 /etc/hosts
sudo echo "${SERVER_IP}    ${HOST_NAME}" >> /etc/hosts
sudo a2dissite 000-default.conf
sudo apache2ctl configtest
sudo systemctl restart apache2

echo "${GREEN}Apache Server Hosted${DEFAULT}"
echo "Check ${BLUE} http://${HOST_NAME}${DEFAULT}"
echo "Run ${YELLOW} make migrate${DEFAULT} to initialize the db if ${RED}NOT${DEFAULT}"
echo "Run ${YELLOW} make help${DEFAULT} for further-${RED}HELP${DEFAULT}"
