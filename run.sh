#!/bin/sh

# Just to get Sudo Permissions
sudo echo "Running App in Background..."

# No stdout, only output Error to a log
$( $(sudo ./build/app > /dev/null 2> stderr.log))&
disown
