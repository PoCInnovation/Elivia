package com.poc.wsleon.ui

import com.poc.wsleon.R
import com.xwray.groupie.Item
import com.xwray.groupie.ViewHolder
import kotlinx.android.synthetic.main.bubble_olivia.view.*
import kotlinx.android.synthetic.main.bubble_user.view.*

class Bubble(private val layoutId: Int, private  val text: String): Item<ViewHolder>() {

    override fun getLayout(): Int {
        return (layoutId)
    }

    override fun bind(viewHolder: ViewHolder, position: Int) {
        if (layoutId == R.layout.bubble_olivia)
            viewHolder.itemView.textViewOlivia.text = text
        else
            viewHolder.itemView.textViewUser.text = text
    }
}