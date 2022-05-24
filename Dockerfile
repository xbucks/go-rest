FROM golang:latest AS build


WORKDIR /app

# copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY . .
RUN make build


FROM scratch
WORKDIR /app/
ARG PORT
COPY --from=build /app/go-rest-api-example .
COPY --from=build /app/config/*.yaml /app/config/
CMD ["./go-rest-api-example"]
EXPOSE $PORT

