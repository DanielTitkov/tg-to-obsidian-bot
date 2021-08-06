FROM golang:latest AS dependencies
ENV NAME "bot"
WORKDIR /opt/${NAME}
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM dependencies AS build
ENV NAME "bot"
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM alpine
ENV NAME "bot"
ARG TG_TOKEN
ARG NOTES_PATH
ARG TEMPLATE_PATH
ARG TEMPLATE_FILE
ARG TIMEOUT
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
CMD ./${NAME} \
    -timeout=${TIMEOUT} \
    -token=${TG_TOKEN} \
    -template=/opt/${NAME}/templates/ \
    -file=${TEMPLATE_FILE} \
    -path=/opt/bot/notes/