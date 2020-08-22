package com.poc.wsleon.socket

import android.util.Log
import com.poc.wsleon.ui.LeonView
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import okio.ByteString
import org.json.JSONObject

class LeonWebSocket(view: LeonView): WebSocketListener() {
    private val chat = view

    override fun onOpen(webSocket: WebSocket, response: Response) {
        val connectionHeader = JSONObject()
        val token = ByteArray(50)

        connectionHeader.put("type", 0)
        connectionHeader.put("content", "")
        connectionHeader.put("user_token", token)
        connectionHeader.put("locale", "en")
        connectionHeader.put("information", null)
        webSocket.send(connectionHeader.toString())
        // webSocket.close(1000,"Socket off")
    }

    override fun onMessage(webSocket: WebSocket, text: String) {
        // Log.d("LeonWebSocket", "Received message : [${text}]")
        val obj = JSONObject(text)
        val oliviaResponse = obj["content"].toString()

        chat.addOliviaBubble(oliviaResponse)
    }

    override fun onMessage(webSocket: WebSocket, bytes: ByteString) {
        Log.d("LeonWebSocket", "Received bytes message : [${bytes.hex()}]")
    }

    override fun onClosing(webSocket: WebSocket, code: Int, reason: String) {
        webSocket.close(1000, null)
        Log.d("LeonWebSocket", "Closing WebSocket connection with code [${code}] reason : [${reason}]")
    }

    override fun onFailure(webSocket: WebSocket, t: Throwable, response: Response?) {
        Log.d("LeonWebSocket", "WebSocket failure : [${t.message}]")
    }
    fun onSend(webSocket: WebSocket, content: String) {
        val token = ByteArray(50)
        val message = JSONObject()

        message.put("type", 1)
        message.put("content", content)
        message.put("user_token", token)
        message.put("locale", "en")
        message.put("information", null)
        webSocket.send(message.toString())
        chat.addUserBubble(content)
    }
}