package com.example.notes_app.service;

import org.springframework.data.domain.Pageable;

import com.example.notes_app.dto.PagedResponse;
import com.example.notes_app.dto.account.AccountResponse;
import com.example.notes_app.dto.account.AccountUpdateRequest;

public interface AccountService {
    PagedResponse<AccountResponse> getAllAccounts(Pageable pageable);
    AccountResponse getAccountById(Integer id);
    AccountResponse getAccountByEmail(String email);
    void deleteAccountById(Integer id);
    AccountResponse updateAccountById(Integer id, AccountUpdateRequest accountUpdateRequest);
}
