FROM golang:alpine

WORKDIR /app

COPY . /app

RUN go install

CMD ["/go/bin/kubernetes-pod-scaler"]