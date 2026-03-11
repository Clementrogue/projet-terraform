#!/bin/bash

apt update -y

cd /home/ubuntu

wget https://aws-tc-largeobjects.s3.us-west-2.amazonaws.com/CUR-TF-200-ACCAP1-1-91571/1-lab-capstone-project-1/s3/UserdataScript-phase-2.sh

chmod +x UserdataScript-phase-2.sh

./UserdataScript-phase-2.sh
