import {HttpCore, HttpMethod, RequestInfo, ResponseHandler} from '@microsoft/kiota-abstractions';
import {ChildFoldersRequestBuilder} from '../childFolders/childFoldersRequestBuilder';
import {MailFolder} from '../mailFolder';
import {MessageRuleRequestBuilder} from '../messageRules/item/messageRuleRequestBuilder';
import {MessageRulesRequestBuilder} from '../messageRules/messageRulesRequestBuilder';
import {MessageRequestBuilder} from '../messages/item/messageRequestBuilder';
import {MessagesRequestBuilder} from '../messages/messagesRequestBuilder';
import {MultiValueLegacyExtendedPropertyRequestBuilder} from '../messages/multiValueExtendedProperties/item/multiValueLegacyExtendedPropertyRequestBuilder';
import {MultiValueExtendedPropertiesRequestBuilder} from '../messages/multiValueExtendedProperties/multiValueExtendedPropertiesRequestBuilder';
import {SingleValueLegacyExtendedPropertyRequestBuilder} from '../messages/singleValueExtendedProperties/item/singleValueLegacyExtendedPropertyRequestBuilder';
import {SingleValueExtendedPropertiesRequestBuilder} from '../messages/singleValueExtendedProperties/singleValueExtendedPropertiesRequestBuilder';

export class MailFolderRequestBuilder {
    public get childFolders(): ChildFoldersRequestBuilder {
        const builder = new ChildFoldersRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get messageRules(): MessageRulesRequestBuilder {
        const builder = new MessageRulesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get messages(): MessagesRequestBuilder {
        const builder = new MessagesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get multiValueExtendedProperties(): MultiValueExtendedPropertiesRequestBuilder {
        const builder = new MultiValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public get singleValueExtendedProperties(): SingleValueExtendedPropertiesRequestBuilder {
        const builder = new SingleValueExtendedPropertiesRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment;
        return builder;
    }
    public readonly get = async (q?: {
                    select?: string[],
                    expand?: string[]
                    } | undefined, h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<MailFolder | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.GET,
            queryParameters: q,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<MailFolder>(requestInfo, responseHandler);
    }
    public readonly patch = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.PATCH,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    public readonly delete = async (h?: {} | undefined, responseHandler?: ResponseHandler | undefined) : Promise<object | undefined> => {
        const requestInfo = {
            URI: this.currentPath ? new URL(this.currentPath): null,
            headers: h,
            httpMethod: HttpMethod.DELETE,
        } as RequestInfo;
        return await this.httpCore?.sendAsync<object>(requestInfo, responseHandler);
    }
    private readonly pathSegment: string = "";
    public currentPath?: string | undefined;
    public httpCore?: HttpCore | undefined;
    public readonly childFoldersById = (id: String) : MailFolderRequestBuilder => {
        const builder = new MailFolderRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/childFolders/" + id;
        return builder;
    }
    public readonly messageRulesById = (id: String) : MessageRuleRequestBuilder => {
        const builder = new MessageRuleRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/messageRules/" + id;
        return builder;
    }
    public readonly messagesById = (id: String) : MessageRequestBuilder => {
        const builder = new MessageRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/messages/" + id;
        return builder;
    }
    public readonly multiValueExtendedPropertiesById = (id: String) : MultiValueLegacyExtendedPropertyRequestBuilder => {
        const builder = new MultiValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/multiValueExtendedProperties/" + id;
        return builder;
    }
    public readonly singleValueExtendedPropertiesById = (id: String) : SingleValueLegacyExtendedPropertyRequestBuilder => {
        const builder = new SingleValueLegacyExtendedPropertyRequestBuilder();
        builder.currentPath = (this.currentPath && this.currentPath) + this.pathSegment + "/singleValueExtendedProperties/" + id;
        return builder;
    }
}
