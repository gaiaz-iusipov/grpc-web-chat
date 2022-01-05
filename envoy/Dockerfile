FROM envoyproxy/envoy-alpine:v1.20.1
COPY envoy.yaml /etc/envoy.yaml
CMD ["envoy", "--config-path", "/etc/envoy.yaml", "--log-level", "warning"]
