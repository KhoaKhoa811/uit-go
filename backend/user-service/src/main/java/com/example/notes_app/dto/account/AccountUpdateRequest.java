package com.example.notes_app.dto.account;

import java.util.List;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class AccountUpdateRequest {
    private Boolean enabled;
    private List<Integer> roleIds;
}
