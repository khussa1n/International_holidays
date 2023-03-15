FROM fedora
EXPOSE  80
COPY . .
CMD [".bin/bot"]