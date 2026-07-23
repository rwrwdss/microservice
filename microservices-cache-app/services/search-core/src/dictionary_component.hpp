#pragma once

#include <cstddef>
#include <string>
#include <vector>

#include <userver/components/component_base.hpp>
#include <userver/components/component_list.hpp>
#include <userver/yaml_config/schema.hpp>

namespace search_core {

class DictionaryComponent final : public userver::components::ComponentBase {
public:
    static constexpr std::string_view kName = "dictionary-component";

    DictionaryComponent(const userver::components::ComponentConfig& config,
                         const userver::components::ComponentContext& context);

    std::size_t Size() const;

    std::vector<std::string> SearchByPrefix(std::string_view prefix, std::size_t limit) const;

    static userver::yaml_config::Schema GetStaticConfigSchema();

private:
    std::vector<std::string> keywords_;
};

void AppendDictionaryComponent(userver::components::ComponentList& component_list);

}  // namespace search_core
