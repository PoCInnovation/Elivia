package com.poc.leonapp.models

import android.os.Parcelable
import com.poc.leonapp.R
import com.xwray.groupie.Item
import com.xwray.groupie.ViewHolder
import kotlinx.android.parcel.Parcelize
import kotlinx.android.synthetic.main.chat_leon.view.*
import kotlinx.android.synthetic.main.chat_user.view.*

class ChatUser(private val user:User): Item<ViewHolder>() {
    override fun bind(viewHolder: ViewHolder, position: Int) {
        viewHolder.itemView.bubbleUser.text = user.request
    }

    override fun getLayout(): Int {
        return (R.layout.chat_user)
    }
}

class ChatLeon(private val user:User): Item<ViewHolder>() {
    override fun bind(viewHolder: ViewHolder, position: Int) {
        viewHolder.itemView.bubbleLeon.text = user.request
    }

    override fun getLayout(): Int {
        return (R.layout.chat_leon)
    }
}

@Parcelize
class User(val request:String): Parcelable {
    constructor() : this("")
}