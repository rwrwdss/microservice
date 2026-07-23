FROM ghcr.io/userver-framework/ubuntu-22.04-userver:latest

WORKDIR /app/services/search-core

COPY proto /app/proto
COPY services/search-core/CMakeLists.txt ./
COPY services/search-core/src ./src

RUN cmake -B build -DCMAKE_BUILD_TYPE=Release \
    && cmake --build build -j "$(nproc)" --target search-core

COPY services/search-core/configs ./configs

CMD ["/app/services/search-core/build/search-core", "--config", "/app/services/search-core/configs/static_config.yaml", "--config_vars", "/app/services/search-core/configs/config_vars.yaml"]
