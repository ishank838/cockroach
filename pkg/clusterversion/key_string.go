// Code generated by "stringer"; DO NOT EDIT.

package clusterversion

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[V21_2-0]
	_ = x[Start22_1-1]
	_ = x[TargetBytesAvoidExcess-2]
	_ = x[AvoidDrainingNames-3]
	_ = x[DrainingNamesMigration-4]
	_ = x[TraceIDDoesntImplyStructuredRecording-5]
	_ = x[AlterSystemTableStatisticsAddAvgSizeCol-6]
	_ = x[MVCCAddSSTable-7]
	_ = x[InsertPublicSchemaNamespaceEntryOnRestore-8]
	_ = x[UnsplitRangesInAsyncGCJobs-9]
	_ = x[ValidateGrantOption-10]
	_ = x[PebbleFormatBlockPropertyCollector-11]
	_ = x[ProbeRequest-12]
	_ = x[SelectRPCsTakeTracingInfoInband-13]
	_ = x[PreSeedTenantSpanConfigs-14]
	_ = x[SeedTenantSpanConfigs-15]
	_ = x[PublicSchemasWithDescriptors-16]
	_ = x[EnsureSpanConfigReconciliation-17]
	_ = x[EnsureSpanConfigSubscription-18]
	_ = x[EnableSpanConfigStore-19]
	_ = x[ScanWholeRows-20]
	_ = x[SCRAMAuthentication-21]
	_ = x[UnsafeLossOfQuorumRecoveryRangeLog-22]
	_ = x[AlterSystemProtectedTimestampAddColumn-23]
	_ = x[EnableProtectedTimestampsForTenant-24]
	_ = x[DeleteCommentsWithDroppedIndexes-25]
	_ = x[RemoveIncompatibleDatabasePrivileges-26]
	_ = x[AddRaftAppliedIndexTermMigration-27]
	_ = x[PostAddRaftAppliedIndexTermMigration-28]
	_ = x[DontProposeWriteTimestampForLeaseTransfers-29]
	_ = x[TenantSettingsTable-30]
	_ = x[EnablePebbleFormatVersionBlockProperties-31]
	_ = x[DisableSystemConfigGossipTrigger-32]
	_ = x[MVCCIndexBackfiller-33]
	_ = x[EnableLeaseHolderRemoval-34]
	_ = x[BackupResolutionInJob-35]
	_ = x[LooselyCoupledRaftLogTruncation-36]
	_ = x[ChangefeedIdleness-37]
	_ = x[BackupDoesNotOverwriteLatestAndCheckpoint-38]
	_ = x[EnableDeclarativeSchemaChanger-39]
	_ = x[RowLevelTTL-40]
	_ = x[PebbleFormatSplitUserKeysMarked-41]
	_ = x[IncrementalBackupSubdir-42]
	_ = x[DateStyleIntervalStyleCastRewrite-43]
	_ = x[EnableNewStoreRebalancer-44]
	_ = x[ClusterLocksVirtualTable-45]
	_ = x[AutoStatsTableSettings-46]
	_ = x[ForecastStats-47]
	_ = x[SuperRegions-48]
	_ = x[EnableNewChangefeedOptions-49]
	_ = x[SpanCountTable-50]
	_ = x[PreSeedSpanCountTable-51]
	_ = x[SeedSpanCountTable-52]
	_ = x[V22_1-53]
	_ = x[Start22_2-54]
	_ = x[LocalTimestamps-55]
	_ = x[EnsurePebbleFormatVersionRangeKeys-56]
	_ = x[EnablePebbleFormatVersionRangeKeys-57]
	_ = x[TrigramInvertedIndexes-58]
	_ = x[RemoveGrantPrivilege-59]
	_ = x[MVCCRangeTombstones-60]
	_ = x[UpgradeSequenceToBeReferencedByID-61]
	_ = x[SampledStmtDiagReqs-62]
}

const _Key_name = "V21_2Start22_1TargetBytesAvoidExcessAvoidDrainingNamesDrainingNamesMigrationTraceIDDoesntImplyStructuredRecordingAlterSystemTableStatisticsAddAvgSizeColMVCCAddSSTableInsertPublicSchemaNamespaceEntryOnRestoreUnsplitRangesInAsyncGCJobsValidateGrantOptionPebbleFormatBlockPropertyCollectorProbeRequestSelectRPCsTakeTracingInfoInbandPreSeedTenantSpanConfigsSeedTenantSpanConfigsPublicSchemasWithDescriptorsEnsureSpanConfigReconciliationEnsureSpanConfigSubscriptionEnableSpanConfigStoreScanWholeRowsSCRAMAuthenticationUnsafeLossOfQuorumRecoveryRangeLogAlterSystemProtectedTimestampAddColumnEnableProtectedTimestampsForTenantDeleteCommentsWithDroppedIndexesRemoveIncompatibleDatabasePrivilegesAddRaftAppliedIndexTermMigrationPostAddRaftAppliedIndexTermMigrationDontProposeWriteTimestampForLeaseTransfersTenantSettingsTableEnablePebbleFormatVersionBlockPropertiesDisableSystemConfigGossipTriggerMVCCIndexBackfillerEnableLeaseHolderRemovalBackupResolutionInJobLooselyCoupledRaftLogTruncationChangefeedIdlenessBackupDoesNotOverwriteLatestAndCheckpointEnableDeclarativeSchemaChangerRowLevelTTLPebbleFormatSplitUserKeysMarkedIncrementalBackupSubdirDateStyleIntervalStyleCastRewriteEnableNewStoreRebalancerClusterLocksVirtualTableAutoStatsTableSettingsForecastStatsSuperRegionsEnableNewChangefeedOptionsSpanCountTablePreSeedSpanCountTableSeedSpanCountTableV22_1Start22_2LocalTimestampsEnsurePebbleFormatVersionRangeKeysEnablePebbleFormatVersionRangeKeysTrigramInvertedIndexesRemoveGrantPrivilegeMVCCRangeTombstonesUpgradeSequenceToBeReferencedByIDSampledStmtDiagReqs"

var _Key_index = [...]uint16{0, 5, 14, 36, 54, 76, 113, 152, 166, 207, 233, 252, 286, 298, 329, 353, 374, 402, 432, 460, 481, 494, 513, 547, 585, 619, 651, 687, 719, 755, 797, 816, 856, 888, 907, 931, 952, 983, 1001, 1042, 1072, 1083, 1114, 1137, 1170, 1194, 1218, 1240, 1253, 1265, 1291, 1305, 1326, 1344, 1349, 1358, 1373, 1407, 1441, 1463, 1483, 1502, 1535, 1554}

func (i Key) String() string {
	if i < 0 || i >= Key(len(_Key_index)-1) {
		return "Key(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
