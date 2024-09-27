#!/bin/bash
# Скрипт для упрощенной заливки всего в main
DATE=`date +%d-%m-%Y`
# git pull
git add -A
git commit -m $DATE
git push -u origin main