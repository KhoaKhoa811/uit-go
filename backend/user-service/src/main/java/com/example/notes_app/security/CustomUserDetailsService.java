package com.example.notes_app.security;

import java.util.Collections;

import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

import com.example.notes_app.entity.AccountEntity;
import com.example.notes_app.repository.AccountRepository;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class CustomUserDetailsService implements UserDetailsService {
    private final AccountRepository accountRepository;

    @Override
    public UserDetails loadUserByUsername(String email) throws UsernameNotFoundException {
        AccountEntity accountEntity = accountRepository.findByEmail(email)
                .orElseThrow(() -> new UsernameNotFoundException("Account not found with email: " + email));

        // Trả về UserDetails đơn giản, không cần authority
        return org.springframework.security.core.userdetails.User
                .withUsername(accountEntity.getEmail())
                .password(accountEntity.getPassword())
                .authorities(Collections.emptyList()) // không có role/permission
                .build();
    }
}
