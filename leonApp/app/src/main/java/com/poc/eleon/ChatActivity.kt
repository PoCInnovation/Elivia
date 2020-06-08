package com.poc.eleon

import android.Manifest
import android.content.pm.PackageManager
import android.media.MediaRecorder
import android.net.Uri
import android.os.Build
import android.os.Bundle
import android.util.Log
import android.view.View
import android.view.animation.AccelerateDecelerateInterpolator
import android.view.animation.TranslateAnimation
import android.widget.ImageView
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.core.app.ActivityCompat
import androidx.core.content.ContextCompat
import com.xwray.groupie.GroupAdapter
import com.xwray.groupie.Item
import com.xwray.groupie.ViewHolder
import io.socket.client.IO
import io.socket.client.Socket
import io.socket.emitter.Emitter
import kotlinx.android.synthetic.main.activity_chat.*
import kotlinx.android.synthetic.main.activity_chat.view.*
import kotlinx.android.synthetic.main.bubble_leon.view.*
import kotlinx.android.synthetic.main.bubble_user.view.*
import org.json.JSONObject
import org.jsoup.Jsoup
import java.io.ByteArrayOutputStream
import java.io.File
import java.io.IOException
import java.io.InputStream

class ChatActivity : AppCompatActivity() {
    private var socket: Socket = IO.socket("http://192.168.0.14:1337")
    var mediaRecorder: MediaRecorder? = null
    private var FILE_RECORDING = ""

    private val PERMISSION_GRANTED = PackageManager.PERMISSION_GRANTED
    private val AUDIO_PERMISSION = Manifest.permission.RECORD_AUDIO
    private val PERMISSION_REQUEST_CODE = 100

