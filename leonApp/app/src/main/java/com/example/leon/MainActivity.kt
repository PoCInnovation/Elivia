package com.example.leon

import android.graphics.Color
import android.os.Bundle
import android.util.TypedValue
import android.widget.LinearLayout
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import kotlinx.android.synthetic.main.activity_main.*


class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        userButton.setOnClickListener {
            userInput()
        }
    }
    private fun addTextToChat(userText:String) {
        // Get chatLayout
        val chatLayout:LinearLayout = findViewById(R.id.chatLayout)
        // Creates the new textView
        val newText:TextView = TextView(this)

        newText.text = userText
        newText.setTextColor(Color.parseColor("#ffffff"))
        newText.setTextSize(TypedValue.COMPLEX_UNIT_SP, 22.0F)
        chatLayout.addView(newText)
    }

    private fun userInput() {
        val userText:String = userKeyboard.text.toString()

        // TODO: Send the user input to Leon

        // Add the input to the scroll view
        addTextToChat(userText)
    }
}
