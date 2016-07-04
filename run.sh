#!/bin/bash
echo 24 > /sys/class/gpio/unexport
/opt/golang/src/github.com/stunti/reboot/main3

