FROM alpine

COPY app-main /app-main
COPY web /web
RUN chmod 755 /app-main

EXPOSE 8079

CMD ["/app-main"]
