FROM jupyter/scipy-notebook

USER root
RUN apt update -y 
RUN apt install libmysqlclient-dev -y
USER $NB_UID