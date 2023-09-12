#!/bin/bash

function setupDocker()
{
  if docker version; then
    echo "docker has been setup!"
  else
    sudo apt update
    sudo apt install docker.io
  fi
}


