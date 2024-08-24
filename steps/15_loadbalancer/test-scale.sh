#!/bin/bash

# Function to start servers
start_servers(){
    for((i=1;i<=3;i++));do
        SERVER_ADDR=":808$i" go run main.go serve &
        sleep 2
    done
}

# Function to trap SIGINT and kill child processes
trap_and_kill(){
    echo "Stopping servers..."
    pkill -P $$
}

# Start servers
start_servers

go run main.go loadbalancer 8081 8082 8083
