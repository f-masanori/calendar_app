FROM python:3.7
USER root

RUN apt-get update && apt-get -y install locales && \
    localedef -f UTF-8 -i ja_JP ja_JP.UTF-8
RUN apt-get install -y vim
RUN apt-get install  -y python-dev  
RUN apt-get install -y default-libmysqlclient-dev
RUN apt-get install -y python3-dev
RUN pip install --upgrade pip
RUN pip install requests==2.18.4 mysql-connector-python 
RUN pip install python-dotenv

ENV LANG ja_JP.UTF-8
ENV LANGUAGE ja_JP:ja
ENV LC_ALL ja_JP.UTF-8
ENV TZ JST-9
