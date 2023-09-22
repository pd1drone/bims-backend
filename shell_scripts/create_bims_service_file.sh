#!/bin/bash

# Create the file bims_backend.service with the specified content
cat << EOF > /etc/systemd/system/bims_backend.service
[Unit]
Description=bims backend http server
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/bims-backend/
ExecStart=/root/bims-backend/cmd/bims
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the bims_backend.service file so it starts on boot
sudo systemctl enable bims_backend.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the bims_backend.service
sudo systemctl start bims_backend.service
