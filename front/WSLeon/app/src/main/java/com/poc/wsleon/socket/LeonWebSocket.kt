package com.poc.wsleon.socket

import android.app.Activity
import android.util.Log
import com.poc.wsleon.MainActivity
import com.poc.wsleon.plugin.core.PluginManager
import com.poc.wsleon.ui.LeonView
import okhttp3.Response
import okhttp3.WebSocket
import okhttp3.WebSocketListener
import okio.ByteString
import org.json.JSONObject

class LeonWebSocket(view: LeonView, plugins: PluginManager, activity: Activity): WebSocketListener() {
    private val chat = view
    private val pluginManager = plugins
    private val token = ByteArray(50)
    val mainActivity = activity

    override fun onOpen(webSocket: WebSocket, response: Response) {
        val connectionHeader = JSONObject()

        connectionHeader.put("type", 0)
        connectionHeader.put("content", "")
        connectionHeader.put("user_token", token)
        connectionHeader.put("locale", "en")
        connectionHeader.put("information", null)
        webSocket.send(connectionHeader.toString())
        // webSocket.close(1000,"Socket off")
    }

    override fun onMessage(webSocket: WebSocket, text: String) {
        val obj = JSONObject(text)
        val oliviaResponse = obj["response"].toString()
        val pluginName: String = obj["package"].toString()

        mainActivity.runOnUiThread {
            chat.addOliviaBubble(oliviaResponse)
        }
        pluginManager.run(pluginName, obj["data"] as JSONObject)
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