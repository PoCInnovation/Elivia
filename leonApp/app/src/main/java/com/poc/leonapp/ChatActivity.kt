package com.poc.leonapp

import android.os.Bundle
import android.util.Log
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
    private var list = ArrayList<String>()
    private var socket: Socket = IO.socket("http://192.168.1.30:1337")

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_chat)
        chatRecyclerView.adapter = adapter

        userSend.setOnClickListener {
            Log.d("ChatActivity", "User asked something")
            performSendUserRequest()
        }

        socket.on(Socket.EVENT_DISCONNECT, Emitter.Listener {
            Log.d("ChatActivity", "User has been disconnected from Leon")
        })
        socket.on(Socket.EVENT_CONNECT, Emitter.Listener {
            Log.d("ChatActivity", "User is now connected to Leon")
            socket.emit("init", "webapp")
        })

        socket.on("answer", Emitter.Listener {  args ->
            var resp: LeonMessage
            if (args.size == 1 && args[0] is JSONObject)
                resp = LeonMessage(args[0] as JSONObject)
            else
                resp = LeonMessage("Invalid response", "none")

            list.add(resp.msg)

            // Adding text to RecyclerView on UI
            runOnUiThread {
                addLeonResponseToChat(list.last())
            }
        })
        socket.connect()
        Log.d("ChatActivity", "User is connecting to Leon..")
    }

    private fun performSendUserRequest() {
        // TODO : Send user request to Leon
        addUserRequestToChat()
    }

    private fun addLeonResponseToChat(response: String) {
        val leon = User(response)

        adapter.add(ChatLeon(leon))
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
    }

    private fun addUserRequestToChat() {
        val msg: String = userKeyboard.text.toString()
        val user = User(msg)
        val obj = JSONObject();

        adapter.add(ChatUser(user))
        userKeyboard.text.clear()
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
        obj.put("value", msg)
        socket.emit("query", obj)
        Log.d("ChatActivity", "User has sent query to Leon")
    }

}
