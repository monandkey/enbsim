package s1ap

const (
	CAMPER_TYPE_TEMPORARY = 0
	CAMPER_TYPE_NORMAL    = 1
)

const (
	reject = iota
	ignore
	notify
)

const (
	idHandoverPreparation                  = 0
	idHandoverResourceAllocation           = 1
	idHandoverNotification                 = 2
	idPathSwitchRequest                    = 3
	idHandoverCancel                       = 4
	idERABSetup                            = 5
	idERABModify                           = 6
	idERABRelease                          = 7
	idERABReleaseIndication                = 8
	idInitialContextSetup                  = 9
	idPaging                               = 10
	iddownlinkNASTransport                 = 11
	idinitialUEMessage                     = 12
	iduplinkNASTransport                   = 13
	idReset                                = 14
	idErrorIndication                      = 15
	idNASNonDeliveryIndication             = 16
	idS1Setup                              = 17
	idUEContextReleaseRequest              = 18
	idDownlinkS1cdma2000tunnelling         = 19
	idUplinkS1cdma2000tunnelling           = 20
	idUEContextModification                = 21
	idUECapabilityInfoIndication           = 22
	idUEContextRelease                     = 23
	ideNBStatusTransfer                    = 24
	idMMEStatusTransfer                    = 25
	idDeactivateTrace                      = 26
	idTraceStart                           = 27
	idTraceFailureIndication               = 28
	idENBConfigurationUpdate               = 29
	idMMEConfigurationUpdate               = 30
	idLocationReportingControl             = 31
	idLocationReportingFailureIndication   = 32
	idLocationReport                       = 33
	idOverloadStart                        = 34
	idOverloadStop                         = 35
	idWriteReplaceWarning                  = 36
	ideNBDirectInformationTransfer         = 37
	idMMEDirectInformationTransfer         = 38
	idPrivateMessage                       = 39
	ideNBConfigurationTransfer             = 40
	idMMEConfigurationTransfer             = 41
	idCellTrafficTrace                     = 42
	idKill                                 = 43
	iddownlinkUEAssociatedLPPaTransport    = 44
	iduplinkUEAssociatedLPPaTransport      = 45
	iddownlinkNonUEAssociatedLPPaTransport = 46
	iduplinkNonUEAssociatedLPPaTransport   = 47
	idUERadioCapabilityMatch               = 48
	idPWSRestartIndication                 = 49
	idERABModificationIndication           = 50
	idPWSFailureIndication                 = 51
	idRerouteNASRequest                    = 52
	idUEContextModificationIndication      = 53
	idConnectionEstablishmentIndication    = 54
	idUEContextSuspend                     = 55
	idUEContextResume                      = 56
	idNASDeliveryIndication                = 57
	idRetrieveUEInformation                = 58
	idUEInformationTransfer                = 59
	ideNBCPRelocationIndication            = 60
	idMMECPRelocationIndication            = 61
	idSecondaryRATDataUsageReport          = 62
	idUERadioCapabilityIDMapping           = 63
	idHandoverSuccess                      = 64
	ideNBEarlyStatusTransfer               = 65
	idMMEEarlyStatusTransfer               = 66
)

var procedureIDs = map[int]string{
	idHandoverPreparation:                  "id-HandoverPreparation",
	idHandoverResourceAllocation:           "id-HandoverResourceAllocation",
	idHandoverNotification:                 "id-HandoverNotification",
	idPathSwitchRequest:                    "id-PathSwitchRequest",
	idHandoverCancel:                       "id-HandoverCancel",
	idERABSetup:                            "id-E-RABSetup",
	idERABModify:                           "id-E-RABModify",
	idERABRelease:                          "id-E-RABRelease",
	idERABReleaseIndication:                "id-E-RABReleaseIndication",
	idInitialContextSetup:                  "id-InitialContextSetup",
	idPaging:                               "id-Paging",
	iddownlinkNASTransport:                 "id-downlinkNASTransport",
	idinitialUEMessage:                     "id-initialUEMessage",
	iduplinkNASTransport:                   "id-uplinkNASTransport",
	idReset:                                "id-Reset",
	idErrorIndication:                      "id-ErrorIndication",
	idNASNonDeliveryIndication:             "id-NASNonDeliveryIndication",
	idS1Setup:                              "id-S1Setup",
	idUEContextReleaseRequest:              "id-UEContextReleaseRequest",
	idDownlinkS1cdma2000tunnelling:         "id-DownlinkS1cdma2000tunnelling",
	idUplinkS1cdma2000tunnelling:           "id-UplinkS1cdma2000tunnelling",
	idUEContextModification:                "id-UEContextModification",
	idUECapabilityInfoIndication:           "id-UECapabilityInfoIndication",
	idUEContextRelease:                     "id-UEContextRelease",
	ideNBStatusTransfer:                    "id-eNBStatusTransfer",
	idMMEStatusTransfer:                    "id-MMEStatusTransfer",
	idDeactivateTrace:                      "id-DeactivateTrace",
	idTraceStart:                           "id-TraceStart",
	idTraceFailureIndication:               "id-TraceFailureIndication",
	idENBConfigurationUpdate:               "id-ENBConfigurationUpdate",
	idMMEConfigurationUpdate:               "id-MMEConfigurationUpdate",
	idLocationReportingControl:             "id-LocationReportingControl",
	idLocationReportingFailureIndication:   "id-LocationReportingFailureIndication",
	idLocationReport:                       "id-LocationReport",
	idOverloadStart:                        "id-OverloadStart",
	idOverloadStop:                         "id-OverloadStop",
	idWriteReplaceWarning:                  "id-WriteReplaceWarning",
	ideNBDirectInformationTransfer:         "id-eNBDirectInformationTransfer",
	idMMEDirectInformationTransfer:         "id-MMEDirectInformationTransfer",
	idPrivateMessage:                       "id-PrivateMessage",
	ideNBConfigurationTransfer:             "id-eNBConfigurationTransfer",
	idMMEConfigurationTransfer:             "id-MMEConfigurationTransfer",
	idCellTrafficTrace:                     "id-CellTrafficTrace",
	idKill:                                 "id-Kill",
	iddownlinkUEAssociatedLPPaTransport:    "id-downlinkUEAssociatedLPPaTransport",
	iduplinkUEAssociatedLPPaTransport:      "id-uplinkUEAssociatedLPPaTransport",
	iddownlinkNonUEAssociatedLPPaTransport: "id-downlinkNonUEAssociatedLPPaTransport",
	iduplinkNonUEAssociatedLPPaTransport:   "id-uplinkNonUEAssociatedLPPaTransport",
	idUERadioCapabilityMatch:               "id-UERadioCapabilityMatch",
	idPWSRestartIndication:                 "id-PWSRestartIndication",
	idERABModificationIndication:           "id-E-RABModificationIndication",
	idPWSFailureIndication:                 "id-PWSFailureIndication",
	idRerouteNASRequest:                    "id-RerouteNASRequest",
	idUEContextModificationIndication:      "id-UEContextModificationIndication",
	idConnectionEstablishmentIndication:    "id-ConnectionEstablishmentIndication",
	idUEContextSuspend:                     "id-UEContextSuspend",
	idUEContextResume:                      "id-UEContextResume",
	idNASDeliveryIndication:                "id-NASDeliveryIndication",
	idRetrieveUEInformation:                "id-RetrieveUEInformation",
	idUEInformationTransfer:                "id-UEInformationTransfer",
	ideNBCPRelocationIndication:            "id-eNBCPRelocationIndication",
	idMMECPRelocationIndication:            "id-MMECPRelocationIndication",
	idSecondaryRATDataUsageReport:          "id-SecondaryRATDataUsageReport",
	idUERadioCapabilityIDMapping:           "id-UERadioCapabilityIDMapping",
	idHandoverSuccess:                      "id-HandoverSuccess",
	ideNBEarlyStatusTransfer:               "id-eNBEarlyStatusTransfer",
	idMMEEarlyStatusTransfer:               "id-MMEEarlyStatusTransfer",
}

