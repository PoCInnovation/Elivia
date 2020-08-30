package com.poc.wsleon

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import com.poc.wsleon.plugin.core.PluginManager
import com.poc.wsleon.socket.LeonWebSocket
import com.poc.wsleon.ui.LeonView
import kotlinx.android.synthetic.main.activity_main.*
import okhttp3.*

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val chat = LeonView()
        chat.initRecyclerViewAdapter(chatView)

        val pluginManager = PluginManager(this, this, chat)
        val userButtonSend: Button = findViewById(R.id.userButtonSend)
        val userTextInput: EditText = findViewById(R.id.userTextInput)

        val client = OkHttpClient();
        val request: Request = Request.Builder().url("ws://192.168.0.14:8080/websocket").build()
        val listener = LeonWebSocket(chat)
        val ws: WebSocket = client.newWebSocket(request, listener)

        userButtonSend.setOnClickListener {
            val userRequest = userTextInput.text.toString()

            listener.onSend(ws, userRequest)
            userTextInput.text.clear()
        }
    }
}