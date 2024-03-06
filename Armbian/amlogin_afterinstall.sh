#!/bin/bash

# My Hello, World! script
echo "Hello, World! - from The-Scripts"

# Disable Armbian Log
systemctl stop armbian-ramlog.service
systemctl disable armbian-ramlog.service

#Update Pakcage
sudo apt update -y

#Upgrade and replace
sudo apt upgrade -y

#Then run manual NMTUI to set network
echo "Run on terminal with command: nmtui"
echo "Jangan lupa bagian ipv4 set manual"
echo "-- isi IP, Gateway dan dns (Samakan dgn gateway)"

