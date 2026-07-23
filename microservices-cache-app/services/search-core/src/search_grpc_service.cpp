#include "search_grpc_service.hpp"

namespace search_core {

namespace {
constexpr std::size_t kDefaultLimit = 10;
}

SearchGrpcService::SearchGrpcService(const userver::components::ComponentConfig& config,
                                      const userver::components::ComponentContext& context)
    : keywords::api::SearchServiceBase::Component(config, context),
      dictionary_(context.FindComponent<DictionaryComponent>()) {}

SearchGrpcService::SearchResult SearchGrpcService::Search(CallContext&, keywords::api::SearchRequest&& request) {
    const auto limit = request.limit() > 0 ? static_cast<std::size_t>(request.limit()) : kDefaultLimit;
    const auto matches = dictionary_.SearchByPrefix(request.query(), limit);

    keywords::api::SearchResponse response;
    for (const auto& match : matches) {
        response.add_results(match);
    }

    return response;
}

}  // namespace search_core
