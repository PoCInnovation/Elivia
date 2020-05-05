package com.poc.leonapp

import android.content.Intent
import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import com.xwray.groupie.GroupAdapter
import com.xwray.groupie.ViewHolder

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        val intent = Intent(this, ChatActivity::class.java)

        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        startActivity(intent)
    }
}