    override fun onCreate(savedInstanceState: Bundle?) {
        val adapter = GroupAdapter<ViewHolder>()
        var showHideKeyboard:Boolean = true

        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_chat)
        initRecyclerViewAdapter(adapter)
        FILE_RECORDING = "${externalCacheDir?.absolutePath}/recorder.mp3"
        var tmp:Boolean = true
        VoiceRecord.setOnClickListener {
            tmp = if (tmp) {
                startRecord()
                false
            } else {
                stopRecord()
                true
            }
        }
        userLand.chatActivity_userSend.setOnClickListener {
            val userText = userLand.chatActivity_userKeyboard.text.toString()

            if (userText.isNotEmpty()) {
                addBubbleToView(adapter, R.layout.bubble_user, userText)
                userLand.chatActivity_userKeyboard.text.clear()
                sendQueryToLeon(userText)
            }
        }
        HideShowKeyboard.setOnClickListener {
            showHideKeyboard = if (!showHideKeyboard) {
                slideUp(findViewById(R.id.userLand))
                userLand.chatActivity_userKeyboard.visibility = View.VISIBLE
                rescaleImageView(VoiceRecord, 1f, 1f, 1000)
                true;
            } else {
                slideDown(findViewById(R.id.userLand))
                userLand.chatActivity_userKeyboard.visibility = View.INVISIBLE
                rescaleImageView(VoiceRecord, 1.5f, 1.5f, 1000)
                false;
            }
        }
        socket.on(Socket.EVENT_DISCONNECT, Emitter.Listener {
            Log.d("ChatActivity", "User has been disconnected from Leon")
        })
        socket.on(Socket.EVENT_CONNECT, Emitter.Listener {
            Log.d("ChatActivity", "User is now connected to Leon")
            socket.emit("init", "webapp")
        })
        socket.on("answer", Emitter.Listener { args ->
            var leonText:String = args[0].toString()

            leonText = Jsoup.parse(leonText).text()
            runOnUiThread {
                addBubbleToView(adapter, R.layout.bubble_leon, leonText)
            }
        })
        socket.on("recognized", Emitter.Listener {  args ->
            var userAudioText = args[0].toString()

            userAudioText = Jsoup.parse(userAudioText).text()
            runOnUiThread {
                addBubbleToView(adapter, R.layout.bubble_user, userAudioText)
            }
            sendQueryToLeon(userAudioText)
        })
        socket.connect()
    }

    private fun rescaleImageView(view: ImageView, x: Float, y: Float, duration: Long) {
        view.animate().scaleY(y).setInterpolator(AccelerateDecelerateInterpolator()).duration = duration
        view.animate().scaleX(x).setInterpolator(AccelerateDecelerateInterpolator()).duration = duration
    }

    private fun isPermissionGranted(): Boolean{
        return if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.M) checkSelfPermission(AUDIO_PERMISSION) == PERMISSION_GRANTED
        else return true

    }

    private fun requestAudioPermission(){
        if(Build.VERSION.SDK_INT >= Build.VERSION_CODES.M){
            requestPermissions(arrayOf(AUDIO_PERMISSION), PERMISSION_REQUEST_CODE)
        }
    }

    private fun startRecord() {
        Log.d("ChatActivity", "Start")
        rescaleImageView(VoiceRecord, 1.5f, 1.5f, 500)
        if(!isPermissionGranted()){
            requestAudioPermission()
            return
        }
        mediaRecorder = MediaRecorder()
        mediaRecorder!!.setAudioSource(MediaRecorder.AudioSource.MIC)
        mediaRecorder!!.setOutputFormat(MediaRecorder.OutputFormat.MPEG_4)
        mediaRecorder!!.setOutputFile(FILE_RECORDING)
        mediaRecorder!!.setAudioEncoder(MediaRecorder.AudioEncoder.AAC)
        mediaRecorder!!.prepare()
        mediaRecorder!!.start()
    }

    private fun stopRecord() {
        Log.d("ChatActivity", "Stop")
        rescaleImageView(VoiceRecord, 1f, 1f, 500)
        if (mediaRecorder == null)
            return
        mediaRecorder?.stop()
        mediaRecorder?.release()
        mediaRecorder = null
        sendAudio(FILE_RECORDING)
    }

    private fun initRecyclerViewAdapter(adapter: GroupAdapter<ViewHolder>) {
        ChatActivity_RecyclerView.adapter = adapter
    }

    private fun addBubbleToView(adapter: GroupAdapter<ViewHolder>, layoutId: Int, text: String) {
        adapter.add(Bubble(layoutId, text))
        ChatActivity_RecyclerView.scrollToPosition(adapter.itemCount - 1)
    }

    private fun sendQueryToLeon(text: String) {
        val obj = JSONObject()

        obj.put("value", text)
        socket.emit("query", obj)
    }

    private fun slideUp(view: View) {
        view.visibility = View.VISIBLE
        val animate = TranslateAnimation(
            0F,  // fromXDelta
            0F,  // toXDelta
            view.height.toFloat(),  // fromYDelta
            0F
        ) // toYDelta
        animate.duration = 200
        animate.fillAfter = true
        view.startAnimation(animate)
    }
    private fun slideDown(view: View) {
        val animate = TranslateAnimation(
            0F,  // fromXDelta
            0F,  // toXDelta
            0F,  // fromYDelta
            view.height.toFloat()
        ) // toYDelta
        animate.duration = 200
        animate.fillAfter = true
        view.startAnimation(animate)
    }

    private fun sendAudio(outputFile:String) {
        Log.d("ChatActivity", "Path to audio record : $outputFile")
        var soundBytes: ByteArray

        if (ContextCompat.checkSelfPermission(this, Manifest.permission.RECORD_AUDIO) != PackageManager.PERMISSION_GRANTED && ContextCompat.checkSelfPermission(this, Manifest.permission.WRITE_EXTERNAL_STORAGE) != PackageManager.PERMISSION_GRANTED) {
            val permissions = arrayOf(android.Manifest.permission.RECORD_AUDIO, android.Manifest.permission.WRITE_EXTERNAL_STORAGE, android.Manifest.permission.READ_EXTERNAL_STORAGE)
            ActivityCompat.requestPermissions(this, permissions,0)
        } else {
            try {
                val r1 = Uri.fromFile(File(outputFile))
                val inputStream = contentResolver.openInputStream(r1)
                Log.d("ChatActivity", "inputStream : $inputStream")
                soundBytes = ByteArray(inputStream!!.available())
                soundBytes = toByteArray(inputStream!!)!!
                Log.d("ChatActivity", "Audio converted")
                Toast.makeText(this, "Recording Finished $soundBytes", Toast.LENGTH_LONG).show()
                socket.emit("recognize", soundBytes)
                Log.d("ChatActivity", "Audio sent")
            } catch (e: Exception) {
                Log.d("ChatActivity", "Got an error : $e.message")
                e.printStackTrace()
            }
        }

    }

    @Throws(IOException::class)
    fun toByteArray(`in`: InputStream): ByteArray? {
        val out = ByteArrayOutputStream()
        var read = 0
        val buffer = ByteArray(1024)
        while (read != -1) {
            read = `in`.read(buffer)
            if (read != -1) out.write(buffer, 0, read)
        }
        out.close()
        return out.toByteArray()
    }
}

class Bubble(private val layoutId: Int, private  val text: String): Item<ViewHolder>() {

    override fun getLayout(): Int {
        return (layoutId)
    }

    override fun bind(viewHolder: ViewHolder, position: Int) {
        if (layoutId == R.layout.bubble_leon)
            viewHolder.itemView.TextView_bubbleLeon.text = text
        else
            viewHolder.itemView.TextView_bubbleUser.text = text
    }
}

