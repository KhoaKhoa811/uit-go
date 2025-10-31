package com.example.notes_app.controller;

import org.springframework.data.domain.Pageable;
import org.springframework.http.ResponseEntity;
import org.springframework.security.core.Authentication;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.ModelAttribute;
import org.springframework.web.bind.annotation.PatchMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.example.notes_app.config.PaginationProperties;
import com.example.notes_app.dto.ApiResponse;
import com.example.notes_app.dto.PagedResponse;
import com.example.notes_app.dto.PaginationRequest;
import com.example.notes_app.dto.account.AccountResponse;
import com.example.notes_app.dto.account.AccountUpdateRequest;
import com.example.notes_app.enums.ErrorCode;
import com.example.notes_app.service.AccountService;
import com.example.notes_app.utils.PaginationUtils;

import lombok.RequiredArgsConstructor;

@RestController
@RequestMapping("/api/v1/accounts")
@RequiredArgsConstructor
public class AccountController {
    private final AccountService accountService;
    private final PaginationProperties paginationProperties;

    @GetMapping
    public ResponseEntity<ApiResponse<PagedResponse<AccountResponse>>> getAllAccounts(@ModelAttribute PaginationRequest paginationRequest) {
        Pageable pageable = PaginationUtils.createPageable(paginationRequest, paginationProperties);
        ApiResponse<PagedResponse<AccountResponse>> accountResponse = ApiResponse.<PagedResponse<AccountResponse>>builder()
                .code(ErrorCode.ACCOUNT_GET_ALL.getCode())
                .data(accountService.getAllAccounts(pageable))
                .build();
        return ResponseEntity.ok(accountResponse);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ApiResponse<AccountResponse>> getAccountById(@PathVariable Integer id) {
        ApiResponse<AccountResponse> accountResponse = ApiResponse.<AccountResponse>builder()
                .code(ErrorCode.ACCOUNT_GET_ALL.getCode())
                .data(accountService.getAccountById(id))
                .build();
        return ResponseEntity.ok(accountResponse);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<ApiResponse<?>> deleteAccountById(@PathVariable Integer id) {
        accountService.deleteAccountById(id);
        ApiResponse<?> accountResponse = ApiResponse.builder()
                .code(ErrorCode.ACCOUNT_DELETED.getCode())
                .message(ErrorCode.ACCOUNT_DELETED.getMessage())
                .build();
        return ResponseEntity.ok(accountResponse);
    }

    @PatchMapping("/{id}")
    public ResponseEntity<ApiResponse<AccountResponse>> updateAccountById(
            @PathVariable Integer id,
            @RequestBody AccountUpdateRequest accountUpdateRequest
    ) {
        ApiResponse<AccountResponse> accountResponse = ApiResponse.<AccountResponse>builder()
                .code(ErrorCode.ACCOUNT_UPDATED.getCode())
                .data(accountService.updateAccountById(id, accountUpdateRequest))
                .build();
        return ResponseEntity.ok(accountResponse);
    }

    @GetMapping("/me")
    public ResponseEntity<ApiResponse<AccountResponse>> getMyInfo(Authentication authentication) {
        ApiResponse<AccountResponse> accountResponse = ApiResponse.<AccountResponse>builder()
                .code(ErrorCode.ACCOUNT_GET_ALL.getCode())
                .data(accountService.getAccountByEmail(authentication.getName()))
                .build();
        return ResponseEntity.ok(accountResponse);
    }

}
