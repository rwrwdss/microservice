#pragma once

#include <userver/components/component.hpp>

#include <keywords/search_service.usrv.pb.hpp>

#include "dictionary_component.hpp"

namespace search_core {

class SearchGrpcService final : public keywords::api::SearchServiceBase::Component {
public:
    static constexpr std::string_view kName = "handler-search-grpc";

    SearchGrpcService(const userver::components::ComponentConfig& config,
                       const userver::components::ComponentContext& context);

    SearchResult Search(CallContext&, keywords::api::SearchRequest&& request) override;

private:
    const DictionaryComponent& dictionary_;
};

}  // namespace search_core