const (
	idHandoverType                                  = 1
	idCause                                         = 2
	idSourceID                                      = 3
	idTargetID                                      = 4
	ideNBUES1APID                                   = 8
	idERABSubjecttoDataForwardingList               = 12
	idERABtoReleaseListHOCmd                        = 13
	idERABDataForwardingItem                        = 14
	idERABReleaseItemBearerRelComp                  = 15
	idERABToBeSetupListBearerSUReq                  = 16
	idERABToBeSetupItemBearerSUReq                  = 17
	idERABAdmittedList                              = 18
	idERABFailedToSetupListHOReqAck                 = 19
	idERABAdmittedItem                              = 20
	idERABFailedtoSetupItemHOReqAck                 = 21
	idERABToBeSwitchedDLList                        = 22
	idERABToBeSwitchedDLItem                        = 23
	idERABToBeSetupListCtxtSUReq                    = 24
	idTraceActivation                               = 25
	idNASPDU                                        = 26
	idERABToBeSetupItemHOReq                        = 27
	idERABSetupListBearerSURes                      = 28
	idERABFailedToSetupListBearerSURes              = 29
	idERABToBeModifiedListBearerModReq              = 30
	idERABModifyListBearerModRes                    = 31
	idERABFailedToModifyList                        = 32
	idERABToBeReleasedList                          = 33
	idERABFailedToReleaseList                       = 34
	idERABItem                                      = 35
	idERABToBeModifiedItemBearerModReq              = 36
	idERABModifyItemBearerModRes                    = 37
	idERABReleaseItem                               = 38
	idERABSetupItemBearerSURes                      = 39
	idSecurityContext                               = 40
	idHandoverRestrictionList                       = 41
	idUEPagingID                                    = 43
	idpagingDRX                                     = 44
	idTAIList                                       = 46
	idTAIItem                                       = 47
	idERABFailedToSetupListCtxtSURes                = 48
	idERABReleaseItemHOCmd                          = 49
	idERABSetupItemCtxtSURes                        = 50
	idERABSetupListCtxtSURes                        = 51
	idERABToBeSetupItemCtxtSUReq                    = 52
	idERABToBeSetupListHOReq                        = 53
	idGERANtoLTEHOInformationRes                    = 55
	idUTRANtoLTEHOInformationRes                    = 57
	idCriticalityDiagnostics                        = 58
	idGlobalENBID                                   = 59
	ideNBname                                       = 60
	idMMEname                                       = 61
	idServedPLMNs                                   = 63
	idSupportedTAs                                  = 64
	idTimeToWait                                    = 65
	iduEaggregateMaximumBitrate                     = 66
	idTAI                                           = 67
	idERABReleaseListBearerRelComp                  = 69
	idcdma2000PDU                                   = 70
	idcdma2000RATType                               = 71
	idcdma2000SectorID                              = 72
	idSecurityKey                                   = 73
	idUERadioCapability                             = 74
	idGUMMEIID                                      = 75
	idERABInformationListItem                       = 78
	idDirectForwardingPathAvailability              = 79
	idUEIdentityIndexValue                          = 80
	idcdma2000HOStatus                              = 83
	idcdma2000HORequiredIndication                  = 84
	idEUTRANTraceID                                 = 86
	idRelativeMMECapacity                           = 87
	idSourceMMEUES1APID                             = 88
	idBearersSubjectToStatusTransferItem            = 89
	ideNBStatusTransferTransparentContainer         = 90
	idUEassociatedLogicalS1ConnectionItem           = 91
	idResetType                                     = 92
	idUEassociatedLogicalS1ConnectionListResAck     = 93
	idERABToBeSwitchedULItem                        = 94
	idERABToBeSwitchedULList                        = 95
	idSTMSI                                         = 96
	idcdma2000OneXRAND                              = 97
	idRequestType                                   = 98
	idUES1APIDs                                     = 99
	idEUTRANCGI                                     = 100
	idOverloadResponse                              = 101
	idcdma2000OneXSRVCCInfo                         = 102
	idERABFailedToBeReleasedList                    = 103
	idSourceToTargetTransparentContainer            = 104
	idServedGUMMEIs                                 = 105
	idSubscriberProfileIDforRFP                     = 106
	idUESecurityCapabilities                        = 107
	idCSFallbackIndicator                           = 108
	idCNDomain                                      = 109
	idERABReleasedList                              = 110
	idMessageIdentifier                             = 111
	idSerialNumber                                  = 112
	idWarningAreaList                               = 113
	idRepetitionPeriod                              = 114
	idNumberofBroadcastRequest                      = 115
	idWarningType                                   = 116
	idWarningSecurityInfo                           = 117
	idDataCodingScheme                              = 118
	idWarningMessageContents                        = 119
	idBroadcastCompletedAreaList                    = 120
	idInterSystemInformationTransferTypeEDT         = 121
	idInterSystemInformationTransferTypeMDT         = 122
	idTargetToSourceTransparentContainer            = 123
	idSRVCCOperationPossible                        = 124
	idSRVCCHOIndication                             = 125
	idNASDownlinkCount                              = 126
	idCSGId                                         = 127
	idCSGIdList                                     = 128
	idSONConfigurationTransferECT                   = 129
	idSONConfigurationTransferMCT                   = 130
	idTraceCollectionEntityIPAddress                = 131
	idMSClassmark2                                  = 132
	idMSClassmark3                                  = 133
	idRRCEstablishmentCause                         = 134
	idNASSecurityParametersfromEUTRAN               = 135
	idNASSecurityParameterstoEUTRAN                 = 136
	idDefaultPagingDRX                              = 137
	idSourceToTargetTransparentContainerSecondary   = 138
	idTargetToSourceTransparentContainerSecondary   = 139
	idEUTRANRoundTripDelayEstimationInfo            = 140
	idBroadcastCancelledAreaList                    = 141
	idConcurrentWarningMessageIndicator             = 142
	idDataForwardingNotPossible                     = 143
	idExtendedRepetitionPeriod                      = 144
	idCellAccessMode                                = 145
	idCSGMembershipStatus                           = 146
	idLPPaPDU                                       = 147
	idRoutingID                                     = 148
	idTimeSynchronisationInfo                       = 149
	idPSServiceNotAvailable                         = 150
	idPagingPriority                                = 151
	idx2TNLConfigurationInfo                        = 152
	ideNBX2ExtendedTransportLayerAddresses          = 153
	idGUMMEIList                                    = 154
	idGWTransportLayerAddress                       = 155
	idCorrelationID                                 = 156
	idSourceMMEGUMMEI                               = 157
	idMMEUES1APID2                                  = 158
	idRegisteredLAI                                 = 159
	idRelayNodeIndicator                            = 160
	idTrafficLoadReductionIndication                = 161
	idMDTConfiguration                              = 162
	idMMERelaySupportIndicator                      = 163
	idGWContextReleaseIndication                    = 164
	idManagementBasedMDTAllowed                     = 165
	idPrivacyIndicator                              = 166
	idTimeUEStayedInCellEnhancedGranularity         = 167
	idHOCause                                       = 168
	idVoiceSupportMatchIndicator                    = 169
	idGUMMEIType                                    = 170
	idM3Configuration                               = 171
	idM4Configuration                               = 172
	idM5Configuration                               = 173
	idMDTLocationInfo                               = 174
	idMobilityInformation                           = 175
	idTunnelInformationforBBF                       = 176
	idManagementBasedMDTPLMNList                    = 177
	idSignallingBasedMDTPLMNList                    = 178
	idULCOUNTValueExtended                          = 179
	idDLCOUNTValueExtended                          = 180
	idReceiveStatusOfULPDCPSDUsExtended             = 181
	idECGIListForRestart                            = 182
	idSIPTOCorrelationID                            = 183
	idSIPTOLGWTransportLayerAddress                 = 184
	idTransportInformation                          = 185
	idLHNID                                         = 186
	idAdditionalCSFallbackIndicator                 = 187
	idTAIListForRestart                             = 188
	idUserLocationInformation                       = 189
	idEmergencyAreaIDListForRestart                 = 190
	idKillAllWarningMessages                        = 191
	idMaskedIMEISV                                  = 192
	ideNBIndirectX2TransportLayerAddresses          = 193
	iduEHistoryInformationFromTheUE                 = 194
	idProSeAuthorized                               = 195
	idExpectedUEBehaviour                           = 196
	idLoggedMBSFNMDT                                = 197
	idUERadioCapabilityForPaging                    = 198
	idERABToBeModifiedListBearerModInd              = 199
	idERABToBeModifiedItemBearerModInd              = 200
	idERABNotToBeModifiedListBearerModInd           = 201
	idERABNotToBeModifiedItemBearerModInd           = 202
	idERABModifyListBearerModConf                   = 203
	idERABModifyItemBearerModConf                   = 204
	idERABFailedToModifyListBearerModConf           = 205
	idSONInformationReport                          = 206
	idMutingAvailabilityIndication                  = 207
	idMutingPatternInformation                      = 208
	idSynchronisationInformation                    = 209
	idERABToBeReleasedListBearerModConf             = 210
	idAssistanceDataForPaging                       = 211
	idCellIdentifierAndCELevelForCECapableUEs       = 212
	idInformationOnRecommendedCellsAndENBsForPaging = 213
	idRecommendedCellItem                           = 214
	idRecommendedENBItem                            = 215
	idProSeUEtoNetworkRelaying                      = 216
	idULCOUNTValuePDCPSNlength18                    = 217
	idDLCOUNTValuePDCPSNlength18                    = 218
	idReceiveStatusOfULPDCPSDUsPDCPSNlength18       = 219
	idM6Configuration                               = 220
	idM7Configuration                               = 221
	idPWSfailedECGIList                             = 222
	idMMEGroupID                                    = 223
	idAdditionalGUTI                                = 224
	idS1Message                                     = 225
	idCSGMembershipInfo                             = 226
	idPagingeDRXInformation                         = 227
	idUERetentionInformation                        = 228
	idUEUsageType                                   = 230
	idextendedUEIdentityIndexValue                  = 231
	idRATType                                       = 232
	idBearerType                                    = 233
	idNBIoTDefaultPagingDRX                         = 234
	idERABFailedToResumeListResumeReq               = 235
	idERABFailedToResumeItemResumeReq               = 236
	idERABFailedToResumeListResumeRes               = 237
	idERABFailedToResumeItemResumeRes               = 238
	idNBIoTPagingeDRXInformation                    = 239
	idV2XServicesAuthorized                         = 240
	idUEUserPlaneCIoTSupportIndicator               = 241
	idCEmodeBSupportIndicator                       = 242
	idSRVCCOperationNotPossible                     = 243
	idNBIoTUEIdentityIndexValue                     = 244
	idRRCResumeCause                                = 245
	idDCNID                                         = 246
	idServedDCNs                                    = 247
	idUESidelinkAggregateMaximumBitrate             = 248
	idDLNASPDUDeliveryAckRequest                    = 249
	idCoverageLevel                                 = 250
	idEnhancedCoverageRestricted                    = 251
	idUELevelQoSParameters                          = 252
	idDLCPSecurityInformation                       = 253
	idULCPSecurityInformation                       = 254
	idextendedeRABMaximumBitrateDL                  = 255
	idextendedeRABMaximumBitrateUL                  = 256
	idextendedeRABGuaranteedBitrateDL               = 257
	idextendedeRABGuaranteedBitrateUL               = 258
	idextendeduEaggregateMaximumBitRateDL           = 259
	idextendeduEaggregateMaximumBitRateUL           = 260
	idNRrestrictioninEPSasSecondaryRAT              = 261
	idUEAppLayerMeasConfig                          = 262
	idUEApplicationLayerMeasurementCapability       = 263
	idSecondaryRATDataUsageReportList               = 264
	idSecondaryRATDataUsageReportItem               = 265
	idHandoverFlag                                  = 266
	idERABUsageReportItem                           = 267
	idSecondaryRATDataUsageRequest                  = 268
	idNRUESecurityCapabilities                      = 269
	idUnlicensedSpectrumRestriction                 = 270
	idCEModeBRestricted                             = 271
	idLTEMIndication                                = 272
	idDownlinkPacketLossRate                        = 273
	idUplinkPacketLossRate                          = 274
	idUECapabilityInfoRequest                       = 275
	idserviceType                                   = 276
	idAerialUEsubscriptionInformation               = 277
	idSubscriptionBasedUEDifferentiationInfo        = 278
	idEndIndication                                 = 280
	idEDTSession                                    = 281
	idCNTypeRestrictions                            = 282
	idPendingDataIndication                         = 283
	idBluetoothMeasurementConfiguration             = 284
	idWLANMeasurementConfiguration                  = 285
	idWarningAreaCoordinates                        = 286
	idNRrestrictionin5GS                            = 287
	idPSCellInformation                             = 288
	idLastNGRANPLMNIdentity                         = 290
	idConnectedengNBList                            = 291
	idConnectedengNBToAddList                       = 292
	idConnectedengNBToRemoveList                    = 293
	idENDCSONConfigurationTransferECT               = 294
	idENDCSONConfigurationTransferMCT               = 295
	idIMSvoiceEPSfallbackfrom5G                     = 296
	idTimeSinceSecondaryNodeRelease                 = 297
	idRequestTypeAdditionalInfo                     = 298
	idAdditionalRRMPriorityIndex                    = 299
	idContextatSource                               = 300
	idIABAuthorized                                 = 301
	idIABNodeIndication                             = 302
	idIABSupported                                  = 303
	idDataSize                                      = 304
	idEthernetType                                  = 305
	idNRV2XServicesAuthorized                       = 306
	idNRUESidelinkAggregateMaximumBitrate           = 307
	idPC5QoSParameters                              = 308
	idIntersystemSONConfigurationTransferMCT        = 309
	idIntersystemSONConfigurationTransferECT        = 310
	idIntersystemMeasurementConfiguration           = 311
	idSourceNodeID                                  = 312
	idNBIoTRLFReportContainer                       = 313
	idUERadioCapabilityID                           = 314
	idUERadioCapabilityNRFormat                     = 315
	idMDTConfigurationNR                            = 316
	idDAPSRequestInfo                               = 317
	idDAPSResponseInfoList                          = 318
	idDAPSResponseInfoItem                          = 319
	idNotifySourceeNB                               = 320
	ideNBEarlyStatusTransferTransparentContainer    = 321
	idBearersSubjectToEarlyStatusTransferItem       = 322
	idWUSAssistanceInformation                      = 323
	idNBIoTPagingDRX                                = 324
	idTraceCollectionEntityURI                      = 325
	idEmergencyIndicator                            = 326
)

