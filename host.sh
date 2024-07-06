#!/bin/bash

SERVER_IP="127.0.0.1"
APACHE_DIR_PATH="/etc/apache2/sites-available"
BLUE='\033[0;34m'
DEFAULT='\033[0m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
RED='\033[0;31m'
SUDO="" 
HOST_NAME=cosign.org
DEFAULT_SLEEP=120

echo "Hosting...."

format_time(){
  local total_seconds="$1"
  local hours=$(( total_seconds / 3600 ))
  local remainder=$(( total_seconds % 3600 ))
  local minutes=$(( remainder / 60 ))
  local seconds=$(( remainder % 60 ))
  printf "%02d:%02d:%02d\n" "$hours" "$minutes" "$seconds"
}


if [[ $# -gt 0 ]]; then
  if [[ "$1" == "--sudo" ]]; then
    SUDO=sudo
  else
    if [[ "$1" =~ ^[0-9]+$ ]]; then
      sleep_duration="$1"
    else
      echo "Invalid sleep duration provided. Using default (120 seconds)."
      sleep_duration=$DEFAULT_SLEEP
    fi
  fi
else
  sleep_duration=$DEFAULT_SLEEP
fi

echo "Sleeping for $(format_time $sleep_duration)... For mysql db to start"
while [[ $sleep_duration -gt 0 ]]; do
  sleep_duration=$(( sleep_duration - 1 ))

  echo -ne "\033[K\r$(format_time $sleep_duration)"

  sleep 1
done

${SUDO} echo "Script execution continues...${SUDO}"



if ! which apache2 >/dev/null 2>&1; then
  echo "Apache2 is not installed. Installing..."
  ${SUDO} apt install -y apache2
else
  echo "Apache2 is already installed."
fi
${SUDO} a2enmod proxy proxy_http

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

${SUDO} touch ${APACHE_DIR_PATH}/${HOST_NAME}.conf
${SUDO} chmod 777 ${APACHE_DIR_PATH}/${HOST_NAME}.conf
${SUDO} echo "${INPUT}" >> ${APACHE_DIR_PATH}/${HOST_NAME}.conf
${SUDO} a2ensite ${APACHE_DIR_PATH}/${HOST_NAME}.conf
${SUDO} chmod 777 /etc/hosts
${SUDO} echo "${SERVER_IP}    ${HOST_NAME}" >> /etc/hosts
${SUDO} a2dissite 000-default.conf
${SUDO} apache2ctl configtest
${SUDO} systemctl restart apache2

echo "${GREEN}Apache Server Hosted${DEFAULT}"
CLEANDB=true ./pkg/config/migrations/migration
echo "CLeaning db.."

echo "Migrating Data.."
./pkg/config/migrations/migration
echo "Check ${BLUE} http://${HOST_NAME}${DEFAULT}"

./shelflove
