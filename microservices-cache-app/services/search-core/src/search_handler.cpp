#include "search_handler.hpp"

#include <exception>
#include <string>

#include <userver/components/component.hpp>
#include <userver/formats/common/type.hpp>
#include <userver/formats/json/serialize.hpp>
#include <userver/formats/json/value_builder.hpp>
#include <userver/server/handlers/http_handler_base.hpp>

#include "dictionary_component.hpp"

namespace search_core {

namespace {

constexpr std::size_t kDefaultLimit = 10;

std::size_t ParseLimit(const std::string& raw) {
    if (raw.empty()) {
        return kDefaultLimit;
    }

    try {
        return static_cast<std::size_t>(std::stoul(raw));
    } catch (const std::exception&) {
        return kDefaultLimit;
    }
}

class SearchHandler final : public userver::server::handlers::HttpHandlerBase {
public:
    static constexpr std::string_view kName = "handler-search";

    SearchHandler(const userver::components::ComponentConfig& config,
                  const userver::components::ComponentContext& context)
        : HttpHandlerBase(config, context), dictionary_(context.FindComponent<DictionaryComponent>()) {}

    std::string HandleRequestThrow(const userver::server::http::HttpRequest& request,
                                    userver::server::request::RequestContext&) const override {
        const auto prefix = request.GetArg("q");
        const auto limit = ParseLimit(request.GetArg("limit"));

        const auto matches = dictionary_.SearchByPrefix(prefix, limit);

        userver::formats::json::ValueBuilder builder(userver::formats::common::Type::kArray);
        for (const auto& match : matches) {
            builder.PushBack(match);
        }

        return userver::formats::json::ToString(builder.ExtractValue());
    }

private:
    const DictionaryComponent& dictionary_;
};

}  // namespace

void AppendSearchHandler(userver::components::ComponentList& component_list) {
    component_list.Append<SearchHandler>();
}

}  // namespace search_core
