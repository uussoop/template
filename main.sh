#!/bin/bash
# set -x

# initiat any other program ie docker compose in detached mode
servicename="myapp"
mypath=$(pwd)
description="doing ai stuff"

#create a service
if [ -f "/lib/systemd/system/$servicename.service" ]; then
    echo "Service $servicename already exists"
else
    echo "Creating service $servicename"
    echo "[Unit]
    Description=$description

    [Service]
    Type=simple
    ExecStart=/bin/bash ${mypath}/read.sh
    Restart=always
    RestartSec=10

    [Install]
    WantedBy=multi-user.target" > /lib/systemd/system/$servicename.service
fi


# touch /lib/systemd/system/$servicename.service 