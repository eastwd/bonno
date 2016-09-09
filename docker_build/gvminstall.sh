#!/bin/bash

curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer | bash
source /root/.gvm/scripts/gvm
gvm install go1.4
gvm use go1.4
gvm install go1.5.3
gvm use go1.5.3 --default
