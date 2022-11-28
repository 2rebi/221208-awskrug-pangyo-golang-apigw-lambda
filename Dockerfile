FROM golang:1.18-alpine3.16 as builder

RUN apk update && apk upgrade && \
    apk --update add make git gcc g++ zip

WORKDIR /app

COPY . .

RUN make install
RUN make build

FROM scratch as distribution

WORKDIR /

COPY --from=builder /app/1_helloworld/function.zip 1_helloworld.zip
COPY --from=builder /app/2_apigw_lambda/function.zip 2_apigw_lambda.zip
COPY --from=builder /app/3_apigw_lambda_route/function.zip 3_apigw_lambda_route.zip
COPY --from=builder /app/4_case_http_net/function.zip 4_case_http_net.zip
COPY --from=builder /app/5_case_echo_framework/function.zip 5_case_echo_framework.zip


