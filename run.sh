#!/bin/sh

# Just to get Sudo Permissions
sudo echo "Running App in Background..."

# No stdout, only output Error to a log
# Run Process with Nice Value (Priority) of 1
$( $(sudo nice -n 1 ./build/app > /dev/null 2> stderr.log))&
disown
