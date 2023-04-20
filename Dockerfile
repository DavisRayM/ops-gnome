FROM golang:1.20 AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /ops-server ./cmd/server

FROM ubuntu:22.10 AS release

ENV DEBIAN_FRONTEND noninteractive

WORKDIR /

COPY --from=build /ops-server /ops-server

RUN apt-get update

RUN apt-get install -y ca-certificates curl gnupg wget maven ssh

RUN mkdir ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

RUN curl -fsSLo /etc/apt/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

RUN curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | tee /usr/share/keyrings/helm.gpg

RUN echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | tee /etc/apt/sources.list.d/helm-stable-debian.list

RUN echo "deb [signed-by=/etc/apt/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | tee /etc/apt/sources.list.d/kubernetes.list


RUN apt-get update && apt-get install -y kubectl helm

RUN wget -q https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb
RUN apt-get --fix-broken install -y ./google-chrome-stable_current_amd64.deb

CMD ["/ops-server"]
