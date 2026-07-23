#include "dictionary_handler.hpp"

#include <fmt/format.h>

#include <userver/components/component.hpp>
#include <userver/server/handlers/http_handler_base.hpp>

#include "dictionary_component.hpp"

namespace search_core {

namespace {

class DictionaryCountHandler final : public userver::server::handlers::HttpHandlerBase {
public:
    static constexpr std::string_view kName = "handler-dictionary-count";

    DictionaryCountHandler(const userver::components::ComponentConfig& config,
                            const userver::components::ComponentContext& context)
        : HttpHandlerBase(config, context),
          dictionary_(context.FindComponent<DictionaryComponent>()) {}

    std::string HandleRequestThrow(const userver::server::http::HttpRequest&,
                                    userver::server::request::RequestContext&) const override {
        return fmt::format(R"({{"count": {}}})", dictionary_.Size());
    }

private:
    const DictionaryComponent& dictionary_;
};

}  // namespace

void AppendDictionaryCountHandler(userver::components::ComponentList& component_list) {
    component_list.Append<DictionaryCountHandler>();
}

}  // namespace search_core
