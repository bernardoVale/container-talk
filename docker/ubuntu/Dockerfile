FROM ubuntu:latest

RUN useradd -m -u 1000 learn

RUN apt-get update && apt-get -y install vim

ADD bash_profile /home/learn/.bash_profile
