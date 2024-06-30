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
  customVelezKeyPath?: string;
  securityDisabled?: boolean;
};

export type Ssh = {
  keyBase64?: Uint8Array;
  addr?: string;
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

export type RunServiceRequest = {
  imageName?: string;
  serviceName?: string;
  replicationFactor?: number;
};

export type RunServiceResponse = Record<string, never>;

export type RunService = Record<string, never>;

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
  static RunService(this:void, req: RunServiceRequest, initReq?: fm.InitReq): Promise<RunServiceResponse> {
    return fm.fetchRequest<RunServiceResponse>(`/service/run`, {...initReq, method: "POST", body: JSON.stringify(req, fm.replacer)});
  }
}