package com.poc.leonapp

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import com.poc.leonapp.models.ChatLeon
import com.poc.leonapp.models.ChatUser
import com.poc.leonapp.models.User
import com.xwray.groupie.GroupAdapter
import com.xwray.groupie.ViewHolder
import io.socket.client.IO
import io.socket.client.Socket
import io.socket.emitter.Emitter
import kotlinx.android.synthetic.main.activity_chat.*
import org.json.JSONObject

class ChatActivity : AppCompatActivity() {
    private val adapter = GroupAdapter<ViewHolder>()
    var socket: Socket = IO.socket("http://192.168.1.30:1337")

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_chat)
        chatRecyclerView.adapter = adapter
        userSend.setOnClickListener {
            performSendUserRequest()
        }

        socket.on(Socket.EVENT_DISCONNECT, Emitter.Listener {
            println("disconnected")
        })
        socket.on(Socket.EVENT_CONNECT, Emitter.Listener {
            println("connected")
            socket.emit("init", "webapp")
        })
        socket.on("answer", Emitter.Listener {
            addLeonResponseToChat("response")
        })
        socket.connect()
        println("SOCKET CONNECT")
    }

    private fun performSendUserRequest() {
        // TODO : Send user request to Leon
        addUserRequestToChat()
        //addLeonResponseToChat("okok")
    }

    private fun addLeonResponseToChat(response: String) {
        val leon = User(response)

        adapter.add(ChatLeon(leon))
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
    }

    private fun addUserRequestToChat() {
        val msg: String = userKeyboard.text.toString()
        val user = User(msg)

        adapter.add(ChatUser(user))
        userKeyboard.text.clear()
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
        val obj = JSONObject();
        obj.put("value", msg)
        socket.emit("query", obj)
    }

}
