
�

grpc/Perun_api.proto	perun_apigoogle/api/annotations.proto":
Version	
Request$
Response
version (	Rversion"�
Node
name (	Rname
addr (	Raddr6
custom_velez_key_path (	H RcustomVelezKeyPath�0
security_disabled (HRsecurityDisabled�B
_custom_velez_key_pathB
_security_disabled"T
Ssh

key_base64 (R	keyBase64
addr (	Raddr
username (	Rusername"l
ConnectVelezP
Request#
node (2.perun_api.NodeRnode 
ssh (2.perun_api.SshRssh

Response":

ListPaging
limit (Rlimit
offset (Roffset"�
	ListNodesw
Request*
search_pattern (	H RsearchPattern�-
paging (2.perun_api.ListPagingRpagingB
_search_pattern1
Response%
nodes (2.perun_api.NodeRnodes"�

RunServicez
Request

image_name (	R	imageName!
service_name (	RserviceName-
replication_factor (RreplicationFactor

Response2�
PerunAPIT
Version.perun_api.Version.Request.perun_api.Version.Response"���
/versiond
ConnectVelez.perun_api.ConnectVelez.Request .perun_api.ConnectVelez.Response"���"/velez:*`
	ListNodes.perun_api.ListNodes.Request.perun_api.ListNodes.Response"���"/velez/list:*d

RunService.perun_api.RunService.Request.perun_api.RunService.Response"���"/service/run:*BZ
/perun_apibproto3