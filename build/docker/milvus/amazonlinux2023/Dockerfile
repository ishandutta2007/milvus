# Copyright (C) 2019-2020 Zilliz. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
# with the License. You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software distributed under the License
# is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
# or implied. See the License for the specific language governing permissions and limitations under the License.

FROM amazonlinux:2023

ARG TARGETARCH
ARG MILVUS_ASAN_LIB

RUN yum install -y wget libgomp libaio libatomic openblas-devel && \
    rm -rf /var/cache/yum/*

# Add Tini
RUN curl -L -o /tini https://github.com/krallin/tini/releases/download/v0.19.0/tini-$TARGETARCH && \
    chmod +x /tini

RUN mkdir -p /milvus/bin
COPY --chown=root:root --chmod=774 ./bin/milvus /milvus/bin/milvus

COPY --chown=root:root --chmod=774 ./configs/ /milvus/configs/

COPY --chown=root:root --chmod=774 ./lib/ /milvus/lib/


ENV PATH=/milvus/bin:$PATH
ENV LD_LIBRARY_PATH=/milvus/lib:$LD_LIBRARY_PATH:/usr/lib
ENV LD_PRELOAD=${MILVUS_ASAN_LIB}:/milvus/lib/libjemalloc.so
ENV MALLOC_CONF=background_thread:true

ENTRYPOINT ["/tini", "--"]

WORKDIR /milvus
