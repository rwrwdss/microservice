cmake_minimum_required(VERSION 3.12...3.31)
project(search-core CXX)

find_package(userver REQUIRED COMPONENTS core grpc)

userver_add_grpc_library(${PROJECT_NAME}_proto
    PROTOS keywords/search.proto
    SOURCE_PATH ${CMAKE_CURRENT_SOURCE_DIR}/../../proto
)

add_library(${PROJECT_NAME}_objs OBJECT
    src/dictionary_component.hpp
    src/dictionary_component.cpp
    src/dictionary_handler.hpp
    src/dictionary_handler.cpp
    src/search_handler.hpp
    src/search_handler.cpp
    src/search_grpc_service.hpp
    src/search_grpc_service.cpp
)
target_link_libraries(${PROJECT_NAME}_objs PUBLIC userver::core userver::grpc ${PROJECT_NAME}_proto)

add_executable(${PROJECT_NAME} src/main.cpp)
target_link_libraries(${PROJECT_NAME} PRIVATE ${PROJECT_NAME}_objs)
