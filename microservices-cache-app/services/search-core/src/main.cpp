#include <userver/clients/dns/component.hpp>
#include <userver/clients/http/component.hpp>
#include <userver/components/minimal_server_component_list.hpp>
#include <userver/server/handlers/ping.hpp>
#include <userver/ugrpc/server/component_list.hpp>
#include <userver/utils/daemon_run.hpp>

#include "dictionary_component.hpp"
#include "dictionary_handler.hpp"
#include "search_grpc_service.hpp"
#include "search_handler.hpp"

int main(int argc, char* argv[]) {
    auto component_list = userver::components::MinimalServerComponentList()
                               .Append<userver::server::handlers::Ping>()
                               .Append<userver::components::HttpClient>()
                               .Append<userver::clients::dns::Component>()
                               .AppendComponentList(userver::ugrpc::server::MinimalComponentList());

    search_core::AppendDictionaryComponent(component_list);
    search_core::AppendDictionaryCountHandler(component_list);
    search_core::AppendSearchHandler(component_list);
    component_list.Append<search_core::SearchGrpcService>();

    return userver::utils::DaemonMain(argc, argv, component_list);
}
