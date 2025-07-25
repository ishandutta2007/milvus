// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
#pragma once

#include "common/JsonCastType.h"
#include "common/Types.h"
#include "common/Consts.h"

namespace milvus::index {

struct NgramParams {
    bool loading_index;
    uintptr_t min_gram;
    uintptr_t max_gram;
};

struct CreateIndexInfo {
    DataType field_type;
    IndexType index_type;
    MetricType metric_type;
    IndexVersion index_engine_version;
    std::string field_name;
    int64_t dim;
    int32_t scalar_index_engine_version{1};
    uint32_t tantivy_index_version{7};
    JsonCastType json_cast_type{JsonCastType::UNKNOWN};
    std::string json_path;
    std::string json_cast_function{UNKNOW_CAST_FUNCTION_NAME};
    std::optional<NgramParams> ngram_params{std::nullopt};
};

}  // namespace milvus::index
