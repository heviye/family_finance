FROM alpine:3.6
ADD family_finance /
EXPOSE 8010
CMD ["/family_finance"]