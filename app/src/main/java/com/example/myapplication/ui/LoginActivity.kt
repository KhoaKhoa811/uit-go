package com.example.myapplication.ui

import android.content.Intent
import android.os.Bundle
import android.widget.*
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.lifecycleScope
import com.example.myapplication.MainActivity
import com.example.myapplication.R
import com.example.myapplication.api.ApiClient
import com.example.myapplication.model.LoginRequest
import com.example.myapplication.util.PrefsManager
import kotlinx.coroutines.launch
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

class LoginActivity : AppCompatActivity() {

    private lateinit var etEmail: EditText
    private lateinit var etPassword: EditText
    private lateinit var btnLogin: Button
    private lateinit var btnGoRegister: Button
    private lateinit var progressBar: ProgressBar

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_login)

        initViews()
        setupListeners()
    }

    private fun initViews() {
        etEmail = findViewById(R.id.etEmail)
        etPassword = findViewById(R.id.etPassword)
        btnLogin = findViewById(R.id.btnLogin)
        btnGoRegister = findViewById(R.id.btnGoRegister)
        progressBar = findViewById(R.id.progressBar) // thêm vào layout nếu chưa có
    }

    private fun setupListeners() {
        btnLogin.setOnClickListener {
            val email = etEmail.text.toString().trim()
            val password = etPassword.text.toString().trim()

            if (email.isEmpty() || password.isEmpty()) {
                Toast.makeText(this, "Please enter email and password", Toast.LENGTH_SHORT).show()
                return@setOnClickListener
            }

            loginUser(email, password)
        }

        btnGoRegister.setOnClickListener {
            startActivity(Intent(this, RegisterActivity::class.java))
        }
    }

    private fun loginUser(email: String, password: String) {
        progressBar.visibility = ProgressBar.VISIBLE
        lifecycleScope.launch {
            try {
                val response = ApiClient.instance.login(LoginRequest(email, password))
                withContext(Dispatchers.Main) {
                    progressBar.visibility = ProgressBar.GONE
                    if (response.isSuccessful) {
                        val tokenResponse = response.body()
                        tokenResponse?.let {
                            PrefsManager.saveTokens(
                                context = this@LoginActivity,
                                accessToken = it.data.accessToken,
                                refreshToken = it.data.refreshToken
                            )
                            Toast.makeText(this@LoginActivity, "Login successful!", Toast.LENGTH_SHORT).show()
                            startActivity(Intent(this@LoginActivity, MainActivity::class.java))
                            finish()
                        }
                    } else {
                        Toast.makeText(this@LoginActivity, "Invalid email or password", Toast.LENGTH_SHORT).show()
                    }
                }
            } catch (e: Exception) {
                withContext(Dispatchers.Main) {
                    progressBar.visibility = ProgressBar.GONE
                    Toast.makeText(this@LoginActivity, "Network error: ${e.localizedMessage}", Toast.LENGTH_LONG).show()
                }
            }
        }
    }
}
