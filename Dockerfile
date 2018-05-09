FROM iron/base

ENV LOGGLY_URL https://logs-01.loggly.com/inputs/%s/tag/http/
ENV LOGGLY_CUSTOMER_TOKEN YOUR_LOGGLY_TOKEN

EXPOSE 12123/udp
ADD gelffy-linux-amd64 /
CMD ["./gelffy-linux-amd64"]