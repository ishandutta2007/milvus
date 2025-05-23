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

#include <cstddef>
#include <cstdint>

#include "bitset/common.h"

namespace milvus {
namespace bitset {
namespace detail {

namespace dynamic {

///////////////////////////////////////////////////////////////////////////
// a facility to run through all acceptable data types
#define ALL_DATATYPES_1(FUNC) \
    FUNC(int8_t);             \
    FUNC(int16_t);            \
    FUNC(int32_t);            \
    FUNC(int64_t);            \
    FUNC(float);              \
    FUNC(double);

// a facility to run through all acceptable forward types
#define ALL_FORWARD_TYPES_1(FUNC) \
    FUNC(uint8_t);                \
    FUNC(uint64_t);

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename T, typename U, CompareOpType Op>
struct OpCompareColumnImpl {
    static bool
    op_compare_column(uint8_t* const __restrict bitmask,
                      const T* const __restrict t,
                      const U* const __restrict u,
                      const size_t size) {
        return false;
    }
};

#define DECLARE_PARTIAL_OP_COMPARE_COLUMN(TTYPE)             \
    template <CompareOpType Op>                              \
    struct OpCompareColumnImpl<TTYPE, TTYPE, Op> {           \
        static bool                                          \
        op_compare_column(uint8_t* const __restrict bitmask, \
                          const TTYPE* const __restrict t,   \
                          const TTYPE* const __restrict u,   \
                          const size_t size);                \
    };

ALL_DATATYPES_1(DECLARE_PARTIAL_OP_COMPARE_COLUMN)

#undef DECLARE_PARTIAL_OP_COMPARE_COLUMN

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename T, CompareOpType Op>
struct OpCompareValImpl {
    static inline bool
    op_compare_val(uint8_t* const __restrict bitmask,
                   const T* const __restrict t,
                   const size_t size,
                   const T& value) {
        return false;
    }
};

#define DECLARE_PARTIAL_OP_COMPARE_VAL(TTYPE)             \
    template <CompareOpType Op>                           \
    struct OpCompareValImpl<TTYPE, Op> {                  \
        static bool                                       \
        op_compare_val(uint8_t* const __restrict bitmask, \
                       const TTYPE* const __restrict t,   \
                       const size_t size,                 \
                       const TTYPE& value);               \
    };

ALL_DATATYPES_1(DECLARE_PARTIAL_OP_COMPARE_VAL)

#undef DECLARE_PARTIAL_OP_COMPARE_VAL

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename T, RangeType Op>
struct OpWithinRangeColumnImpl {
    static inline bool
    op_within_range_column(uint8_t* const __restrict bitmask,
                           const T* const __restrict lower,
                           const T* const __restrict upper,
                           const T* const __restrict values,
                           const size_t size) {
        return false;
    }
};

#define DECLARE_PARTIAL_OP_WITHIN_RANGE_COLUMN(TTYPE)                \
    template <RangeType Op>                                          \
    struct OpWithinRangeColumnImpl<TTYPE, Op> {                      \
        static bool                                                  \
        op_within_range_column(uint8_t* const __restrict bitmask,    \
                               const TTYPE* const __restrict lower,  \
                               const TTYPE* const __restrict upper,  \
                               const TTYPE* const __restrict values, \
                               const size_t size);                   \
    };

ALL_DATATYPES_1(DECLARE_PARTIAL_OP_WITHIN_RANGE_COLUMN)

#undef DECLARE_PARTIAL_OP_WITHIN_RANGE_COLUMN

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename T, RangeType Op>
struct OpWithinRangeValImpl {
    static inline bool
    op_within_range_val(uint8_t* const __restrict bitmask,
                        const T& lower,
                        const T& upper,
                        const T* const __restrict values,
                        const size_t size) {
        return false;
    }
};

#define DECLARE_PARTIAL_OP_WITHIN_RANGE_VAL(TTYPE)                \
    template <RangeType Op>                                       \
    struct OpWithinRangeValImpl<TTYPE, Op> {                      \
        static bool                                               \
        op_within_range_val(uint8_t* const __restrict bitmask,    \
                            const TTYPE& lower,                   \
                            const TTYPE& upper,                   \
                            const TTYPE* const __restrict values, \
                            const size_t size);                   \
    };

ALL_DATATYPES_1(DECLARE_PARTIAL_OP_WITHIN_RANGE_VAL)

#undef DECLARE_PARTIAL_OP_WITHIN_RANGE_VAL

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename T, ArithOpType AOp, CompareOpType CmpOp>
struct OpArithCompareImpl {
    static inline bool
    op_arith_compare(uint8_t* const __restrict bitmask,
                     const T* const __restrict src,
                     const ArithHighPrecisionType<T>& right_operand,
                     const ArithHighPrecisionType<T>& value,
                     const size_t size) {
        return false;
    }
};

#define DECLARE_PARTIAL_OP_ARITH_COMPARE(TTYPE)                              \
    template <ArithOpType AOp, CompareOpType CmpOp>                          \
    struct OpArithCompareImpl<TTYPE, AOp, CmpOp> {                           \
        static bool                                                          \
        op_arith_compare(uint8_t* const __restrict bitmask,                  \
                         const TTYPE* const __restrict src,                  \
                         const ArithHighPrecisionType<TTYPE>& right_operand, \
                         const ArithHighPrecisionType<TTYPE>& value,         \
                         const size_t size);                                 \
    };

ALL_DATATYPES_1(DECLARE_PARTIAL_OP_ARITH_COMPARE)

#undef DECLARE_PARTIAL_OP_ARITH_COMPARE

///////////////////////////////////////////////////////////////////////////
// the default implementation
template <typename ElementT>
struct ForwardOpsImpl {
    static inline bool
    op_and(ElementT* const left,
           const ElementT* const right,
           const size_t start_left,
           const size_t start_right,
           const size_t size) {
        return false;
    }

