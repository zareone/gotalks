FROM golang:1.11

# install present tool
RUN go get golang.org/x/tools/cmd/present

# add slides 
# ADD . /slides

VOLUME /slides

# setup service
WORKDIR /slides

EXPOSE 3999
EXPOSE 8080

ENTRYPOINT ["/go/bin/present"]
CMD ["-http=:3999"]
