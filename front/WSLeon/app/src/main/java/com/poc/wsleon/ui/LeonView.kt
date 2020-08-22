package com.poc.wsleon.ui

import androidx.recyclerview.widget.RecyclerView
import com.poc.wsleon.R
import com.xwray.groupie.GroupAdapter
import com.xwray.groupie.ViewHolder

class LeonView() {
    private val adapter = GroupAdapter<ViewHolder>()

    fun initRecyclerViewAdapter(view: RecyclerView) {
        view.adapter = adapter
    }
    fun addOliviaBubble(message: String) {
        this.adapter.add(Bubble(R.layout.bubble_olivia, message))
    }
    fun addUserBubble(message: String) {
        this.adapter.add(Bubble(R.layout.bubble_user, message))
    }
}