    static inline bool
    op_and_multiple(ElementT* const left,
                    const ElementT* const* const rights,
                    const size_t start_left,
                    const size_t* const __restrict start_rights,
                    const size_t n_rights,
                    const size_t size) {
        return false;
    }

    static inline bool
    op_or(ElementT* const left,
          const ElementT* const right,
          const size_t start_left,
          const size_t start_right,
          const size_t size) {
        return false;
    }

    static inline bool
    op_or_multiple(ElementT* const left,
                   const ElementT* const* const rights,
                   const size_t start_left,
                   const size_t* const __restrict start_rights,
                   const size_t n_rights,
                   const size_t size) {
        return false;
    }

    static inline bool
    op_xor(ElementT* const left,
           const ElementT* const right,
           const size_t start_left,
           const size_t start_right,
           const size_t size) {
        return false;
    }

    static inline bool
    op_sub(ElementT* const left,
           const ElementT* const right,
           const size_t start_left,
           const size_t start_right,
           const size_t size) {
        return false;
    }
};

#define DECLARE_PARTIAL_FORWARD_OPS(ELEMENTTYPE)                     \
    template <>                                                      \
    struct ForwardOpsImpl<ELEMENTTYPE> {                             \
        static bool                                                  \
        op_and(ELEMENTTYPE* const left,                              \
               const ELEMENTTYPE* const right,                       \
               const size_t start_left,                              \
               const size_t start_right,                             \
               const size_t size);                                   \
                                                                     \
        static bool                                                  \
        op_and_multiple(ELEMENTTYPE* const left,                     \
                        const ELEMENTTYPE* const* const rights,      \
                        const size_t start_left,                     \
                        const size_t* const __restrict start_rights, \
                        const size_t n_rights,                       \
                        const size_t size);                          \
                                                                     \
        static bool                                                  \
        op_or(ELEMENTTYPE* const left,                               \
              const ELEMENTTYPE* const right,                        \
              const size_t start_left,                               \
              const size_t start_right,                              \
              const size_t size);                                    \
                                                                     \
        static bool                                                  \
        op_or_multiple(ELEMENTTYPE* const left,                      \
                       const ELEMENTTYPE* const* const rights,       \
                       const size_t start_left,                      \
                       const size_t* const __restrict start_rights,  \
                       const size_t n_rights,                        \
                       const size_t size);                           \
                                                                     \
        static bool                                                  \
        op_sub(ELEMENTTYPE* const left,                              \
               const ELEMENTTYPE* const right,                       \
               const size_t start_left,                              \
               const size_t start_right,                             \
               const size_t size);                                   \
                                                                     \
        static bool                                                  \
        op_xor(ELEMENTTYPE* const left,                              \
               const ELEMENTTYPE* const right,                       \
               const size_t start_left,                              \
               const size_t start_right,                             \
               const size_t size);                                   \
    };

ALL_FORWARD_TYPES_1(DECLARE_PARTIAL_FORWARD_OPS)

#undef DECLARE_PARTIAL_FORWARD_OPS

///////////////////////////////////////////////////////////////////////////

#undef ALL_DATATYPES_1
#undef ALL_FORWARD_TYPES_1

}  // namespace dynamic

///////////////////////////////////////////////////////////////////////////

//
struct VectorizedDynamic {
    // Fills a bitmask by comparing two arrays element-wise.
    // API requirement: size % 8 == 0
    template <typename T, typename U, CompareOpType Op>
    static bool
    op_compare_column(uint8_t* const __restrict bitmask,
                      const T* const __restrict t,
                      const U* const __restrict u,
                      const size_t size) {
        return dynamic::OpCompareColumnImpl<T, U, Op>::op_compare_column(
            bitmask, t, u, size);
    }

