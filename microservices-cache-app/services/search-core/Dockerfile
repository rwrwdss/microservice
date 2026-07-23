FROM ghcr.io/userver-framework/ubuntu-22.04-userver:latest

WORKDIR /app

COPY CMakeLists.txt ./
COPY src ./src

RUN cmake -B build -DCMAKE_BUILD_TYPE=Release \
    && cmake --build build -j "$(nproc)" --target search-core

COPY configs ./configs

CMD ["/app/build/search-core", "--config", "/app/configs/static_config.yaml", "--config_vars", "/app/configs/config_vars.yaml"]