var ieIDs = map[int]string{
	idHandoverType:                                  "id-HandoverType",
	idCause:                                         "id-Cause",
	idSourceID:                                      "id-SourceID",
	idTargetID:                                      "id-TargetID",
	ideNBUES1APID:                                   "id-eNB-UE-S1AP-ID",
	idERABSubjecttoDataForwardingList:               "id-E-RABSubjecttoDataForwardingList",
	idERABtoReleaseListHOCmd:                        "id-E-RABtoReleaseListHOCmd",
	idERABDataForwardingItem:                        "id-E-RABDataForwardingItem",
	idERABReleaseItemBearerRelComp:                  "id-E-RABReleaseItemBearerRelComp",
	idERABToBeSetupListBearerSUReq:                  "id-E-RABToBeSetupListBearerSUReq",
	idERABToBeSetupItemBearerSUReq:                  "id-E-RABToBeSetupItemBearerSUReq",
	idERABAdmittedList:                              "id-E-RABAdmittedList",
	idERABFailedToSetupListHOReqAck:                 "id-E-RABFailedToSetupListHOReqAck",
	idERABAdmittedItem:                              "id-E-RABAdmittedItem",
	idERABFailedtoSetupItemHOReqAck:                 "id-E-RABFailedtoSetupItemHOReqAck",
	idERABToBeSwitchedDLList:                        "id-E-RABToBeSwitchedDLList",
	idERABToBeSwitchedDLItem:                        "id-E-RABToBeSwitchedDLItem",
	idERABToBeSetupListCtxtSUReq:                    "id-E-RABToBeSetupListCtxtSUReq",
	idTraceActivation:                               "id-TraceActivation",
	idNASPDU:                                        "id-NAS-PDU",
	idERABToBeSetupItemHOReq:                        "id-E-RABToBeSetupItemHOReq",
	idERABSetupListBearerSURes:                      "id-E-RABSetupListBearerSURes",
	idERABFailedToSetupListBearerSURes:              "id-E-RABFailedToSetupListBearerSURes",
	idERABToBeModifiedListBearerModReq:              "id-E-RABToBeModifiedListBearerModReq",
	idERABModifyListBearerModRes:                    "id-E-RABModifyListBearerModRes",
	idERABFailedToModifyList:                        "id-E-RABFailedToModifyList",
	idERABToBeReleasedList:                          "id-E-RABToBeReleasedList",
	idERABFailedToReleaseList:                       "id-E-RABFailedToReleaseList",
	idERABItem:                                      "id-E-RABItem",
	idERABToBeModifiedItemBearerModReq:              "id-E-RABToBeModifiedItemBearerModReq",
	idERABModifyItemBearerModRes:                    "id-E-RABModifyItemBearerModRes",
	idERABReleaseItem:                               "id-E-RABReleaseItem",
	idERABSetupItemBearerSURes:                      "id-E-RABSetupItemBearerSURes",
	idSecurityContext:                               "id-SecurityContext",
	idHandoverRestrictionList:                       "id-HandoverRestrictionList",
	idUEPagingID:                                    "id-UEPagingID",
	idpagingDRX:                                     "id-pagingDRX",
	idTAIList:                                       "id-TAIList",
	idTAIItem:                                       "id-TAIItem",
	idERABFailedToSetupListCtxtSURes:                "id-E-RABFailedToSetupListCtxtSURes",
	idERABReleaseItemHOCmd:                          "id-E-RABReleaseItemHOCmd",
	idERABSetupItemCtxtSURes:                        "id-E-RABSetupItemCtxtSURes",
	idERABSetupListCtxtSURes:                        "id-E-RABSetupListCtxtSURes",
	idERABToBeSetupItemCtxtSUReq:                    "id-E-RABToBeSetupItemCtxtSUReq",
	idERABToBeSetupListHOReq:                        "id-E-RABToBeSetupListHOReq",
	idGERANtoLTEHOInformationRes:                    "id-GERANtoLTEHOInformationRes",
	idUTRANtoLTEHOInformationRes:                    "id-UTRANtoLTEHOInformationRes",
	idCriticalityDiagnostics:                        "id-CriticalityDiagnostics ",
	idGlobalENBID:                                   "id-Global-ENB-ID",
	ideNBname:                                       "id-eNBname",
	idMMEname:                                       "id-MMEname",
	idServedPLMNs:                                   "id-ServedPLMNs",
	idSupportedTAs:                                  "id-SupportedTAs",
	idTimeToWait:                                    "id-TimeToWait",
	iduEaggregateMaximumBitrate:                     "id-uEaggregateMaximumBitrate",
	idTAI:                                           "id-TAI",
	idERABReleaseListBearerRelComp:                  "id-E-RABReleaseListBearerRelComp",
	idcdma2000PDU:                                   "id-cdma2000PDU",
	idcdma2000RATType:                               "id-cdma2000RATType",
	idcdma2000SectorID:                              "id-cdma2000SectorID",
	idSecurityKey:                                   "id-SecurityKey",
	idUERadioCapability:                             "id-UERadioCapability",
	idGUMMEIID:                                      "id-GUMMEI-ID",
	idERABInformationListItem:                       "id-E-RABInformationListItem",
	idDirectForwardingPathAvailability:              "id-Direct-Forwarding-Path-Availability",
	idUEIdentityIndexValue:                          "id-UEIdentityIndexValue",
	idcdma2000HOStatus:                              "id-cdma2000HOStatus",
	idcdma2000HORequiredIndication:                  "id-cdma2000HORequiredIndication",
	idEUTRANTraceID:                                 "id-E-UTRAN-Trace-ID",
	idRelativeMMECapacity:                           "id-RelativeMMECapacity",
	idSourceMMEUES1APID:                             "id-SourceMME-UE-S1AP-ID",
	idBearersSubjectToStatusTransferItem:            "id-Bearers-SubjectToStatusTransfer-Item",
	ideNBStatusTransferTransparentContainer:         "id-eNB-StatusTransfer-TransparentContainer",
	idUEassociatedLogicalS1ConnectionItem:           "id-UE-associatedLogicalS1-ConnectionItem",
	idResetType:                                     "id-ResetType",
	idUEassociatedLogicalS1ConnectionListResAck:     "id-UE-associatedLogicalS1-ConnectionListResAck",
	idERABToBeSwitchedULItem:                        "id-E-RABToBeSwitchedULItem",
	idERABToBeSwitchedULList:                        "id-E-RABToBeSwitchedULList",
	idSTMSI:                                         "id-S-TMSI",
	idcdma2000OneXRAND:                              "id-cdma2000OneXRAND",
	idRequestType:                                   "id-RequestType",
	idUES1APIDs:                                     "id-UE-S1AP-IDs",
	idEUTRANCGI:                                     "id-EUTRAN-CGI",
	idOverloadResponse:                              "id-OverloadResponse",
	idcdma2000OneXSRVCCInfo:                         "id-cdma2000OneXSRVCCInfo",
	idERABFailedToBeReleasedList:                    "id-E-RABFailedToBeReleasedList",
	idSourceToTargetTransparentContainer:            "id-Source-ToTarget-TransparentContainer",
	idServedGUMMEIs:                                 "id-ServedGUMMEIs",
	idSubscriberProfileIDforRFP:                     "id-SubscriberProfileIDforRFP",
	idUESecurityCapabilities:                        "id-UESecurityCapabilities",
	idCSFallbackIndicator:                           "id-CSFallbackIndicator",
	idCNDomain:                                      "id-CNDomain",
	idERABReleasedList:                              "id-E-RABReleasedList",
	idMessageIdentifier:                             "id-MessageIdentifier",
	idSerialNumber:                                  "id-SerialNumber",
	idWarningAreaList:                               "id-WarningAreaList",
	idRepetitionPeriod:                              "id-RepetitionPeriod",
	idNumberofBroadcastRequest:                      "id-NumberofBroadcastRequest",
	idWarningType:                                   "id-WarningType",
	idWarningSecurityInfo:                           "id-WarningSecurityInfo",
	idDataCodingScheme:                              "id-DataCodingScheme",
	idWarningMessageContents:                        "id-WarningMessageContents",
	idBroadcastCompletedAreaList:                    "id-BroadcastCompletedAreaList",
	idInterSystemInformationTransferTypeEDT:         "id-Inter-SystemInformationTransferTypeEDT",
	idInterSystemInformationTransferTypeMDT:         "id-Inter-SystemInformationTransferTypeMDT",
	idTargetToSourceTransparentContainer:            "id-Target-ToSource-TransparentContainer",
	idSRVCCOperationPossible:                        "id-SRVCCOperationPossible",
	idSRVCCHOIndication:                             "id-SRVCCHOIndication",
	idNASDownlinkCount:                              "id-NAS-DownlinkCount",
	idCSGId:                                         "id-CSG-Id",
	idCSGIdList:                                     "id-CSG-IdList",
	idSONConfigurationTransferECT:                   "id-SONConfigurationTransferECT",
	idSONConfigurationTransferMCT:                   "id-SONConfigurationTransferMCT",
	idTraceCollectionEntityIPAddress:                "id-TraceCollectionEntityIPAddress",
	idMSClassmark2:                                  "id-MSClassmark2",
	idMSClassmark3:                                  "id-MSClassmark3",
	idRRCEstablishmentCause:                         "id-RRC-Establishment-Cause",
	idNASSecurityParametersfromEUTRAN:               "id-NASSecurityParametersfromE-UTRAN",
	idNASSecurityParameterstoEUTRAN:                 "id-NASSecurityParameterstoE-UTRAN",
	idDefaultPagingDRX:                              "id-DefaultPagingDRX",
	idSourceToTargetTransparentContainerSecondary:   "id-Source-ToTarget-TransparentContainer-Secondary",
	idTargetToSourceTransparentContainerSecondary:   "id-Target-ToSource-TransparentContainer-Secondary",
	idEUTRANRoundTripDelayEstimationInfo:            "id-EUTRANRoundTripDelayEstimationInfo",
	idBroadcastCancelledAreaList:                    "id-BroadcastCancelledAreaList",
	idConcurrentWarningMessageIndicator:             "id-ConcurrentWarningMessageIndicator",
	idDataForwardingNotPossible:                     "id-Data-Forwarding-Not-Possible",
	idExtendedRepetitionPeriod:                      "id-ExtendedRepetitionPeriod",
	idCellAccessMode:                                "id-CellAccessMode",
	idCSGMembershipStatus:                           "id-CSGMembershipStatus ",
	idLPPaPDU:                                       "id-LPPa-PDU",
	idRoutingID:                                     "id-Routing-ID",
	idTimeSynchronisationInfo:                       "id-Time-Synchronisation-Info",
	idPSServiceNotAvailable:                         "id-PS-ServiceNotAvailable",
	idPagingPriority:                                "id-PagingPriority",
	idx2TNLConfigurationInfo:                        "id-x2TNLConfigurationInfo",
	ideNBX2ExtendedTransportLayerAddresses:          "id-eNBX2ExtendedTransportLayerAddresses",
	idGUMMEIList:                                    "id-GUMMEIList",
	idGWTransportLayerAddress:                       "id-GW-TransportLayerAddress",
	idCorrelationID:                                 "id-Correlation-ID",
	idSourceMMEGUMMEI:                               "id-SourceMME-GUMMEI",
	idMMEUES1APID2:                                  "id-MME-UE-S1AP-ID-2",
	idRegisteredLAI:                                 "id-RegisteredLAI",
	idRelayNodeIndicator:                            "id-RelayNode-Indicator",
	idTrafficLoadReductionIndication:                "id-TrafficLoadReductionIndication",
	idMDTConfiguration:                              "id-MDTConfiguration",
	idMMERelaySupportIndicator:                      "id-MMERelaySupportIndicator",
	idGWContextReleaseIndication:                    "id-GWContextReleaseIndication",
	idManagementBasedMDTAllowed:                     "id-ManagementBasedMDTAllowed",
	idPrivacyIndicator:                              "id-PrivacyIndicator",
	idTimeUEStayedInCellEnhancedGranularity:         "id-Time-UE-StayedInCell-EnhancedGranularity",
	idHOCause:                                       "id-HO-Cause",
	idVoiceSupportMatchIndicator:                    "id-VoiceSupportMatchIndicator",
	idGUMMEIType:                                    "id-GUMMEIType",
	idM3Configuration:                               "id-M3Configuration",
	idM4Configuration:                               "id-M4Configuration",
	idM5Configuration:                               "id-M5Configuration",
	idMDTLocationInfo:                               "id-MDT-Location-Info",
	idMobilityInformation:                           "id-MobilityInformation",
	idTunnelInformationforBBF:                       "id-Tunnel-Information-for-BBF",
	idManagementBasedMDTPLMNList:                    "id-ManagementBasedMDTPLMNList",
	idSignallingBasedMDTPLMNList:                    "id-SignallingBasedMDTPLMNList",
	idULCOUNTValueExtended:                          "id-ULCOUNTValueExtended",
	idDLCOUNTValueExtended:                          "id-DLCOUNTValueExtended",
	idReceiveStatusOfULPDCPSDUsExtended:             "id-ReceiveStatusOfULPDCPSDUsExtended",
	idECGIListForRestart:                            "id-ECGIListForRestart",
	idSIPTOCorrelationID:                            "id-SIPTO-Correlation-ID",
	idSIPTOLGWTransportLayerAddress:                 "id-SIPTO-L-GW-TransportLayerAddress",
	idTransportInformation:                          "id-TransportInformation",
	idLHNID:                                         "id-LHN-ID",
	idAdditionalCSFallbackIndicator:                 "id-AdditionalCSFallbackIndicator",
	idTAIListForRestart:                             "id-TAIListForRestart",
	idUserLocationInformation:                       "id-UserLocationInformation",
	idEmergencyAreaIDListForRestart:                 "id-EmergencyAreaIDListForRestart",
	idKillAllWarningMessages:                        "id-KillAllWarningMessages",
	idMaskedIMEISV:                                  "id-Masked-IMEISV",
	ideNBIndirectX2TransportLayerAddresses:          "id-eNBIndirectX2TransportLayerAddresses",
	iduEHistoryInformationFromTheUE:                 "id-uE-HistoryInformationFromTheUE",
	idProSeAuthorized:                               "id-ProSeAuthorized",
	idExpectedUEBehaviour:                           "id-ExpectedUEBehaviour",
	idLoggedMBSFNMDT:                                "id-LoggedMBSFNMDT",
	idUERadioCapabilityForPaging:                    "id-UERadioCapabilityForPaging",
	idERABToBeModifiedListBearerModInd:              "id-E-RABToBeModifiedListBearerModInd",
	idERABToBeModifiedItemBearerModInd:              "id-E-RABToBeModifiedItemBearerModInd",
	idERABNotToBeModifiedListBearerModInd:           "id-E-RABNotToBeModifiedListBearerModInd",
	idERABNotToBeModifiedItemBearerModInd:           "id-E-RABNotToBeModifiedItemBearerModInd",
	idERABModifyListBearerModConf:                   "id-E-RABModifyListBearerModConf",
	idERABModifyItemBearerModConf:                   "id-E-RABModifyItemBearerModConf",
	idERABFailedToModifyListBearerModConf:           "id-E-RABFailedToModifyListBearerModConf",
	idSONInformationReport:                          "id-SON-Information-Report",
	idMutingAvailabilityIndication:                  "id-Muting-Availability-Indication",
	idMutingPatternInformation:                      "id-Muting-Pattern-Information",
	idSynchronisationInformation:                    "id-Synchronisation-Information",
	idERABToBeReleasedListBearerModConf:             "id-E-RABToBeReleasedListBearerModConf",
	idAssistanceDataForPaging:                       "id-AssistanceDataForPaging",
	idCellIdentifierAndCELevelForCECapableUEs:       "id-CellIdentifierAndCELevelForCECapableUEs",
	idInformationOnRecommendedCellsAndENBsForPaging: "id-InformationOnRecommendedCellsAndENBsForPaging",
	idRecommendedCellItem:                           "id-RecommendedCellItem",
	idRecommendedENBItem:                            "id-RecommendedENBItem",
	idProSeUEtoNetworkRelaying:                      "id-ProSeUEtoNetworkRelaying",
	idULCOUNTValuePDCPSNlength18:                    "id-ULCOUNTValuePDCP-SNlength18",
	idDLCOUNTValuePDCPSNlength18:                    "id-DLCOUNTValuePDCP-SNlength18",
	idReceiveStatusOfULPDCPSDUsPDCPSNlength18:       "id-ReceiveStatusOfULPDCPSDUsPDCP-SNlength18",
	idM6Configuration:                               "id-M6Configuration",
	idM7Configuration:                               "id-M7Configuration",
	idPWSfailedECGIList:                             "id-PWSfailedECGIList",
	idMMEGroupID:                                    "id-MME-Group-ID",
	idAdditionalGUTI:                                "id-Additional-GUTI",
	idS1Message:                                     "id-S1-Message",
	idCSGMembershipInfo:                             "id-CSGMembershipInfo",
	idPagingeDRXInformation:                         "id-Paging-eDRXInformation",
	idUERetentionInformation:                        "id-UE-RetentionInformation",
	idUEUsageType:                                   "id-UE-Usage-Type",
	idextendedUEIdentityIndexValue:                  "id-extended-UEIdentityIndexValue",
	idRATType:                                       "id-RAT-Type",
	idBearerType:                                    "id-BearerType",
	idNBIoTDefaultPagingDRX:                         "id-NB-IoT-DefaultPagingDRX",
	idERABFailedToResumeListResumeReq:               "id-E-RABFailedToResumeListResumeReq",
	idERABFailedToResumeItemResumeReq:               "id-E-RABFailedToResumeItemResumeReq",
	idERABFailedToResumeListResumeRes:               "id-E-RABFailedToResumeListResumeRes",
	idERABFailedToResumeItemResumeRes:               "id-E-RABFailedToResumeItemResumeRes",
	idNBIoTPagingeDRXInformation:                    "id-NB-IoT-Paging-eDRXInformation",
	idV2XServicesAuthorized:                         "id-V2XServicesAuthorized",
	idUEUserPlaneCIoTSupportIndicator:               "id-UEUserPlaneCIoTSupportIndicator ",
	idCEmodeBSupportIndicator:                       "id-CE-mode-B-SupportIndicator ",
	idSRVCCOperationNotPossible:                     "id-SRVCCOperationNotPossible",
	idNBIoTUEIdentityIndexValue:                     "id-NB-IoT-UEIdentityIndexValue ",
	idRRCResumeCause:                                "id-RRC-Resume-Cause",
	idDCNID:                                         "id-DCN-ID",
	idServedDCNs:                                    "id-ServedDCNs",
	idUESidelinkAggregateMaximumBitrate:             "id-UESidelinkAggregateMaximumBitrate ",
	idDLNASPDUDeliveryAckRequest:                    "id-DLNASPDUDeliveryAckRequest",
	idCoverageLevel:                                 "id-Coverage-Level ",
	idEnhancedCoverageRestricted:                    "id-EnhancedCoverageRestricted",
	idUELevelQoSParameters:                          "id-UE-Level-QoS-Parameters",
	idDLCPSecurityInformation:                       "id-DL-CP-SecurityInformation",
	idULCPSecurityInformation:                       "id-UL-CP-SecurityInformation",
	idextendedeRABMaximumBitrateDL:                  "id-extended-e-RAB-MaximumBitrateDL",
	idextendedeRABMaximumBitrateUL:                  "id-extended-e-RAB-MaximumBitrateUL",
	idextendedeRABGuaranteedBitrateDL:               "id-extended-e-RAB-GuaranteedBitrateDL",
	idextendedeRABGuaranteedBitrateUL:               "id-extended-e-RAB-GuaranteedBitrateUL",
	idextendeduEaggregateMaximumBitRateDL:           "id-extended-uEaggregateMaximumBitRateDL",
	idextendeduEaggregateMaximumBitRateUL:           "id-extended-uEaggregateMaximumBitRateUL",
	idNRrestrictioninEPSasSecondaryRAT:              "id-NRrestrictioninEPSasSecondaryRAT",
	idUEAppLayerMeasConfig:                          "id-UEAppLayerMeasConfig",
	idUEApplicationLayerMeasurementCapability:       "id-UE-Application-Layer-Measurement-Capability",
	idSecondaryRATDataUsageReportList:               "id-SecondaryRATDataUsageReportList",
	idSecondaryRATDataUsageReportItem:               "id-SecondaryRATDataUsageReportItem",
	idHandoverFlag:                                  "id-HandoverFlag",
	idERABUsageReportItem:                           "id-E-RABUsageReportItem",
	idSecondaryRATDataUsageRequest:                  "id-SecondaryRATDataUsageRequest",
	idNRUESecurityCapabilities:                      "id-NRUESecurityCapabilities",
	idUnlicensedSpectrumRestriction:                 "id-UnlicensedSpectrumRestriction",
	idCEModeBRestricted:                             "id-CE-ModeBRestricted",
	idLTEMIndication:                                "id-LTE-M-Indication",
	idDownlinkPacketLossRate:                        "id-DownlinkPacketLossRate",
	idUplinkPacketLossRate:                          "id-UplinkPacketLossRate",
	idUECapabilityInfoRequest:                       "id-UECapabilityInfoRequest",
	idserviceType:                                   "id-serviceType",
	idAerialUEsubscriptionInformation:               "id-AerialUEsubscriptionInformation",
	idSubscriptionBasedUEDifferentiationInfo:        "id-Subscription-Based-UE-DifferentiationInfo",
	idEndIndication:                                 "id-EndIndication",
	idEDTSession:                                    "id-EDT-Session",
	idCNTypeRestrictions:                            "id-CNTypeRestrictions",
	idPendingDataIndication:                         "id-PendingDataIndication",
	idBluetoothMeasurementConfiguration:             "id-BluetoothMeasurementConfiguration",
	idWLANMeasurementConfiguration:                  "id-WLANMeasurementConfiguration",
	idWarningAreaCoordinates:                        "id-WarningAreaCoordinates",
	idNRrestrictionin5GS:                            "id-NRrestrictionin5GS",
	idPSCellInformation:                             "id-PSCellInformation",
	idLastNGRANPLMNIdentity:                         "id-LastNG-RANPLMNIdentity",
	idConnectedengNBList:                            "id-ConnectedengNBList",
	idConnectedengNBToAddList:                       "id-ConnectedengNBToAddList",
	idConnectedengNBToRemoveList:                    "id-ConnectedengNBToRemoveList",
	idENDCSONConfigurationTransferECT:               "id-EN-DCSONConfigurationTransfer-ECT",
	idENDCSONConfigurationTransferMCT:               "id-EN-DCSONConfigurationTransfer-MCT",
	idIMSvoiceEPSfallbackfrom5G:                     "id-IMSvoiceEPSfallbackfrom5G",
	idTimeSinceSecondaryNodeRelease:                 "id-TimeSinceSecondaryNodeRelease",
	idRequestTypeAdditionalInfo:                     "id-RequestTypeAdditionalInfo",
	idAdditionalRRMPriorityIndex:                    "id-AdditionalRRMPriorityIndex",
	idContextatSource:                               "id-ContextatSource",
	idIABAuthorized:                                 "id-IAB-Authorized",
	idIABNodeIndication:                             "id-IAB-Node-Indication",
	idIABSupported:                                  "id-IAB-Supported",
	idDataSize:                                      "id-DataSize",
	idEthernetType:                                  "id-Ethernet-Type",
	idNRV2XServicesAuthorized:                       "id-NRV2XServicesAuthorized",
	idNRUESidelinkAggregateMaximumBitrate:           "id-NRUESidelinkAggregateMaximumBitrate",
	idPC5QoSParameters:                              "id-PC5QoSParameters",
	idIntersystemSONConfigurationTransferMCT:        "id-IntersystemSONConfigurationTransferMCT",
	idIntersystemSONConfigurationTransferECT:        "id-IntersystemSONConfigurationTransferECT",
	idIntersystemMeasurementConfiguration:           "id-IntersystemMeasurementConfiguration",
	idSourceNodeID:                                  "id-SourceNodeID",
	idNBIoTRLFReportContainer:                       "id-NB-IoT-RLF-Report-Container",
	idUERadioCapabilityID:                           "id-UERadioCapabilityID",
	idUERadioCapabilityNRFormat:                     "id-UERadioCapability-NR-Format",
	idMDTConfigurationNR:                            "id-MDTConfigurationNR",
	idDAPSRequestInfo:                               "id-DAPSRequestInfo",
	idDAPSResponseInfoList:                          "id-DAPSResponseInfoList",
	idDAPSResponseInfoItem:                          "id-DAPSResponseInfoItem",
	idNotifySourceeNB:                               "id-NotifySourceeNB",
	ideNBEarlyStatusTransferTransparentContainer:    "id-eNB-EarlyStatusTransfer-TransparentContainer",
	idBearersSubjectToEarlyStatusTransferItem:       "id-Bearers-SubjectToEarlyStatusTransfer-Item",
	idWUSAssistanceInformation:                      "id-WUS-Assistance-Information",
	idNBIoTPagingDRX:                                "id-NB-IoT-PagingDRX",
	idTraceCollectionEntityURI:                      "id-TraceCollectionEntityURI",
	idEmergencyIndicator:                            "id-EmergencyIndicator",
}
