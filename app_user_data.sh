#!/bin/bash
set -eux

apt update -y
apt install -y unzip curl wget mysql-client awscli jq

cd /home/ubuntu

wget https://aws-tc-largeobjects.s3.us-west-2.amazonaws.com/CUR-TF-200-ACCAP1-1-91571/1-lab-capstone-project-1/s3/UserdataScript-phase-3.sh -O UserdataScript-phase-3.sh
chmod +x UserdataScript-phase-3.sh
bash UserdataScript-phase-3.sh || true

pkill node || true

cat > /usr/local/bin/start-student-app.sh <<EOF
#!/bin/bash
set -e

export APP_DB_HOST="${db_host}"
export APP_DB_USER="${db_user}"
export APP_DB_PASSWORD="${db_password}"
export APP_DB_NAME="${db_name}"
export APP_PORT=80

cd /home/ubuntu/resources/codebase_partner
npm start
EOF

chmod +x /usr/local/bin/start-student-app.sh

cat > /etc/systemd/system/studentapp.service <<'EOF'
[Unit]
Description=Student App
After=network.target

[Service]
Type=simple
User=ubuntu
ExecStart=/usr/local/bin/start-student-app.sh
Restart=always

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl enable studentapp
systemctl restart studentapp
