FROM gcc:bullseye as builder

WORKDIR /app
COPY main.c main.c
RUN gcc main.c -o main

FROM debian:bullseye

WORKDIR /app
COPY --from=builder /app/main main
CMD ["./main"]