package com.example.notes_app.grpc.impl;

import com.example.notes_app.grpc.account.*;
import com.example.notes_app.service.AccountService;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.server.service.GrpcService;

@GrpcService
@RequiredArgsConstructor
public class AccountGrpcServiceImpl extends AccountServiceGrpc.AccountServiceImplBase {
    private final AccountService accountService;

    @Override
    public void getAccountById(AccountRequest request, StreamObserver<AccountResponse> responseObserver) {
        var result = accountService.getAccountById(request.getId());
        System.out.println(">> Email in gRPC: " + result.getEmail());
        AccountResponse response = AccountResponse.newBuilder()
                .setId(result.getId())
                .setEmail(result.getEmail())
                .build();
        responseObserver.onNext(response);
        responseObserver.onCompleted();
    }

    @Override
    public void deleteAccountById(AccountRequest request, StreamObserver<Empty> responseObserver) {
        accountService.deleteAccountById(request.getId());
        responseObserver.onNext(Empty.newBuilder().build());
        responseObserver.onCompleted();
    }
}
