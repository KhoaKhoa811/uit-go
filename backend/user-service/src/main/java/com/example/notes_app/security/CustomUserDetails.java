package com.example.notes_app.security;

import java.util.Collection;

import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import com.example.notes_app.entity.AccountEntity;

import lombok.AllArgsConstructor;
import lombok.RequiredArgsConstructor;

@AllArgsConstructor
@RequiredArgsConstructor
public class CustomUserDetails implements UserDetails {
    private final AccountEntity accountEntity;

    private Collection<? extends GrantedAuthority> authorities;

    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return null;
    }

    @Override
    public String getPassword() {
        return accountEntity.getPassword();
    }

    @Override
    public boolean isEnabled() {
        return accountEntity.getEnabled();
    }

    @Override
    public String getUsername() {
        return accountEntity.getEmail();
    }
}
