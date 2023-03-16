FROM fedora
EXPOSE  80
COPY ./.bin ./.env ./
COPY ./configs ./configs
CMD ["/bot"]