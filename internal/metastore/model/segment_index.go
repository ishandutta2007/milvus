package model

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/proto/indexpb"
)

type SegmentIndex struct {
	SegmentID      int64
	CollectionID   int64
	PartitionID    int64
	NumRows        int64
	IndexID        int64
	BuildID        int64
	NodeID         int64
	IndexVersion   int64
	IndexState     commonpb.IndexState
	FailReason     string
	IsDeleted      bool
	CreatedUTCTime uint64
	IndexFileKeys  []string
	// The byte size of serialized index data at oos. (compressed)
	IndexSerializedSize uint64
	// The byte size of index data in memory. (uncompressed)
	// The IndexMemSize may not be stored at old milvus implementation, so it may be not accurate.
	// (generated by the IndexSerializedSize multiplied with a configured compress-ratio).
	IndexMemSize uint64
	// deprecated
	WriteHandoff              bool
	CurrentIndexVersion       int32
	IndexStoreVersion         int64
	FinishedUTCTime           uint64
	CurrentScalarIndexVersion int32
}

func UnmarshalSegmentIndexModel(segIndex *indexpb.SegmentIndex) *SegmentIndex {
	if segIndex == nil {
		return nil
	}
	return &SegmentIndex{
		SegmentID:                 segIndex.SegmentID,
		CollectionID:              segIndex.CollectionID,
		PartitionID:               segIndex.PartitionID,
		NumRows:                   segIndex.NumRows,
		IndexID:                   segIndex.IndexID,
		BuildID:                   segIndex.BuildID,
		NodeID:                    segIndex.NodeID,
		IndexState:                segIndex.State,
		FailReason:                segIndex.FailReason,
		IndexVersion:              segIndex.IndexVersion,
		IsDeleted:                 segIndex.Deleted,
		CreatedUTCTime:            segIndex.CreateTime,
		IndexFileKeys:             common.CloneStringList(segIndex.IndexFileKeys),
		IndexSerializedSize:       segIndex.SerializeSize,
		IndexMemSize:              segIndex.MemSize,
		WriteHandoff:              segIndex.WriteHandoff,
		CurrentIndexVersion:       segIndex.GetCurrentIndexVersion(),
		FinishedUTCTime:           segIndex.FinishedTime,
		CurrentScalarIndexVersion: segIndex.CurrentScalarIndexVersion,
	}
}

func MarshalSegmentIndexModel(segIdx *SegmentIndex) *indexpb.SegmentIndex {
	if segIdx == nil {
		return nil
	}

	return &indexpb.SegmentIndex{
		CollectionID:              segIdx.CollectionID,
		PartitionID:               segIdx.PartitionID,
		SegmentID:                 segIdx.SegmentID,
		NumRows:                   segIdx.NumRows,
		IndexID:                   segIdx.IndexID,
		BuildID:                   segIdx.BuildID,
		NodeID:                    segIdx.NodeID,
		State:                     segIdx.IndexState,
		FailReason:                segIdx.FailReason,
		IndexVersion:              segIdx.IndexVersion,
		IndexFileKeys:             common.CloneStringList(segIdx.IndexFileKeys),
		Deleted:                   segIdx.IsDeleted,
		CreateTime:                segIdx.CreatedUTCTime,
		SerializeSize:             segIdx.IndexSerializedSize,
		MemSize:                   segIdx.IndexMemSize,
		WriteHandoff:              segIdx.WriteHandoff,
		CurrentIndexVersion:       segIdx.CurrentIndexVersion,
		FinishedTime:              segIdx.FinishedUTCTime,
		CurrentScalarIndexVersion: segIdx.CurrentScalarIndexVersion,
	}
}

func CloneSegmentIndex(segIndex *SegmentIndex) *SegmentIndex {
	return &SegmentIndex{
		SegmentID:                 segIndex.SegmentID,
		CollectionID:              segIndex.CollectionID,
		PartitionID:               segIndex.PartitionID,
		NumRows:                   segIndex.NumRows,
		IndexID:                   segIndex.IndexID,
		BuildID:                   segIndex.BuildID,
		NodeID:                    segIndex.NodeID,
		IndexState:                segIndex.IndexState,
		FailReason:                segIndex.FailReason,
		IndexVersion:              segIndex.IndexVersion,
		IsDeleted:                 segIndex.IsDeleted,
		CreatedUTCTime:            segIndex.CreatedUTCTime,
		IndexFileKeys:             common.CloneStringList(segIndex.IndexFileKeys),
		IndexSerializedSize:       segIndex.IndexSerializedSize,
		IndexMemSize:              segIndex.IndexMemSize,
		WriteHandoff:              segIndex.WriteHandoff,
		CurrentIndexVersion:       segIndex.CurrentIndexVersion,
		FinishedUTCTime:           segIndex.FinishedUTCTime,
		CurrentScalarIndexVersion: segIndex.CurrentScalarIndexVersion,
	}
}
