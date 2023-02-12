#!/bin/sh
echo -e "\e[34m >>> Starting the server \e[97m"
uvicorn server.main:app --host 0.0.0.0 --reload