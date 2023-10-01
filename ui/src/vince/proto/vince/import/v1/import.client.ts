// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "vince/import/v1/import.proto" (package "v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Import } from "./import";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { ImportResponse } from "./import";
import type { ImportRequest } from "./import";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service v1.Import
 */
export interface IImportClient {
    /**
     * @generated from protobuf rpc: Import(v1.ImportRequest) returns (v1.ImportResponse);
     */
    import(input: ImportRequest, options?: RpcOptions): UnaryCall<ImportRequest, ImportResponse>;
}
/**
 * @generated from protobuf service v1.Import
 */
export class ImportClient implements IImportClient, ServiceInfo {
    typeName = Import.typeName;
    methods = Import.methods;
    options = Import.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Import(v1.ImportRequest) returns (v1.ImportResponse);
     */
    import(input: ImportRequest, options?: RpcOptions): UnaryCall<ImportRequest, ImportResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<ImportRequest, ImportResponse>("unary", this._transport, method, opt, input);
    }
}