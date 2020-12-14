#!/bin/sh

cp -p /etc/hosts /etc/hosts.back
sed '/yourhost.yourdyndns.com/d' /etc/hosts > /etc/hosts.new
cp -p /etc/hosts.new /etc/hosts
getent hosts yourhost.yourdyndns.com >> /etc/hosts
