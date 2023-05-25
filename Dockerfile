FROM alpine:3.13

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
ADD env/sample.config env/config
ADD integrations/env/sample.config integrations/env/config
ADD migrations/sql migrations/sql
ADD locale/English.json locale/English.json
ADD locale/Indonesian.json locale/Indonesian.json

COPY brankas-disburse brankas-disburse
COPY migrate migrate

ENTRYPOINT ["sh", "-c", "./migrate up && ./brankas-disburse"]