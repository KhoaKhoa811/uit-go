package com.example.notes_app.dto.account;

import java.util.List;

import com.example.notes_app.dto.role.RoleResponse;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class AccountResponse {
    private Integer id;
    private String email;
    private Boolean enabled;
    private List<RoleResponse> roles;
}
