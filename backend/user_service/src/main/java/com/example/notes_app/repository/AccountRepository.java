package com.example.notes_app.repository;

import com.example.notes_app.entity.AccountEntity;
import com.example.notes_app.entity.RoleEntity;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

import java.util.Optional;

public interface AccountRepository extends JpaRepository<AccountEntity, Integer> {
    Optional<AccountEntity> findByEmail(String email);
    boolean existsByEmail(String email);
    @Query("SELECT a FROM AccountEntity a LEFT JOIN FETCH a.roles WHERE a.id = :id")
    Optional<AccountEntity> findByAccountId(@Param("id") Integer id);

}
