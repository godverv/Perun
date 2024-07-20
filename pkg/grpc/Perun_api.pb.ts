/* eslint-disable */
// @ts-nocheck

/**
 * This file is a generated Typescript file for GRPC Gateway, DO NOT MODIFY
 */

import * as fm from "../fetch.pb";


export type VersionRequest = Record<string, never>;

export type VersionResponse = {
  version?: string;
};

export type Version = Record<string, never>;

export type Node = {
  name?: string;
  addr?: string;
  port?: number;
  customVelezKeyPath?: string;
  securityDisabled?: boolean;
};

export type Ssh = {
  keyBase64?: Uint8Array;
  port?: string;
  username?: string;
};

export type ConnectVelezRequest = {
  node?: Node;
  ssh?: Ssh;
};

export type ConnectVelezResponse = Record<string, never>;

export type ConnectVelez = Record<string, never>;

export type ListPaging = {
  limit?: number;
  offset?: number;
};

export type ListNodesRequest = {
  searchPattern?: string;
  paging?: ListPaging;
};

export type ListNodesResponse = {
  nodes?: Node[];
};

export type ListNodes = Record<string, never>;

export type CreateServiceRequest = {
  imageName?: string;
  serviceName?: string;
  replicas?: number;
};

export type CreateServiceResponse = Record<string, never>;

export type CreateService = Record<string, never>;

export type RefreshServiceRequest = {
  serviceName?: string;
};

export type RefreshServiceResponse = Record<string, never>;

export type RefreshService = Record<string, never>;

export class PerunAPI {
  static Version(this:void, req: VersionRequest, initReq?: fm.InitReq): Promise<VersionResponse> {
    return fm.fetchRequest<VersionResponse>(`/version?${fm.renderURLSearchParams(req, [])}`, {...initReq, method: "GET"});
  }
  static ConnectVelez(this:void, req: ConnectVelezRequest, initReq?: fm.InitReq): Promise<ConnectVelezResponse> {
    return fm.fetchRequest<ConnectVelezResponse>(`/velez`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static ListNodes(this:void, req: ListNodesRequest, initReq?: fm.InitReq): Promise<ListNodesResponse> {
    return fm.fetchRequest<ListNodesResponse>(`/velez/list`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static CreateService(this:void, req: CreateServiceRequest, initReq?: fm.InitReq): Promise<CreateServiceResponse> {
    return fm.fetchRequest<CreateServiceResponse>(`/service/new`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
  static RefreshService(this:void, req: RefreshServiceRequest, initReq?: fm.InitReq): Promise<RefreshServiceResponse> {
    return fm.fetchRequest<RefreshServiceResponse>(`/service/${req.serviceName}/refresh`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
}