package com.example.myapplication.ui

import android.os.Bundle
import android.widget.*
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.myapplication.R
import com.example.myapplication.api.ApiClient
import com.example.myapplication.model.RegisterRequest
import kotlinx.coroutines.launch

class RegisterActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_register)

        val email = findViewById<EditText>(R.id.etEmail)
        val password = findViewById<EditText>(R.id.etPassword)
        val btnRegister = findViewById<Button>(R.id.btnRegister)

        btnRegister.setOnClickListener {
            lifecycleScope.launch {
                val response = ApiClient.instance.register(
                    RegisterRequest(email.text.toString(), password.text.toString(), listOf(1))
                )
                if (response.isSuccessful) {
                    Toast.makeText(this@RegisterActivity, "Register success!", Toast.LENGTH_SHORT).show()
                    finish()
                } else {
                    Toast.makeText(this@RegisterActivity, "Register failed!", Toast.LENGTH_SHORT).show()
                }
            }
        }
    }
}