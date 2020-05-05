package com.poc.leonapp

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.os.Parcel
import android.os.Parcelable
import com.poc.leonapp.models.ChatLeon
import com.poc.leonapp.models.ChatUser
import com.poc.leonapp.models.User
import com.xwray.groupie.GroupAdapter
import com.xwray.groupie.Item
import com.xwray.groupie.ViewHolder
import kotlinx.android.parcel.Parcelize
import kotlinx.android.synthetic.main.activity_chat.*
import kotlinx.android.synthetic.main.chat_leon.view.*
import kotlinx.android.synthetic.main.chat_user.view.*

class ChatActivity : AppCompatActivity() {
    private val adapter = GroupAdapter<ViewHolder>()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_chat)
        chatRecyclerView.adapter = adapter
        userSend.setOnClickListener {
            performSendUserRequest()
        }
    }

    private fun performSendUserRequest() {
        // TODO : Send user request to Leon
        addUserRequestToChat()
        addLeonResponseToChat()
    }

    private fun addLeonResponseToChat() {
        val leon = User("Leon says hello !")

        adapter.add(ChatLeon(leon))
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
    }

    private fun addUserRequestToChat() {
        val user = User(userKeyboard.text.toString())

        adapter.add(ChatUser(user))
        userKeyboard.text.clear()
        chatRecyclerView.scrollToPosition(adapter.itemCount - 1)
    }
}