    // Fills a bitmask by comparing elements of a given array to a
    //   given value.
    // API requirement: size % 8 == 0
    template <typename T, CompareOpType Op>
    static bool
    op_compare_val(uint8_t* const __restrict bitmask,
                   const T* const __restrict t,
                   const size_t size,
                   const T& value) {
        return dynamic::OpCompareValImpl<T, Op>::op_compare_val(
            bitmask, t, size, value);
    }

    // API requirement: size % 8 == 0
    template <typename T, RangeType Op>
    static bool
    op_within_range_column(uint8_t* const __restrict bitmask,
                           const T* const __restrict lower,
                           const T* const __restrict upper,
                           const T* const __restrict values,
                           const size_t size) {
        return dynamic::OpWithinRangeColumnImpl<T, Op>::op_within_range_column(
            bitmask, lower, upper, values, size);
    }

    // API requirement: size % 8 == 0
    template <typename T, RangeType Op>
    static bool
    op_within_range_val(uint8_t* const __restrict bitmask,
                        const T& lower,
                        const T& upper,
                        const T* const __restrict values,
                        const size_t size) {
        return dynamic::OpWithinRangeValImpl<T, Op>::op_within_range_val(
            bitmask, lower, upper, values, size);
    }

    // API requirement: size % 8 == 0
    template <typename T, ArithOpType AOp, CompareOpType CmpOp>
    static inline bool
    op_arith_compare(uint8_t* const __restrict bitmask,
                     const T* const __restrict src,
                     const ArithHighPrecisionType<T>& right_operand,
                     const ArithHighPrecisionType<T>& value,
                     const size_t size) {
        return dynamic::OpArithCompareImpl<T, AOp, CmpOp>::op_arith_compare(
            bitmask, src, right_operand, value, size);
    }

    // The following functions just forward parameters to the reference code,
    //   generated for a particular platform.

    template <typename ElementT>
    static inline bool
    forward_op_and(ElementT* const left,
                   const ElementT* const right,
                   const size_t start_left,
                   const size_t start_right,
                   const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_and(
            left, right, start_left, start_right, size);
    }

    template <typename ElementT>
    static inline bool
    forward_op_and_multiple(ElementT* const left,
                            const ElementT* const* const rights,
                            const size_t start_left,
                            const size_t* const __restrict start_rights,
                            const size_t n_rights,
                            const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_and_multiple(
            left, rights, start_left, start_rights, n_rights, size);
    }

    template <typename ElementT>
    static inline bool
    forward_op_or(ElementT* const left,
                  const ElementT* const right,
                  const size_t start_left,
                  const size_t start_right,
                  const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_or(
            left, right, start_left, start_right, size);
    }

    template <typename ElementT>
    static inline bool
    forward_op_or_multiple(ElementT* const left,
                           const ElementT* const* const rights,
                           const size_t start_left,
                           const size_t* const __restrict start_rights,
                           const size_t n_rights,
                           const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_or_multiple(
            left, rights, start_left, start_rights, n_rights, size);
    }

    template <typename ElementT>
    static inline bool
    forward_op_xor(ElementT* const left,
                   const ElementT* const right,
                   const size_t start_left,
                   const size_t start_right,
                   const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_xor(
            left, right, start_left, start_right, size);
    }

    template <typename ElementT>
    static inline bool
    forward_op_sub(ElementT* const left,
                   const ElementT* const right,
                   const size_t start_left,
                   const size_t start_right,
                   const size_t size) {
        return dynamic::ForwardOpsImpl<ElementT>::op_sub(
            left, right, start_left, start_right, size);
    }
};

}  // namespace detail
}  // namespace bitset
}  // namespace milvus