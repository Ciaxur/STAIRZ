#!/bin/sh

# No stdout, only output Error to a log
$( $(sudo ./build/app > /dev/null 2> stderr.log))&
disown
