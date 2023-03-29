FROM alpine:latest

COPY dist/luogu_stats_card_linux_amd64 /app/luogu_stats_card

WORKDIR /app

EXPOSE 10127

CMD ["./luogu_stats_card"]