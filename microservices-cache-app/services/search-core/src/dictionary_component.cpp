#include "dictionary_component.hpp"

#include <fstream>
#include <stdexcept>

#include <userver/components/component.hpp>
#include <userver/logging/log.hpp>
#include <userver/yaml_config/merge_schemas.hpp>

namespace search_core {

namespace {

std::vector<std::string> LoadKeywordsFromFile(const std::string& path) {
    std::ifstream file(path);
    if (!file.is_open()) {
        throw std::runtime_error("failed to open dictionary file: " + path);
    }

    std::vector<std::string> keywords;
    std::string line;
    while (std::getline(file, line)) {
        if (!line.empty()) {
            keywords.push_back(line);
        }
    }

    return keywords;
}

}  // namespace

DictionaryComponent::DictionaryComponent(const userver::components::ComponentConfig& config,
                                          const userver::components::ComponentContext& context)
    : userver::components::ComponentBase(config, context) {
    const auto path = config["file-path"].As<std::string>();

    LOG_INFO() << "Loading dictionary from " << path;

    keywords_ = LoadKeywordsFromFile(path);

    LOG_INFO() << "Loaded " << keywords_.size() << " keywords";
}

std::size_t DictionaryComponent::Size() const { return keywords_.size(); }

std::vector<std::string> DictionaryComponent::SearchByPrefix(std::string_view prefix, std::size_t limit) const {
    std::vector<std::string> results;

    for (const auto& keyword : keywords_) {
        if (results.size() >= limit) {
            break;
        }

        if (keyword.rfind(prefix, 0) == 0) {
            results.push_back(keyword);
        }
    }

    return results;
}

userver::yaml_config::Schema DictionaryComponent::GetStaticConfigSchema() {
    return userver::yaml_config::MergeSchemas<userver::components::ComponentBase>(R"(
type: object
description: loads a text dictionary file into memory on startup
additionalProperties: false
properties:
    file-path:
        type: string
        description: path to the dictionary file
)");
}

void AppendDictionaryComponent(userver::components::ComponentList& component_list) {
    component_list.Append<DictionaryComponent>();
}

}  // namespace search_core
