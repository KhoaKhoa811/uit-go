package com.example.myapplication.model

data class RegisterRequest(
    val email: String,
    val password: String,
    val roleIds: List<Int>
)