// tslint:disable
// eslint-disable
// Generated by Microsoft Kiota
import { createDateTimeTimeZoneFromDiscriminatorValue, serializeDateTimeTimeZone, type DateTimeTimeZone } from './dateTimeTimeZone';
import { FollowupFlagStatus } from './followupFlagStatus';
import { type AdditionalDataHolder, type Parsable, type ParseNode, type SerializationWriter } from '@microsoft/kiota-abstractions';

export function createFollowupFlagFromDiscriminatorValue(parseNode: ParseNode | undefined) {
    if(!parseNode) throw new Error("parseNode cannot be undefined");
    return deserializeIntoFollowupFlag;
}
export function deserializeIntoFollowupFlag(followupFlag: FollowupFlag | undefined = {} as FollowupFlag) : Record<string, (node: ParseNode) => void> {
    return {
        "completedDateTime": n => { followupFlag.completedDateTime = n.getObjectValue<DateTimeTimeZone>(createDateTimeTimeZoneFromDiscriminatorValue); },
        "dueDateTime": n => { followupFlag.dueDateTime = n.getObjectValue<DateTimeTimeZone>(createDateTimeTimeZoneFromDiscriminatorValue); },
        "flagStatus": n => { followupFlag.flagStatus = n.getEnumValue<FollowupFlagStatus>(FollowupFlagStatus); },
        "startDateTime": n => { followupFlag.startDateTime = n.getObjectValue<DateTimeTimeZone>(createDateTimeTimeZoneFromDiscriminatorValue); },
    }
}
export interface FollowupFlag extends AdditionalDataHolder, Parsable {
    /**
     * Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
     */
    additionalData?: Record<string, unknown>;
    /**
     * The completedDateTime property
     */
    completedDateTime?: DateTimeTimeZone;
    /**
     * The dueDateTime property
     */
    dueDateTime?: DateTimeTimeZone;
    /**
     * The flagStatus property
     */
    flagStatus?: FollowupFlagStatus;
    /**
     * The startDateTime property
     */
    startDateTime?: DateTimeTimeZone;
}
export function serializeFollowupFlag(writer: SerializationWriter, followupFlag: FollowupFlag | undefined = {} as FollowupFlag) : void {
        writer.writeObjectValue<DateTimeTimeZone>("completedDateTime", followupFlag.completedDateTime, serializeDateTimeTimeZone);
        writer.writeObjectValue<DateTimeTimeZone>("dueDateTime", followupFlag.dueDateTime, serializeDateTimeTimeZone);
        writer.writeEnumValue<FollowupFlagStatus>("flagStatus", followupFlag.flagStatus);
        writer.writeObjectValue<DateTimeTimeZone>("startDateTime", followupFlag.startDateTime, serializeDateTimeTimeZone);
        writer.writeAdditionalData(followupFlag.additionalData);
}
// tslint:enable
// eslint-enable
